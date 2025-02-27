#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#ifdef _WIN32
    #include <winsock2.h>
    #include <windows.h>
    #define PLATFORM_CLOSE_SOCKET(s) closesocket(s)
#else
    #include <sys/socket.h>
    #include <netinet/in.h>
    #include <arpa/inet.h>
    #include <unistd.h>
    #define PLATFORM_CLOSE_SOCKET(s) close(s)
#endif
#include <ws2tcpip.h>
#include <pthread.h>

#include "client_handler.h"
#include "common.h"

#include "services_paths.h"
#include "route_handler.h"
#include "../security/suspicious_checker.h"
#include "headers.h"
#include "../ip_filter/ip.h"
#include "../utiling/logger.h"
#include "../config/config.h"

#define BUF_SIZE 8192
#define MAX_REQUEST_SIZE 4096

/*
 * Функция forward_to_backend() - пересылает HTTP-запрос на backend‑сервер
 */
int forward_to_backend(int client_sock, const char *backend_ip, int backend_port, const char *method, const char *path, const char *request) {
    int backend_sock;
    struct sockaddr_in backend_addr;
    // char url[256];

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
        PLATFORM_CLOSE_SOCKET(backend_sock);
        return -1;
    }

    if (connect(backend_sock, (struct sockaddr *)&backend_addr, sizeof(backend_addr)) < 0) {
        perror("connect");
        PLATFORM_CLOSE_SOCKET(backend_sock);
        return -1;
    }

    // Формируем HTTP-запрос
    char http_request[MAX_REQUEST_SIZE];
    snprintf(http_request, sizeof(http_request), "%s %s HTTP/1.1\r\nHost: %s:%d\r\n\r\n",
             method, path, backend_ip, backend_port);

    printf("%s", http_request);

    // Отправляем запрос на backend-сервер
    if (send(backend_sock, http_request, strlen(http_request), 0) == -1) {
        perror("send");
        PLATFORM_CLOSE_SOCKET(backend_sock);
        return -1;
    }

    // Получаем ответ от бэкенда и пересылаем его клиенту
    char response[MAX_REQUEST_SIZE];
    ssize_t bytes_received;
    while ((bytes_received = recv(backend_sock, response, MAX_REQUEST_SIZE, 0)) > 0) {
        send(client_sock, response, bytes_received, 0);
    }

    PLATFORM_CLOSE_SOCKET(backend_sock);
    return 0;
}

/*
 * Функция handle_client() - обработка запроса в отдельном потоке
 */
void *handle_client(void *arg) {
    int client_sock = *((int *)arg);
    free(arg);

    // Получаем IP клиента
    char clientIP[INET_ADDRSTRLEN];
    get_client_ip(client_sock, clientIP, sizeof(clientIP));

    // Проверка частоты запросов с IP
    if (check_ip_rate(clientIP)) {
        char *banned_message = "HTTP/1.1 403 Forbidden\r\n\r\nIP banned";
        send(client_sock, banned_message, strlen(banned_message), 0);
        printf("HTTP/1.1 403 Forbidden\r\n\r\nIP banned");
        log_request(clientIP, "BLOCKED", "N/A", "403", "IP banned");
        log_security("IP banned");
        PLATFORM_CLOSE_SOCKET(client_sock);
        pthread_exit(NULL);
    }

    char buffer[BUF_SIZE];
    int received = recv(client_sock, buffer, BUF_SIZE - 1, 0);
    if (received <= 0) {
        PLATFORM_CLOSE_SOCKET(client_sock);
        pthread_exit(NULL);
    }
    buffer[received] = '\0';

    // Если запрос содержит подозрительную активность - прекращаем выполнение
    if (is_request_suspicious(buffer)) {
        char *forbidden_msg = "HTTP/1.1 403 Forbidden\r\n\r\nSuspicious activity detected";
        send(client_sock, forbidden_msg, strlen(forbidden_msg), 0);
        printf("HTTP/1.1 403 Forbidden\r\n\r\nSuspicious activity detected");
        log_security("Suspicious activity detected");
        PLATFORM_CLOSE_SOCKET(client_sock);
        pthread_exit(NULL);
    }

    char method[16], url[256], protocol[16];
    if (sscanf(buffer, "%15s %255s %15s", method, url, protocol) != 3) {
        char *error_msg = "HTTP/1.1 400 Bad Request\r\n\r\n";
        send(client_sock, error_msg, strlen(error_msg), 0);
        log_request(clientIP, "UNKNOWN", "UNKNOWN", "400", error_msg);
        log_error("Bad Request: Unable to parse the request.");
        PLATFORM_CLOSE_SOCKET(client_sock);
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

        char modified_request[BUF_SIZE];
        transform_request(buffer, modified_request, BUF_SIZE, selected.ip, selected.port);

        char new_url[256];
        const char *old_prefix = "/api/v1";
        const char *micro_prefix = "/micro";
        size_t old_prefix_len = strlen(old_prefix);
        if (strncmp(url, old_prefix, old_prefix_len) == 0) {
            snprintf(new_url, sizeof(new_url), "%s%s", micro_prefix, url + old_prefix_len);
        } else {
            strncpy(new_url, url, sizeof(new_url));
            new_url[sizeof(new_url) - 1] = '\0';
        }

        printf("We forward the request to the microservice user to %s:%d\n", selected.ip, selected.port);
        char message[256];
        sprintf(message, "We forward the request to the microservice user to %s:%d", selected.ip, selected.port);
        log_request(clientIP, method, url, status, message);
        forward_to_backend(client_sock, selected.ip, selected.port, method, new_url, modified_request);
    }
    else if (strstr(url, "/blog/") != NULL) {
        Backend selected;
        pthread_mutex_lock(&rr_mutex);
        selected = blog_backends[blog_rr_index];
        blog_rr_index = (blog_rr_index + 1) % blog_backend_count;
        pthread_mutex_unlock(&rr_mutex);

        char modified_request[BUF_SIZE];
        transform_request(buffer, modified_request, BUF_SIZE, selected.ip, selected.port);

        char new_url[256];
        const char *old_prefix = "/api/v1";
        const char *micro_prefix = "/micro";
        size_t old_prefix_len = strlen(old_prefix);
        if (strncmp(url, old_prefix, old_prefix_len) == 0) {
            snprintf(new_url, sizeof(new_url), "%s%s", micro_prefix, url + old_prefix_len);
        } else {
            strncpy(new_url, url, sizeof(new_url));
            new_url[sizeof(new_url) - 1] = '\0';
        }

        printf("We are forwarding the request to the blog microservice to %s:%d\n", selected.ip, selected.port);
        char message[256];
        sprintf(message, "We forward the request to the microservice blog to %s:%d", selected.ip, selected.port);
        log_request(clientIP, method, url, status, message);
        forward_to_backend(client_sock, selected.ip, selected.port, method, new_url, modified_request);
    }
    else if (strstr(url, "/movie/") != NULL) {
        Backend selected;
        pthread_mutex_lock(&rr_mutex);
        selected = movie_backends[movie_rr_index];
        movie_rr_index = (movie_rr_index + 1) % movie_backend_count;
        pthread_mutex_unlock(&rr_mutex);

        char modified_request[BUF_SIZE];
        transform_request(buffer, modified_request, BUF_SIZE, selected.ip, selected.port);

        char new_url[256];
        const char *old_prefix = "/api/v1";
        const char *micro_prefix = "/micro";
        size_t old_prefix_len = strlen(old_prefix);
        if (strncmp(url, old_prefix, old_prefix_len) == 0) {
            snprintf(new_url, sizeof(new_url), "%s%s", micro_prefix, url + old_prefix_len);
        } else {
            strncpy(new_url, url, sizeof(new_url));
            new_url[sizeof(new_url) - 1] = '\0';
        }

        printf("We are forwarding the request to the movie microservice to %s:%d\n", selected.ip, selected.port);
        char message[256];
        sprintf(message, "We forward the request to the microservice movie to %s:%d", selected.ip, selected.port);
        log_request(clientIP, method, url, status, message);
        forward_to_backend(client_sock, selected.ip, selected.port, method, new_url, modified_request);
    }
    else {
        char *not_found = "HTTP/1.1 404 Not Found\r\n\r\n";
        send(client_sock, not_found, strlen(not_found), 0);
        log_request(clientIP, method, url, status, not_found);
    }

    PLATFORM_CLOSE_SOCKET(client_sock);
    pthread_exit(NULL);
}
