@echo off
setlocal enabledelayedexpansion

rem Переменные
rem URL сервера
set base_url=http://localhost:8080/api/v1/recommendations/
rem Массив UUIDs
set uuids=51ee2a52-0618-4f94-bba8-13c433b4a164 e53eb127-df5e-4dfb-86e0-89c9805123c5

rem Получаем методы нагрузки (get, post, put, delete)
set method=%1

rem Проверка метода
if /I "%method%" == "get" (
    echo gets
    goto:get_recommendations
) else if /I "%method%" == "post" (
    echo posts
    goto:post_recommendations
) else (
    echo Method not found
)

rem Функция получение рекомендация пользователя
rem /recommendations/{uuid}
:get_recommendations
    rem Кол-во запросов
    set requests_per_uuid=1000

    rem Кол-во одновременных запросов
    set concurrent_requests=20

    rem Проход по каждому UUID
    for %%u in (%uuids%) do (
        set "url=%base_url%%%u"
        echo Sending requests %requests_per_uuid% on the url: !url!

        rem Отправка запросов с помощью hey
        hey -n %requests_per_uuid% -c %concurrent_requests% -m GET !url!

        echo Completed sending requests for %%u
        echo .
    )
goto:eof

rem Функция добавление рекомендаций пользователя
rem /recommendations
:post_recommendations
    rem Кол-во запросов
    set requests_count=1000

    rem Кол-во одновременных запросов
    set concurrent_requests=20

    set uuid0=51ee2a52-0618-4f94-bba8-13c433b4a164
    set uuid1=e53eb127-df5e-4dfb-86e0-89c9805123c5

    rem Массив жанров
    set /A genres0=28
    set /A genres1=12
    set /A genres2=16
    set /A genres3=35
    set /A genres4=80
    set genres=!genres0! !genres1! !genres2! !genres3! !genres4!

    rem Получение кол-ва жанров
    set /A length_genres=0
    for %%g in (%genres%) do (
        set /A length_genres+=1
    )

    for /L %%i in (1, 1, %requests_count%) do (
        rem Случайные числа
        set /A random_index_genre=!RANDOM! %% !length_genres!
        set /A random_index_uuid=!RANDOM! %% 2

        if !random_index_uuid! == 0 (
            set selected_uuid=!uuid0!
        ) else (
            set selected_uuid=!uuid1!
        )

        if !random_index_genre! == 0 (
            set selected_genre=!genres0!
        ) else if !random_index_genre! == 1 (
            set selected_genre=!genres1!
        ) else if !random_index_genre! == 2 (
            set selected_genre=!genres2!
        ) else if !random_index_genre! == 3 (
            set selected_genre=!genres3!
        ) else if !random_index_genre! == 4 (
            set selected_genre=!genres4!
        )

        rem Тело запроса
        set body={uuid:!selected_uuid!,title:Title %%i,genre:!selected_genre!}

        echo Sending request: !body!

        rem Отправка запроса
        curl.exe -H "Content-Type: application/json" -d "!body!" -X POST !base_url!
    )
goto:eof

endlocal
pause
