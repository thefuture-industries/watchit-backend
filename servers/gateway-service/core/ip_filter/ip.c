/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include "ip.h"
#include <string.h>
#include <stdio.h>


void get_client_ip(SOCKET client_sock, char *ip_buffer, int buffer_len) {
    struct sockaddr_in client_addr;
    int addr_len = sizeof(client_addr);

    if (getpeername(client_sock, (struct sockaddr *)&client_addr, &addr_len)) {
        // Если произошла ошибка — записываем "unknown" в буфер
        strncpy(ip_buffer, "unknown", buffer_len);
        ip_buffer[buffer_len - 1] = '\0';
        return;
    }

    // Преобразуем бинарный IP в строковое представление
    if (inet_ntop(AF_INET, &client_addr.sin_addr, ip_buffer, buffer_len) == NULL) {
        // Если произошла ошибка преобразования
        strncpy(ip_buffer, "unknown", buffer_len);
        ip_buffer[buffer_len - 1] = '\0';
    }
}
