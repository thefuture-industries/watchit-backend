/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#ifndef ROUTE_HANDLER_H
#define ROUTE_HANDLER_H

#include <stddef.h>

// Функция выполняет преобразование HTTP-запроса.
// Если в первой строке URL начинается с "/api/v1", она заменяется на "/micro".
void transform_request(const char *orig_req, char *new_req, size_t new_req_size);

#endif
