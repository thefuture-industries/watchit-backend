/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#ifndef SERVICES_PATHS_H
#define SERVICES_PATHS_H

typedef struct {
    char ip[16];
    int port;
} Backend;

extern int movie_rr_index;
extern int movie_backend_count;
extern Backend* movie_backends;

extern int blog_rr_index;
extern int blog_backend_count;
extern Backend* blog_backends;

extern int user_rr_index;
extern int user_backend_count;
extern Backend* user_backends;

#endif
