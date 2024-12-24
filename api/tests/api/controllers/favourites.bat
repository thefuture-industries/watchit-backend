@echo off
setlocal enabledelayedexpansion

@REM UUIDs пользователей
set uuids=51ee2a52-0618-4f94-bba8-13c433b4a164 e53eb127-df5e-4dfb-86e0-89c9805123c5
@REM Кол-во запросов
set requests_count=10000
@REM Кол-во одновременных запросов
set concurrent_requests=40
@REM URL сервера
set base_url=http://localhost:8080/api/v1/favourites/

set method=%1

if /I "%method%"=="get" (
  goto:get_favourites
) else (
  echo Method not found
)

@REM Функция получение избранных пользователя
:get_favourites
  @REM Рандомное uuid
  set /A random_uuid=!RANDOM! %% 2

  set "selected_uuid="
  for %%u in (%uuids%) do (
    if !random_uuid!==0 (
      set "selected_uuid=%%u"
    )
    set random_uuid=-1
  )

  @REM Отпарвка запроса
  set "url=%base_url%%selected_uuid%"
  echo Sending requests %requests_count% on the url: !url!
  hey -n %requests_count% -c %concurrent_requests% -m GET !url!
  echo Completed sending requests for !selected_uuid!
  echo .
goto:eof

endlocal
pause
