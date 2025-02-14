/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>

#define BUF_SIZE 4096

void add_header(char *buffer, const char *header) {
    size_t buf_len = strlen(buffer);
    size_t hdr_len = strlen(header);
    if (buf_len + hdr_len + 1 < BUF_SIZE) {
        memmove(buffer + hdr_len, buffer, buf_len + 1);
        memcpy(buffer, header, hdr_len);
    }
}
