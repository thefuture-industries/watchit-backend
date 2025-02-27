/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include "route_handler.h"
#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <ctype.h>

void transform_request(const char *orig_req, char *new_req, size_t new_req_size, const char *backend_ip, int backend_port) {
    // Находим конец первой строки (оканчивается на "\r\n")
    const char *eol = strstr(orig_req, "\r\n");
    if (!eol) {
        // Если нет символов перевода строки, просто копируем оригинал
        strncpy(new_req, orig_req, new_req_size);
        new_req[new_req_size - 1] = '\0';
        return;
    }
    size_t first_line_len = eol - orig_req;

    // Разбираем первую строку на метод, URL и протокол
    char method[32], url[256], protocol[32];
    if (sscanf(orig_req, "%31s %255s %31s", method, url, protocol) != 3) {
        strncpy(new_req, orig_req, new_req_size);
        new_req[new_req_size - 1] = '\0';
        return;
    }

    // Если URL начинается с "/api/v1", заменяем этот префикс на "/micro"
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

    // Собираем новую первую строку
    char new_first_line[512];
    snprintf(new_first_line, sizeof(new_first_line), "%s %s %s\r\n", method, new_url, protocol);

    // Формируем полный преобразованный запрос: новая первая строка + оставшиеся заголовки
    const char *headers = orig_req + first_line_len;

    // Копируем заголовки до первого вхождения Host:
    char temp_headers[2048];  // Буфер для хранения временных заголовков
    char *temp_ptr = temp_headers;
    const char *current_pos = headers;
    // bool found_host = false;

    while (*current_pos) {
        if (strncasecmp(current_pos, "Host:", 5) == 0) {
            // Пропускаем заголовок Host
            current_pos += 6;  // Пропускаем "Host:"
            while (*current_pos != '\n') {
                current_pos++;
            }
            current_pos++;     // Переход к следующему заголовку
        } else {
            // Копируем текущую часть заголовков в temp_headers
            while (*current_pos != '\n') {
                *temp_ptr++ = *current_pos++;
            }
            *temp_ptr++ = '\n';  // Добавляем символ новой строки
            current_pos++;       // Переход к следующему заголовку
        }
    }
    *temp_ptr = '\0';  // Завершаем строку нулевым символом

    // char host_header[256];
    // snprintf(host_header, sizeof(host_header), "Host: %s:%d\r\n", backend_ip, backend_port);
    // strcat(temp_headers, host_header);

    snprintf(new_req, new_req_size, "%s%s", new_first_line, temp_headers);
}
