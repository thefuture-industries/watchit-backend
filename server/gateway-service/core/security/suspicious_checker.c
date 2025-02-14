/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include "suspicious_checker.h"
#include <string.h>

#define MAX_BANNED_WORDS 128
#define MAX_WORD_LENGTH 128

typedef struct {
    char words[MAX_BANNED_WORDS][MAX_WORD_LENGTH];
    int count;
} BannedWordsList;

bool add_banned_word(BannedWordsList *list, const char *word) {
    if (list->count < MAX_BANNED_WORDS) {
        strncpy(list->words[list->count], word, MAX_WORD_LENGTH - 1);
        list->words[list->count][MAX_WORD_LENGTH - 1] = '\0';
        list->count++;
        return true;
    }
    else {
        fprintf(stderr, "Error: Banned words list is full.\n");
        return false;
    }
}

bool contains_banned_word(const BannedWordsList *list, const char *text) {
    for (int i = 0; i < list->count; i++) {
        if (strstr(text, list->words[i]) != NULL) {
            printf(list->words[i]);
            return true;
        }
    }
    return false;
}

void initialize_banned_words_list(BannedWordsList *list) {
    list->count = 0;
}

int is_request_suspicious(const char* request) {
    BannedWordsList banned_words;
    initialize_banned_words_list(&banned_words);

    // XSS
    add_banned_word(&banned_words, "<script");
    add_banned_word(&banned_words, "</script");
    add_banned_word(&banned_words, "<iframe");
    add_banned_word(&banned_words, "</iframe");
    add_banned_word(&banned_words, "javascript:");
    add_banned_word(&banned_words, "vbscript:");
    add_banned_word(&banned_words, "data:text/html");
    add_banned_word(&banned_words, "onclick=");
    add_banned_word(&banned_words, "onerror=");
    add_banned_word(&banned_words, "onmouseover=");
    add_banned_word(&banned_words, "onfocus=");
    add_banned_word(&banned_words, "eval(");
    add_banned_word(&banned_words, "expression(");
    add_banned_word(&banned_words, "document.cookie");
    add_banned_word(&banned_words, "window.onload");
    add_banned_word(&banned_words, "setTimeout(");
    add_banned_word(&banned_words, "setInterval(");
    add_banned_word(&banned_words, "ActiveXObject(");
    add_banned_word(&banned_words, "appendChild(");
    add_banned_word(&banned_words, "fromCharCode(");
    add_banned_word(&banned_words, "base64,");
    add_banned_word(&banned_words, "fromCharCode");
    add_banned_word(&banned_words, "location.hash");
    add_banned_word(&banned_words, "location.replace");
    add_banned_word(&banned_words, "unescape");
    add_banned_word(&banned_words, "%3cscript");

    // SQL Injection
    add_banned_word(&banned_words, "union select");
    add_banned_word(&banned_words, "union all select");
    add_banned_word(&banned_words, "information_schema");
    add_banned_word(&banned_words, "xp_cmdshell");
    add_banned_word(&banned_words, "exec(");
    add_banned_word(&banned_words, "INSERT INTO");
    add_banned_word(&banned_words, "DELETE FROM");
    add_banned_word(&banned_words, "UPDATE ");
    add_banned_word(&banned_words, "CREATE TABLE");
    add_banned_word(&banned_words, "DROP TABLE");
    add_banned_word(&banned_words, "TRUNCATE TABLE");
    add_banned_word(&banned_words, "ALTER TABLE");
    add_banned_word(&banned_words, "sleep(");
    add_banned_word(&banned_words, "benchmark(");

    // LFI/RFI (Local/Remote File Inclusion)
    add_banned_word(&banned_words, "../");
    add_banned_word(&banned_words, "..\\");
    add_banned_word(&banned_words, "file://");
    add_banned_word(&banned_words, "http://");
    add_banned_word(&banned_words, "https://");
    add_banned_word(&banned_words, "%2e%2e%2f");
    add_banned_word(&banned_words, "%2e%2e%5c");
    add_banned_word(&banned_words, "php://");
    add_banned_word(&banned_words, "data://");

    // OS Command Injection
    add_banned_word(&banned_words, "system(");
    add_banned_word(&banned_words, "passthru(");
    add_banned_word(&banned_words, "shell_exec(");
    add_banned_word(&banned_words, "exec(");
    add_banned_word(&banned_words, "popen(");
    add_banned_word(&banned_words, "proc_open(");

    // Sensitive Data
    add_banned_word(&banned_words, "passwd");
    add_banned_word(&banned_words, "api_key");
    add_banned_word(&banned_words, "secret_key");
    add_banned_word(&banned_words, "private_key");
    add_banned_word(&banned_words, "oauth_token");
    add_banned_word(&banned_words, "credit_card");
    add_banned_word(&banned_words, "ssn"); // Social Security Number
    add_banned_word(&banned_words, "social security number"); // Полная фраза
    add_banned_word(&banned_words, "bank_account");

    // Общие подозрительные вещи
    add_banned_word(&banned_words, "<?php");
    add_banned_word(&banned_words, "<%");
    add_banned_word(&banned_words, "file_get_contents(");
    add_banned_word(&banned_words, "curl_exec(");
    add_banned_word(&banned_words, "$_GET");
    add_banned_word(&banned_words, "$_POST");
    add_banned_word(&banned_words, "$_REQUEST");
    add_banned_word(&banned_words, "getenv(");
    add_banned_word(&banned_words, "require_once(");
    add_banned_word(&banned_words, "include_once(");
    add_banned_word(&banned_words, "require(");
    add_banned_word(&banned_words, "include(");
    add_banned_word(&banned_words, "<?xml");
    add_banned_word(&banned_words, "<!DOCTYPE");

    add_banned_word(&banned_words, "onerror");
    add_banned_word(&banned_words, "onload");
    add_banned_word(&banned_words, "iframe");
    add_banned_word(&banned_words, "onclick=");
    add_banned_word(&banned_words, "onerror=");
    add_banned_word(&banned_words, "javascript:");
    add_banned_word(&banned_words, "vbscript:");
    add_banned_word(&banned_words, "%3Cscript%3E");
    add_banned_word(&banned_words, "%3Ciframe%3E");
    add_banned_word(&banned_words, "data:");
    add_banned_word(&banned_words, "expression(");
    add_banned_word(&banned_words, "</script>");
    add_banned_word(&banned_words, "url(");
    add_banned_word(&banned_words, "prompt(");


    //Примеры обхода - их тоже нужно учитывать.
    add_banned_word(&banned_words, "String.fromCharCode");
    add_banned_word(&banned_words, "unescape");
    add_banned_word(&banned_words, "decodeURIComponent");
    add_banned_word(&banned_words, "charCodeAt");

    if (request == NULL) {
        return 0;
    }

    if (contains_banned_word(&banned_words, request)) {
        return 1;
    }

    return 0;
}
