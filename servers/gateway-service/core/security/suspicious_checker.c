/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include "suspicious_checker.h"
#include <string.h>

#ifdef _WIN32
#include <ctype.h>
// Реализация strcasestr для Windows
static char *strcasestr(const char *haystack, const char *needle) {
    if (!*needle)
        return (char *)haystack;

    for (; *haystack; haystack++) {
        if (tolower((unsigned char)*haystack) == tolower((unsigned char)*needle)) {
            const char *h = haystack;
            const char *n = needle;
            while (*h && *n && tolower((unsigned char)*h) == tolower((unsigned char)*n)) {
                h++;
                n++;
            }
            if (!*n)
                return (char *)haystack;
        }
    }
    return NULL;
}
#endif

static const char *suspicious_keywords[] = {
    "DROP",
    "DELETE",
    "INSERT",
    "SELECT",
    "<script",
    "onerror",
    "onload",
    "</script>",
    NULL
};

int is_request_suspicious(const char* request) {
    if (request == NULL) {
        return 0;
    }

    for (int i = 0; suspicious_keywords[i] != NULL; i++) {
    if (strcasestr(request, suspicious_keywords[i]) != NULL)
        return 1;
    }

    return 0;
}
