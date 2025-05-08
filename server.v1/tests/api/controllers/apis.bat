@echo off
setlocal enabledelayedexpansion

@REM Переменные
@REM URL сервера
set base_url=http://localhost:8080/api/v1/
@REM Массив UUIDs
set uuids=51ee2a52-0618-4f94-bba8-13c433b4a164 e53eb127-df5e-4dfb-86e0-89c9805123c5
@REM Кол-во запросов
set requests_count=500
@REM Кол-во одновременных запросов
set concurrent_requests=10
@REM Потоки запросов
set threads=3

@REM Получаем router запроса (similar, popular, etc.)
set router=%1

if /I "%router%"=="similar" (
  goto:similar
) else if /I "%router%"=="details" (
  goto:details
) else if /I "%router%" == "text" (
  goto:text
) else (
  echo Method not found
)

@REM Фукнция для получения похожих фильмов
@REM /movies/similar
:similar
  @REM Массив жанров
  set genres=28 12 16 35 80

  @REM Массив описаний фильмов
  set "description0=Five years after surviving Art the Clown's Halloween massacre, Sienna and Jonathan are still struggling to rebuild their shattered lives. As the holiday season approaches, they try to embrace the Christmas spirit and leave the horrors of the past behind. But just when they think they're safe, Art returns, determined to turn their holiday cheer into a new nightmare. The festive season quickly unravels as Art unleashes his twisted brand of terror, proving that no holiday is safe."
  set "description1=Eddie and Venom are on the run. Hunted by both of their worlds and with the net closing in, the duo are forced into a devastating decision that will bring the curtains down on Venom and Eddie's last dance."
  set "description2=A fading celebrity decides to use a black market drug, a cell-replicating substance that temporarily creates a younger, better version of herself."
  set "description3=While struggling with his dual identity, Arthur Fleck not only stumbles upon true love, but also finds the music that's always been inside him."
  set "description4=A listless Wade Wilson toils away in civilian life with his days as the morally flexible mercenary, Deadpool, behind him. But when his homeworld faces an existential threat, Wade must reluctantly suit-up again with an even more reluctant Wolverine."
  set descriptions=!description0! !description1! !description2! !description3! !description4!

  @REM Получение кол-ва жанров
  set /A length_genres=0
  for %%g in (%genres%) do (
    set /A length_genres+=1
  )

  @REM Получение кол-во описаний
  set /A length_descriptions=0
  for %%d in (%descriptions%) do (
    set /A length_descriptions+=1
  )

  @REM Случайные числа
  set /A random_index_genre=!RANDOM! %% !length_genres!
  set /A random_index_description=!RANDOM! %% 4

  @REM Выбор случайного жанра
  set "selected_genre="
  for %%g in (%genres%) do (
    if !random_index_genre!==0 (
      set selected_genre=%%g
    )
    set /A random_index_genre-=1
  )

  @REM Выбор случайного описания
  set "selected_description=!description%random_index_description%!"
  @REM Замена пробелов на %20
  set "encoded_description=!selected_description: =%%20!"

  @REM Формирование URL запроса
  set "url=%base_url%movies/similar?genre=!selected_genre!&overview=!encoded_description!"

  @REM Выполнение запроса
  echo Sending requests %requests_count% on the url: !url!
  hey -n %requests_count% -c %concurrent_requests% -m GET !url!
  echo Completed sending requests for similar movies
  echo .
goto:eof

@REM Функция для получения деталий фильма
@REM /movie/{id}
:details
  @REM Массив ID
  set "ids=1184918 933260 698687 889737 533535 1084736 1051896 335983 945961 580489 845781"

  @REM Получение кол-ва ID
  set /A ids_length=0
  for %%i in (%ids%) do (
    set /A ids_length+=1
  )

  @REM Случайный индекс
  set /A random_index_id=!RANDOM! %% !ids_length!

  @REM Выбор случайного ID
  set "selected_id="
  for %%j in (%ids%) do (
    if !random_index_id!==0 (
      set selected_id=%%j
    )
    set /A random_index_id-=1
  )

  @REM Формирование URL запроса
  set "url=%base_url%movie/!selected_id!"

  @REM Выполнение запроса
  echo Sending requests %requests_count% on the url: !url!
  hey -n %requests_count% -c %concurrent_requests% -m GET !url!
  echo Completed sending requests for movie details
  echo .
goto:eof

@REM Функция для получения фильмов по тексту
@REM /text/movies
:text
  @REM Текст для поиска
  set "text=This is test"
  set "encoded_text=!text: =%%20!"
  @REM В недалёком будущем человечество сталкивается с глобальным кризисом ресурсов. Земля истощена, и люди вынуждены искать новые источники энергии и сырья. Ученые разрабатывают технологию межзвездных путешествий, но ресурсы для реализации проекта ограничены.

  @REM Случайный UUID
  set "uuid="
  set /A random_index_uuid=%RANDOM% %% 2
  for %%u in (%uuids%) do (
    if !random_index_uuid!==0 (
      set uuid=%%u
    )
    set /A random_index_uuid-=1
  )

  @REM Формирование URL запроса
  set "url=%base_url%text/movies"

  REM Формирование тела запроса
  set body={uuid:!uuid!,text:!encoded_text!,lege:simple}
  echo UUID: !uuid!
  echo Encoded Text: !encoded_text!
  echo Body: !body!

  @REM Выполнение запроса
  echo Sending requests %requests_count% on the url: !url!
  curl.exe -H "Content-Type: application/json" -d "!body!" -X POST !url!
  echo Completed sending requests for text movies
  echo .
goto:eof

endlocal
pause
