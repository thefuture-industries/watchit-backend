# Flicksfi - Backend v2

## Дата релиза второй версии приложения `28.12.2025`

> Первая версия сервера v1 храниться в папке server-ooold, сейчас создана папка server/ это вторая v2 версия сервера.

Backend написан на микросервисной архитектуре (основной язык Golang).

user-service/ - Регистрация/Вход/Обновление/Подписка/Профиль
movie-service/ - Рекомендации/Фильмы/Избранные/Лимиты
blog-service/ - Комментарии/Лайки к фильмам/Подписчики?

> Цель: сделать быстрый сервис для нахождения фильмов когда нечего посмотреть, чтобы пользователь зашел и сразу нашел свой фильм фаворит.

## База данных фильмов

Данные хранятся в виде json файла но при продакшене используется .gz вид.

Пример хранения одного фильма в json:

```json
{
    "adult": false,
    "backdrop_path": "/gMQibswELoKmB60imE7WFMlCuqY.jpg",
    "genre_ids": [27, 53, 9648],
    "id": 1034541,
    "original_language": "en",
    "original_title": "Terrifier 3",
    "overview": "Five years after surviving Art the Clown's Halloween massacre, Sienna and Jonathan are still struggling to rebuild their shattered lives. As the holiday season approaches, they try to embrace the Christmas spirit and leave the horrors of the past behind. But just when they think they're safe, Art returns, determined to turn their holiday cheer into a new nightmare. The festive season quickly unravels as Art unleashes his twisted brand of terror, proving that no holiday is safe.",
    "popularity": 4184.074,
    "poster_path": "/63xYQj1BwRFielxsBDXvHIJyXVm.jpg",
    "release_date": "2024-10-09",
    "title": "Terrifier 3",
    "video": false,
    "vote_average": 7.201,
    "vote_count": 698
    },
```

## Акцент

Текущий акцент на улучшенный алгоритм [fom](https://github.com/thefuture-industries/flicksfi-backend/issues/188)

## Разработка

1. Для быстрого вхождения в разработку прочти CONTRIBUTING.md
2. Для быстрой разработки рекомендуем использовать [git auto-commit](https://github.com/thefuture-industries/git-auto-commit)

## Перед началом разработке

1. Запустите script создания .env файлов:

```bash
cd scripts
bash create_envs
```

2. Установите pre-commit

```bash
npx husky install
```

3. Установите auto-commit через [https://github.com/thefuture-industries/git-auto-commit]
