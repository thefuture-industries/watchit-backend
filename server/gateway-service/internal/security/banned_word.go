package security

import (
	"log"
	"strings"
)

const (
	MAX_BANNED_WORDS = 128
	MAX_WORD_LENGTH  = 128
)

type BannedWordsList struct {
	words [MAX_BANNED_WORDS]string
	count int
}

func addBannedWord(list *BannedWordsList, word string) bool {
	if list.count < MAX_BANNED_WORDS {
		list.words[list.count] = word
		list.count++
		return true
	}

	log.Println("Error: Banned words list is full.")
	return false
}

func containsBannedWord(list *BannedWordsList, text string) bool {
	for i := 0; i < list.count; i++ {
		if strings.Contains(text, list.words[i]) {
			log.Println("Found banned word:", list.words[i])
			return true
		}
	}
	return false
}

func initializeBannedWordsList() *BannedWordsList {
	return &BannedWordsList{
		count: 0,
	}
}

func IsRequestSuspicious(request string) int {
	bannedWords := initializeBannedWordsList()

	// Добавляем XSS, SQL инъекции, LFI/RFI и другие опасные фразы
	bannedList := []string{
		"<script>", "</script>", "<iframe>", "</iframe>", "javascript:", "vbscript:",
		"data:text/html", "onclick=", "onerror=", "onmouseover=", "onfocus=", "eval(",
		"expression(", "document.cookie", "window.onload", "setTimeout(", "setInterval(",
		"ActiveXObject(", "appendChild(", "fromCharCode(", "base64,", "fromCharCode",
		"location.hash", "location.replace", "unescape", "%3cscript", "union select",
		"union all select", "information_schema", "xp_cmdshell", "exec(", "INSERT INTO",
		"DELETE FROM", "UPDATE ", "CREATE TABLE", "DROP TABLE", "TRUNCATE TABLE", "ALTER TABLE",
		"sleep(", "benchmark(", "../", "..\\", "file://", "%2e%2e%2f",
		"%2e%2e%5c", "php://", "data://", "system(", "passthru(", "shell_exec(", "exec(", "popen(",
		"proc_open(", "passwd", "api_key", "secret_key", "private_key", "oauth_token", "credit_card",
		"ssn", "social security number", "bank_account", "<?php", "<%", "file_get_contents(",
		"curl_exec(", "$_GET", "$_POST", "$_REQUEST", "getenv(", "require_once(", "include_once(",
		"require(", "include(", "<?xml", "<!DOCTYPE", "onerror", "onload", "iframe", "onclick=",
		"onerror=", "javascript:", "vbscript:", "%3Cscript%3E", "%3Ciframe%3E", "data:", "expression(",
		"</script>", "url(", "prompt(", "String.fromCharCode", "unescape", "decodeURIComponent", "charCodeAt",
	}

	// Добавляем слова в список
	for _, word := range bannedList {
		addBannedWord(bannedWords, word)
	}

	if request == "" {
		return 0
	}

	// Проверяем запрос на наличие запрещенных слов
	if containsBannedWord(bannedWords, request) {
		return 1
	}

	return 0
}
