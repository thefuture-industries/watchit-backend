/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include <stdio.h>
#include <stdlib.h>
#ifdef _WIN32
    #include <winsock2.h>
    #include <windows.h>
    #define PLATFORM_CLOSE_SOCKET(s) closesocket(s)
#else
    #include <sys/socket.h>
    #include <netinet/in.h>
    #include <arpa/inet.h>
    #include <unistd.h>
    #include <unistd.h>
    #define PLATFORM_CLOSE_SOCKET(s) close(s)
#endif
#include "config/config.h"
#include "server.h"

// #pragma comment(lib, "ws2_32.lib")

int main() {
    // Устанавливаем кодировку консоли для корректного отображения UTF-8
    SetConsoleCP(65001);
    SetConsoleOutputCP(65001);

    // Инициализация Winsock
    WSADATA wsaData;
    if (WSAStartup(MAKEWORD(2,2), &wsaData) != 0) {
        fprintf(stderr, "WSAStartup failed: %d\n", WSAGetLastError());
        system("pause");
        exit(EXIT_FAILURE);
    }

    // Получаем настройки из конфигурационного файла
    char* listen_port = get_config("listen_port");
    char* listen_addr = get_config("listen_addr");

    // Запускаем сервер
    start_server(listen_addr, listen_port);

    system("pause");
    return 0;
}
