/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#ifndef IP_FILTER_H
#define IP_FILTER_H

#ifdef _WIN32
    #include <winsock2.h>
#else
    #include <sys/socket.h>
    #include <netinet/in.h>
    #include <arpa/inet.h>
    #include <unistd.h>
#endif
#include <ws2tcpip.h>

void get_client_ip(SOCKET client_sock, char *ip_buffer, int buffer_len);

int check_ip_rate(const char *clientIP);

#endif
