#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <winsock2.h>
#include <ws2tcpip.h>

#define SERVER_URL "http://localhost:8080/api/v1/user/sync"
#define MICRO_PATH "/micro/user/sync"
#define SERVER_IP "127.0.0.1"
#define SERVER_PORT 8080

int main() {
    WSADATA wsaData;
    SOCKET sock;
    struct sockaddr_in server_addr;
    char request[1024];
    char response[4096];
    int bytes_sent, bytes_received;

    // Инициализация Winsock
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0) {
        fprintf(stderr, "WSAStartup failed: %d\n", WSAGetLastError());
        return 1;
    }

    // Создание сокета
    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == INVALID_SOCKET) {
        fprintf(stderr, "Socket creation failed: %d\n", WSAGetLastError());
        WSACleanup();
        return 1;
    }

    // Настройка адреса сервера
    memset(&server_addr, 0, sizeof(server_addr));
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(SERVER_PORT);
    if (inet_pton(AF_INET, SERVER_IP, &server_addr.sin_addr) <= 0) {
        fprintf(stderr, "Invalid address/ Address not supported\n");
        closesocket(sock);
        WSACleanup();
        return 1;
    }

    // Подключение к серверу
    if (connect(sock, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        fprintf(stderr, "Connection failed: %d\n", WSAGetLastError());
        closesocket(sock);
        WSACleanup();
        return 1;
    }

    // Формирование HTTP-запроса
    snprintf(request, sizeof(request), "GET %s HTTP/1.1\r\nHost: %s:%d\r\nConnection: close\r\n\r\n", MICRO_PATH, SERVER_IP, SERVER_PORT);

    // Получение ответа от сервера
    bytes_received = recv(sock, response, sizeof(response) - 1, 0);
    if (bytes_received < 0) {
        fprintf(stderr, "Receive failed: %d\n", WSAGetLastError());
        closesocket(sock);
        WSACleanup();
        return 1;
    }

    // Завершение строки ответа нулевым символом
    response[bytes_received] = '\0';

    // Вывод ответа
    printf("Response:\n%s\n", response);

    // Закрытие сокета
    closesocket(sock);
    WSACleanup();

    return 0;
}
