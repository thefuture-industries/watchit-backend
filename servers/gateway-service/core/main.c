/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <windows.h>
#include <winsock2.h>
#include <ws2tcpip.h>
#include <pthread.h>

#include "networking/services_paths.h"
#include "networking/route_handler.h"
#include "security/suspicious_checker.h"
#include "networking/headers.h"
#include "ip_filter/ip.h"
#include "utiling/logger.h"
#include "config/config.h"

#pragma comment(lib, "ws2_32.lib")

// Declare the variable without initialization
char* listen_port;
char* listen_addr;

#define BUF_SIZE 8192
#define MAX_REQUEST_SIZE 4096

// Declare the mutex
pthread_mutex_t rr_mutex;

/*
 * Функция forward_to_backend() - пересылает запрос на backend‑сервер
 */
// int forward_to_backend(int client_sock, const char* backend_ip, int backend_port, const char* request, int request_len) {
//     int backend_sock;
//     struct sockaddr_in backend_addr;

//     backend_sock = socket(AF_INET, SOCK_STREAM, 0);
//     if (backend_sock < 0) {
//         perror("socket");
//         return -1;
//     }

//     memset(&backend_addr, 0, sizeof(backend_addr));
//     backend_addr.sin_family = AF_INET;
//     backend_addr.sin_port = htons(backend_port);
//     if (inet_pton(AF_INET, backend_ip, &backend_addr.sin_addr) <= 0) {
//         perror("inet_pton");
//         closesocket(backend_sock);
//         return -1;
//     }

//     if (connect(backend_sock, (struct sockaddr *)&backend_addr, sizeof(backend_addr)) < 0) {
//         perror("connect");
//         closesocket(backend_sock);
//         return -1;
//     }

//     int sent = 0;
//     while (sent < request_len) {
//         int n = send(backend_sock, request + sent, request_len - sent, 0);
//         if (n < 0) {
//             perror("send");
//             closesocket(backend_sock);
//             return -1;
//         }
//         sent += n;
//     }

//     char backend_buf[BUF_SIZE];
//     int r;
//     while ((r = recv(backend_sock, backend_buf, BUF_SIZE, 0)) > 0) {
//         int total_sent = 0;
//         while (total_sent < r) {
//             int n = send(client_sock, backend_buf + total_sent, r - total_sent, 0);
//             if (n < 0) {
//                 perror("send to client");
//                 break;
//             }
//             total_sent += n;
//         }
//     }

//     closesocket(backend_sock);
//     return 0;
// }

int forward_to_backend(int client_sock, const char *backend_ip, int backend_port, const char *request, int request_len) {
    int backend_sock;
    struct sockaddr_in backend_addr;

    printf("Request: %s\n", request);

    backend_sock = socket(AF_INET, SOCK_STREAM, 0);
    if (backend_sock < 0) {
        perror("socket");
        return -1;
    }

    memset(&backend_addr, 0, sizeof(backend_addr));
    backend_addr.sin_family = AF_INET;
    backend_addr.sin_port = htons(backend_port);
    if (inet_pton(AF_INET, backend_ip, &backend_addr.sin_addr) <= 0) {
        perror("inet_pton");
        close(backend_sock);
        return -1;
    }

    if (connect(backend_sock, (struct sockaddr *)&backend_addr, sizeof(backend_addr)) < 0) {
        perror("connect");
        close(backend_sock);
        return -1;
    }

    // Отправляем запрос на бэкенд-сервер
    if (send(backend_sock, request, strlen(request), 0) == -1) {
        perror("send");
        close(backend_sock);
        return -1;
    }

    // Получаем ответ от бэкенда и передаем его клиенту
    char response[MAX_REQUEST_SIZE];
    ssize_t bytes_received;
    while ((bytes_received = recv(backend_sock, response, MAX_REQUEST_SIZE, 0)) > 0) {
        send(client_sock, response, bytes_received, 0);
    }

    close(backend_sock);
    return 0;
}

/*
 * Функция handle_client() - обрабатывает запрос в отдельном потоке
 */
void *handle_client(void *arg) {
    int client_sock = *((int *)arg);
    free(arg);

    // Получаем IP клиента (устройство)
    char clientIP[INET_ADDRSTRLEN];
    get_client_ip(client_sock, clientIP, sizeof(clientIP));

    // Проверяем частоту запросов с данного IP
    if (check_ip_rate(clientIP)) {
        char *banned_message = "HTTP/1.1 403 Forbidden\r\n\r\nIP banned";
        send(client_sock, banned_message, strlen(banned_message), 0);
        printf("HTTP/1.1 403 Forbidden\r\n\r\nIP banned");
        log_request(clientIP, "BLOCKED", "N/A", "403", "IP banned");
        log_security("IP banned");
        closesocket(client_sock);
        pthread_exit(NULL);
    }

    char buffer[BUF_SIZE];
    int received = recv(client_sock, buffer, BUF_SIZE - 1, 0);
    if (received <= 0) {
        closesocket(client_sock);
        pthread_exit(NULL);
    }
    buffer[received] = '\0';

    // Если в запросе (либо URL, либо тело) обнаружен подозрительный текст, закрываем соединение
    if (is_request_suspicious(buffer)) {
        char *forbidden_msg = "HTTP/1.1 403 Forbidden\r\n\r\nSuspicious activity detected";
        send(client_sock, forbidden_msg, strlen(forbidden_msg), 0);
        printf("HTTP/1.1 403 Forbidden\r\n\r\nSuspicious activity detected");
        log_security("Suspicious activity detected");
        closesocket(client_sock);
        pthread_exit(NULL);
    }

    // Преобразуем запрос: меняем префикс "/api/v1" на "/micro"
    char modified_request[BUF_SIZE];
    transform_request(buffer, modified_request, BUF_SIZE);

    char method[16], url[256], protocol[16];
    if (sscanf(buffer, "%15s %255s %15s", method, url, protocol) != 3) {
        char *error_msg = "HTTP/1.1 400 Bad Request\r\n\r\n";
        send(client_sock, error_msg, strlen(error_msg), 0);
        log_request(clientIP, "UNKNOWN", "UNKNOWN", "400", error_msg);
        log_error("Bad Request: Unable to parse the request.");
        closesocket(client_sock);
        pthread_exit(NULL);
    }

    printf("Request: %s %s %s\n", method, url, protocol);

    const char *status = "200 (Forwarded)";
    if (strstr(url, "/user/") != NULL) {
        Backend selected;
        pthread_mutex_lock(&rr_mutex);
        selected = user_backends[user_rr_index];
        user_rr_index = (user_rr_index + 1) % user_backend_count;
        pthread_mutex_unlock(&rr_mutex);

        printf("We forward the request to the microservice user to %s:%d\n", selected.ip, selected.port);
        char message[256];
        sprintf(message, "We forward the request to the microservice user to %s:%d", selected.ip, selected.port);
        log_request(clientIP, method, url, status, message);
        forward_to_backend(client_sock, selected.ip, selected.port, modified_request, strlen(modified_request));
    } else if (strstr(url, "/blog/") != NULL) {
        Backend selected;
        pthread_mutex_lock(&rr_mutex);
        selected = blog_backends[blog_rr_index];
        blog_rr_index = (blog_rr_index + 1) % blog_backend_count;
        pthread_mutex_unlock(&rr_mutex);

        printf("We are forwarding the request to the blog microservice to %s:%d\n", selected.ip, selected.port);
        char message[256];
        sprintf(message, "We forward the request to the microservice blog to %s:%d", selected.ip, selected.port);
        log_request(clientIP, method, url, status, message);
        forward_to_backend(client_sock, selected.ip, selected.port, modified_request, strlen(modified_request));
    } else if (strstr(url, "/movie/") != NULL) {
        Backend selected;
        pthread_mutex_lock(&rr_mutex);
        selected = movie_backends[movie_rr_index];
        movie_rr_index = (movie_rr_index + 1) % movie_backend_count;
        pthread_mutex_unlock(&rr_mutex);

        printf("We are forwarding the request to the movie microservice to %s:%d\n", selected.ip, selected.port);
        char message[256];
        sprintf(message, "We forward the request to the microservice movie to %s:%d", selected.ip, selected.port);
        log_request(clientIP, method, url, status, message);
        forward_to_backend(client_sock, selected.ip, selected.port, modified_request, strlen(modified_request));
    } else {
        char *not_found = "HTTP/1.1 404 Not Found\r\n\r\n";
        send(client_sock, not_found, strlen(not_found), 0);
        log_request(clientIP, method, url, status, not_found);
    }

    closesocket(client_sock);
    pthread_exit(NULL);
}

/*
 * Функция main() - инициирует Winsock, запускает сервер и принимает подключения
 */
int main() {
    SetConsoleCP(65001);
    SetConsoleOutputCP(65001);

    WSADATA wsaData;
    if (WSAStartup(MAKEWORD(2,2), &wsaData) != 0) {
        fprintf(stderr, "WSAStartup failed: %d\n", WSAGetLastError());
        system("pause");
        exit(EXIT_FAILURE);
    }

    listen_port = get_config("listen_port");
    listen_addr = get_config("listen_addr");

    pthread_mutex_init(&rr_mutex, NULL);

    int server_sock;
    struct sockaddr_in server_addr;

    server_sock = socket(AF_INET, SOCK_STREAM, 0);
    if (server_sock < 0) {
        perror("socket");
        WSACleanup();
        exit(EXIT_FAILURE);
    }

    int opt = 1;
    if (setsockopt(server_sock, SOL_SOCKET, SO_REUSEADDR, (char *)&opt, sizeof(opt)) < 0) {
        perror("setsockopt");
        closesocket(server_sock);
        WSACleanup();
        exit(EXIT_FAILURE);
    }

    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(atoi(listen_port));
    // server_addr.sin_addr.s_addr = INADDR_ANY;

    if (inet_pton(AF_INET, listen_addr, &server_addr.sin_addr) <= 0) {
        perror("Invalid listen_addr");
        WSACleanup();
        exit(EXIT_FAILURE);
    }

    memset(server_addr.sin_zero, 0, sizeof(server_addr.sin_zero));

    if (bind(server_sock, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        int err = WSAGetLastError();
        fprintf(stderr, "bind failed with error: %d\n", err);
        closesocket(server_sock);
        WSACleanup();
        exit(EXIT_FAILURE);
    }

    if (listen(server_sock, 10) < 0) {
        perror("listen");
        closesocket(server_sock);
        WSACleanup();
        exit(EXIT_FAILURE);
    }

    printf("The multithreaded server listens on the port %s...\n", listen_port);

    while (1) {
        struct sockaddr_in client_addr;
        socklen_t addr_len = sizeof(client_addr);
        int *client_sock_ptr = malloc(sizeof(int));
        if (!client_sock_ptr) {
            perror("malloc");
            continue;
        }
        *client_sock_ptr = accept(server_sock, (struct sockaddr *)&client_addr, &addr_len);
        if (*client_sock_ptr < 0) {
            perror("accept");
            free(client_sock_ptr);
            continue;
        }

        pthread_t tid;
        if (pthread_create(&tid, NULL, handle_client, client_sock_ptr) != 0) {
            perror("pthread_create");
            closesocket(*client_sock_ptr);
            free(client_sock_ptr);
            continue;
        }
        pthread_detach(tid);
    }

    closesocket(server_sock);
    pthread_mutex_destroy(&rr_mutex);
    WSACleanup();
    system("pause");
    return 0;
}
