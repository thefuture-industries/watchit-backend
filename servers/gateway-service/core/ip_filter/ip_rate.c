/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include "ip.h"
#include "../include/uthash.h"
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <pthread.h>

typedef struct ip_entry {
    // ключ ip address
    char ip[INET_ADDRSTRLEN];
    // кол-во запросов в секунду
    int count;
    // время последнего обновления счетчика
    time_t last_update;
    // флаг блокировки
    int banned;
    // для uthash
    UT_hash_handle hh;
} ip_entry;

static ip_entry *ip_table = NULL;
static pthread_mutex_t ip_mutex = PTHREAD_MUTEX_INITIALIZER;

int check_ip_rate(const char *clientIP) {
    time_t current_time = time(NULL);
    int is_banned = 0;

    pthread_mutex_lock(&ip_mutex);

    ip_entry *entry = NULL;
    HASH_FIND_STR(ip_table, clientIP, entry);
    if (entry == NULL) {
        // Если записи нет — создаём новую
        entry = malloc(sizeof(ip_entry));
        if (entry != NULL) {
            strncpy(entry->ip, clientIP, INET_ADDRSTRLEN);
            entry->ip[INET_ADDRSTRLEN - 1] = '\0';
            entry->count = 1;
            entry->last_update = current_time;
            entry->banned = 0;
            HASH_ADD_STR(ip_table, ip, entry);
        }
    }
    else {
        // Если пришёл запрос в новой секунде — сбрасываем счётчик
        if (entry->last_update != current_time) {
            entry->count = 1;
            entry->last_update = current_time;
            entry->banned = 0;
        }
        else {
            entry->count++;
            if (entry->count >= 20) {
                entry->banned = 1;
            }
        }
    }

    if (entry != NULL){
        is_banned = entry->banned;
    }

    pthread_mutex_unlock(&ip_mutex);
    return is_banned;
}
