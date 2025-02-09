/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <winsock2.h>
#include <ws2tcpip.h>
#include <pthread.h>
#include "include/services_config.c"

#pragma comment(lib, "ws2_32.lib")

#define LISTEN_PORT 8080
#define BUF_SIZE 8192

// Declare the mutex
pthread_mutex_t rr_mutex;

/*
 * Функция forward_to_backend() - пересылает запрос на backend‑сервер
 */
int forward_to_backend(int client_sock, const char* backend_ip, int backend_port, const char* request, int request_len) {
    int backend_sock;
    struct sockaddr_in backend_addr;

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
        closesocket(backend_sock);
        return -1;
    }

    if (connect(backend_sock, (struct sockaddr *)&backend_addr, sizeof(backend_addr)) < 0) {
        perror("connect");
        closesocket(backend_sock);
        return -1;
    }

    int sent = 0;
    while (sent < request_len) {
        int n = send(backend_sock, request + sent, request_len - sent, 0);
        if (n < 0) {
            perror("send");
            closesocket(backend_sock);
            return -1;
        }
        sent += n;
    }

    char backend_buf[BUF_SIZE];
    int r;
    while ((r = recv(backend_sock, backend_buf, BUF_SIZE, 0)) > 0) {
        int total_sent = 0;
        while (total_sent < r) {
            int n = send(client_sock, backend_buf + total_sent, r - total_sent, 0);
            if (n < 0) {
                perror("send to client");
                break;
            }
            total_sent += n;
        }
    }

    closesocket(backend_sock);
    return 0;
}

/*
 * Функция handle_client() - обрабатывает запрос в отдельном потоке
 */
void *handle_client(void *arg) {
    int client_sock = *((int *)arg);
    free(arg);

    char buffer[BUF_SIZE];
    int received = recv(client_sock, buffer, BUF_SIZE - 1, 0);
    if (received <= 0) {
        closesocket(client_sock);
        pthread_exit(NULL);
    }
    buffer[received] = '\0';

    char method[16], url[256], protocol[16];
    if (sscanf(buffer, "%15s %255s %15s", method, url, protocol) != 3) {
        char *error_msg = "HTTP/1.1 400 Bad Request\r\n\r\n";
        send(client_sock, error_msg, strlen(error_msg), 0);
        closesocket(client_sock);
        pthread_exit(NULL);
    }

    printf("Запрос: %s %s %s\n", method, url, protocol);

    if (strstr(url, "/user/") != NULL) {
        Backend selected;
        pthread_mutex_lock(&rr_mutex);
        selected = user_backends[user_rr_index];
        user_rr_index = (user_rr_index + 1) % user_backend_count;
        pthread_mutex_unlock(&rr_mutex);

        printf("Пересылаем запрос в микросервис user на %s:%d\n", selected.ip, selected.port);
        forward_to_backend(client_sock, selected.ip, selected.port, buffer, received);
    } else {
        char *not_found = "HTTP/1.1 404 Not Found\r\n\r\n";
        send(client_sock, not_found, strlen(not_found), 0);
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
    server_addr.sin_port = htons(LISTEN_PORT);
    server_addr.sin_addr.s_addr = INADDR_ANY;
    memset(server_addr.sin_zero, 0, sizeof(server_addr.sin_zero));

    if (bind(server_sock, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        perror("bind");
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

    printf("Многопоточный сервер слушает порт %d...\n", LISTEN_PORT);

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
