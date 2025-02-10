/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#ifndef SERVICES_CONFIG_H
#define SERVICES_CONFIG_H

#include "services_paths.h"

// Настройки для user-service
Backend user_backends_uri[] = {
    {"127.0.0.1", 8001},
    {"127.0.0.1", 8002},
};
Backend* user_backends = user_backends_uri;
int user_backend_count = sizeof(user_backends_uri) / sizeof(user_backends_uri[0]);
int user_rr_index = 0;

// Настройки для blog-service
Backend blog_backends_uri[] = {
    {"127.0.0.1", 8011},
    {"127.0.0.1", 8011},
};
Backend* blog_backends = blog_backends_uri;
int blog_backend_count = sizeof(blog_backends_uri) / sizeof(blog_backends_uri[0]);
int blog_rr_index = 0;

// Настройки для movie-service
Backend movie_backends_uri[] = {
    {"127.0.0.1", 8020},
    {"127.0.0.1", 8021},
};
Backend* movie_backends = movie_backends_uri;
int movie_backend_count = sizeof(movie_backends_uri) / sizeof(movie_backends_uri[0]);
int movie_rr_index = 0;

#endif
