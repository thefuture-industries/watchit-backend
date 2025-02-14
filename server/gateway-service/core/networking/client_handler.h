#ifndef CLIENT_HANDLER_H
#define CLIENT_HANDLER_H

#ifdef __cplusplus
extern "C" {
#endif

// Функция-обработчик клиента (функция потока)
void *handle_client(void *arg);

// Функция пересылки запроса на backend-сервер
int forward_to_backend(int client_sock, const char *backend_ip, int backend_port, const char *method, const char *path, const char *request);

#ifdef __cplusplus
}
#endif

#endif
