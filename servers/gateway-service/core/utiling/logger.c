/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include "logger.h"
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>

#define LOG_DIR "logs"
#define LOG_FILE "logs/server.log"
#define ERROR_LOG_FILE "logs/error.log"

// Функция создаёт папку, если её не существует
static void ensure_log_dir_exists() {
#ifdef _WIN32
    _mkdir(LOG_DIR);
#else
    mkdir(LOG_DIR, 0777);
#endif
}

void log_request(const char *device, const char *method, const char *url, const char *status, const char *message) {
    ensure_log_dir_exists();
    FILE *fp = fopen(LOG_FILE, "a");
    if (!fp) return;

    time_t now = time(NULL);
    struct tm *tm_info = localtime(&now);
    char time_str[64];
    strftime(time_str, sizeof(time_str), "%Y-%m-%d %H:%M:%S", tm_info);

    fprintf(fp, "[%s] Device: %s | Method: %s | URL: %s | Status: %s | Message: %s",
            time_str,
            device,
            method,
            url,
            status,
            message);
    fclose(fp);
}

void log_error(const char *message) {
    ensure_log_dir_exists();
    FILE *fp = fopen(ERROR_LOG_FILE, "a");
    if (!fp) return;

    time_t now = time(NULL);
    struct tm *tm_info = localtime(&now);
    char time_str[64];
    strftime(time_str, sizeof(time_str), "%Y-%m-%d %H:%M:%S", tm_info);
    fprintf(fp, "[%s] ERROR: %s\n", time_str, message);
    fclose(fp);
}
