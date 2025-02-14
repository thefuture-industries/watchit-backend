#ifndef COMMON_H
#define COMMON_H

#include <pthread.h>

// Глобальный мьютекс для синхронизации round-robin выбора backend-сервера
extern pthread_mutex_t rr_mutex;

#endif
