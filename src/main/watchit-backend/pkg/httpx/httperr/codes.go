package httperr

import "errors"

var (
	Err_DuplicateEmail          = errors.New("пользователь с таким адресом электронной почты уже существует")
	Err_UserNotFound            = errors.New("пользователь не был найден")
	Err_ContextDeadlineExceeded = errors.New("долгое время ожидания ответа от базы данных, повторите попытку позже")
	Err_ContextCanceled         = errors.New("операция была отменена")
	Err_UniqueViolation         = errors.New("введенные данные уже существует")
	Err_DbTimeout               = errors.New("соединение с базой данных превысило время ожидания")
	Err_DbNetworkTemporary      = errors.New("временная проблема с сетью. попробуйте ещё раз")
	Err_DbNetwork               = errors.New("сетевая ошибка при подключении к базе данных")
	Err_NotDeleted              = errors.New("ошибка, ни одна запись не была удалена")
	Err_NotUpdated              = errors.New("ошибка, ни одна запись не была обновлена")
)
