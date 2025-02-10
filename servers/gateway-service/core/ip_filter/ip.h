/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#ifndef CLIENT_INFO_H
#define CLIENT_INFO_H

#include <winsock2.h>
#include <ws2tcpip.h>

// Функция заполняет буфер ip_buffer IP-адресом клиента, полученным из client_sock
void get_client_ip(SOCKET client_sock, char *ip_buffer, int buffer_len);

#endif
