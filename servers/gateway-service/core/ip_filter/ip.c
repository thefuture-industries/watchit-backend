/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include "ip.h"
#include <string.h>
#include <stdio.h>

void get_client_ip(SOCKET client_sock, char *ip_buffer, int buffer_len) {
    struct sockaddr_in client_addr;
    int addrLen = sizeof(client_addr);
    if (getpeername(client_sock, (struct sockaddr*)&client_addr, &addrLen) != 0) {
        // Если ошибка – записываем значение "unknown"
        strncpy(ip_buffer, "unknown", buffer_len);
        ip_buffer[buffer_len - 1] = '\0';
        return;
    }
    inet_ntop(AF_INET, &client_addr.sin_addr, ip_buffer, buffer_len);
}
