#ifndef SERVER_H
#define SERVER_H

#ifdef __cplusplus
extern "C" {
#endif

// Функция запуска сервера, принимает адрес и порт для прослушивания
void start_server(const char* listen_addr, const char* listen_port);

#ifdef __cplusplus
}
#endif

#endif
