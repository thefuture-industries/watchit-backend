#ifndef SERVICES_CONFIG_H
#define SERVICES_CONFIG_H

typedef struct {
    char ip[16];
    int port;
} Backend;

// Настройки для user-service
Backend user_backends[] = {
    {"127.0.0.1", 8001},
    {"127.0.0.1", 8002},
};
int user_backend_count = sizeof(user_backends) / sizeof(user_backends[0]);
int user_rr_index = 0;

// Настройки для blog-service
Backend blog_backends[] = {
    {"127.0.0.1", 8011},
    {"127.0.0.1", 8011},
};
int blog_backend_count = sizeof(blog_backends) / sizeof(blog_backends[0]);
int blog_rr_index = 0;

// Настройки для movie-service
Backend movie_backends[] = {
    {"127.0.0.1", 8020},
    {"127.0.0.1", 8021},
};
int movie_backend_count = sizeof(movie_backends) / sizeof(movie_backends[0]);
int movie_rr_index = 0;

#endif
