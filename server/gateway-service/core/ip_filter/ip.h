/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#ifndef IP_FILTER_H
#define IP_FILTER_H

#include <winsock2.h>
#include <ws2tcpip.h>

void get_client_ip(SOCKET client_sock, char *ip_buffer, int buffer_len);

int check_ip_rate(const char *clientIP);

#endif
