/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#ifndef LOGGER_H
#define LOGGER_H

#ifdef _WIN32
#include <direct.h>  // _mkdir
#else
#include <sys/stat.h>
#endif

// Записывает данные о запросе в лог-файл
void log_request(const char *device, const char *method, const char *url, const char *status, const char *message);

// Записывает сообщение об ошибке в лог-файл
void log_error(const char *message);

// Записывает сообщение об безопастности в лог-файл
void log_security(const char *message);

#endif
