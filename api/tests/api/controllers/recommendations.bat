@echo off
setlocal enabledelayedexpansion

@REM Переменные
@REM URL сервера
set base_url=http://localhost:8080/api/v1/recommendations/
@REM Массив UUIDs
set uuids=51ee2a52-0618-4f94-bba8-13c433b4a164 e53eb127-df5e-4dfb-86e0-89c9805123c5

@REM Получаем методы нагрузки (get, post, put, delete)
set method=%1

@REM Проверка метода
if /I "%method%"=="get" (
    goto:get_recommendations
) else if /I "%method%"=="post" (
    goto:post_recommendations
) else (
    echo Method not found
)

@REM Функция получение рекомендация пользователя
@REM /recommendations/{uuid}
:get_recommendations
    @REM Кол-во запросов
    set requests_per_uuid=1000

    @REM Кол-во одновременных запросов
    set concurrent_requests=20

    @REM Проход по каждому UUID
    for %%u in (%uuids%) do (
        set "url=%base_url%%%u"
        echo Sending requests %requests_per_uuid% on the url: !url!

        rem Отправка запросов с помощью hey
        hey -n %requests_per_uuid% -c %concurrent_requests% -m GET !url!

        echo Completed sending requests for %%u
        echo .
    )
goto:eof

@REM Функция добавление рекомендаций пользователя
@REM /recommendations
:post_recommendations
    @REM Кол-во запросов
    set requests_count=1000

    @REM Кол-во одновременных запросов
    set concurrent_requests=20

    @REM Массив жанров
    set genres=28 12 16 35 80

    @REM Получение кол-ва жанров
    set /A length_genres=0
    for %%g in (%genres%) do (
        set /A length_genres+=1
    )

    for /L %%i in (1, 1, %requests_count%) do (
        @REM Случайные числа
        set /A random_index_genre=!RANDOM! %% !length_genres!
        set /A random_index_uuid=!RANDOM! %% 2

        @REM Выбор случайного UUID
        set "selected_uuid="
        for %%u in (%uuids%) do (
            if !random_index_uuid!==0 (
                set selected_uuid=%%u
            )
            set /A random_index_uuid-=1
        )
    
        @REM Выбор случайного жанра
        set "selected_genre="
        for %%g in (%genres%) do (
            if !random_index_genre!==0 (
                set selected_genre=%%g
            )
            set /A random_index_genre-=1
        )

        @REM Тело запроса
        set body={uuid:!selected_uuid!,title:Title %%i,genre:!selected_genre!}

        echo Sending request: !body!

        @REM Отправка запроса
        curl.exe -H "Content-Type: application/json" -d "!body!" -X POST !base_url!
    )
goto:eof

endlocal
pause
