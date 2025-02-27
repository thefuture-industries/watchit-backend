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

#include "server.h"
#include "networking/client_handler.h"
#include "config/config.h"
#include "networking/common.h"

// #pragma comment(lib, "ws2_32.lib")

/*
 * Функция start_server() - создаёт серверный сокет, привязывает его к указанному
 * адресу и порту, начинает прослушивание и для каждого входящего соединения запускает
 * новый поток для его обработки
 */
void start_server(const char* listen_addr, const char* listen_port) {
    // Инициализация мьютекса
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
        PLATFORM_CLOSE_SOCKET(server_sock);
        WSACleanup();
        exit(EXIT_FAILURE);
    }

    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(atoi(listen_port));

    if (inet_pton(AF_INET, listen_addr, &server_addr.sin_addr) <= 0) {
        perror("Invalid listen_addr");
        WSACleanup();
        exit(EXIT_FAILURE);
    }
    memset(server_addr.sin_zero, 0, sizeof(server_addr.sin_zero));

    if (bind(server_sock, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        int err = WSAGetLastError();
        fprintf(stderr, "bind failed with error: %d\n", err);
        PLATFORM_CLOSE_SOCKET(server_sock);
        WSACleanup();
        exit(EXIT_FAILURE);
    }

    if (listen(server_sock, 10) < 0) {
        perror("listen");
        PLATFORM_CLOSE_SOCKET(server_sock);
        WSACleanup();
        exit(EXIT_FAILURE);
    }

    printf("The multithreaded server listens on the port %s...\n", listen_port);

    while (1) {
        struct sockaddr_in client_addr;
        int addr_len = sizeof(client_addr);
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
            PLATFORM_CLOSE_SOCKET(*client_sock_ptr);
            free(client_sock_ptr);
            continue;
        }
        pthread_detach(tid);
    }

    PLATFORM_CLOSE_SOCKET(server_sock);
    pthread_mutex_destroy(&rr_mutex);
    WSACleanup();
}
