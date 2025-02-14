/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

typedef struct conf {
    char* listen_addr;
    char* listen_port;
} conf;

conf config = {
    "0.0.0.0",
    "8080",
};

char* get_config(const char* key) {
    if (strcmp(key, "listen_addr") == 0) {
        return config.listen_addr;
    }
    else if (strcmp(key, "listen_port") == 0) {
        return config.listen_port;
    }

    return NULL;
}
