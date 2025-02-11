#include "suspicious_checker.h"
#include <string.h>

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
