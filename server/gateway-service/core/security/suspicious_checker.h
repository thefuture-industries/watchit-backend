/*---------------------------------------------------------------------------------------------
 *  Copyright (c). All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

#ifndef SUSPICIOUS_CHECKER_H
#define SUSPICIOUS_CHECKER_H

// Функция возвращает 1, если в строке request найден подозрительный фрагмент,
// и 0 — если подозрительных элементов не найдено.
int is_request_suspicious(const char* request);

#endif
