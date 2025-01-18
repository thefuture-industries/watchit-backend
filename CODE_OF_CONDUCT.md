-----ENGLISH VERSION-----

# Protocol of operation and interaction with the application

1. Architecture

- server/ a server for processing, working with the database, and working with the business logic of the application.
- client-mobile/ folder for working with android/iOS platforms.
- client-windows/ folder for working with windows platforms.

2. Starting the server

- To start the server, go to the directory

```bash
cd server
```

- Launching the server

- To start the server, install golang before starting - [golang](https://go.dev/dl /)
- And select go1.23.5.windows-386.msi > maybe newer (but the name should be something like that)

```bash
go run main.go
```

3. Launching the Android app

- To launch the application, go to the directory

```bash
cd client-mobile
```

```bash
npx expo start
```

4. Launching the Windows application

- To launch the application, go to the directory

```bash
cd client-windows
```

```bash
dotnet run
```

5. Making changes/tracking

## IMPORTANT

- After making changes to the code, upload them to github in a new branch! (How? see below)
- After making the changes, write them down in the file CHANGELOG.md (English and Russian versions) (how? see the example in the file CHANGELOG.md)

## GitHub changes

- Go to the main catologist it will be (flicksfi/) And follow the steps in strict order.

```bash
git add .

git commit -m "the name of the changes (in English), if done on issues before, indicate the number in parentheses - example below"

git commit -m "added functions (#32)"

git push origin v0.1.2
```

-----RUSSIAN VERSION-----

# Протокол работы и взаимодействия с приложением

1. Архитектура

- server/ сервер для обработки, работы с БД, работа с бизнес логикой приложения.
- client-mobile/ папка для работы с android/ios платформами.
- client-windows/ папка для работы с windows платформами.

2. Запуск сервера

- Для запуска сервера перейдите в каталог

  ```bash
  cd server
  ```

- Запуск сервера

  - Для запуска сервера перед началом установите golang - [golang](https://go.dev/dl/)
  - И выберите go1.23.5.windows-386.msi > может быть новее(но название должено быть примерно таким)

  ```bash
  go run main.go
  ```

3. Запуск приложения для Android

- Для запуска приложения перейдите в каталог

  ```bash
  cd client-mobile
  ```

  ```bash
  npx expo start
  ```

4. Запуск приложения для Windows

- Для запуска приложения перейдите в каталог

  ```bash
  cd client-windows
  ```

  ```bash
  dotnet run
  ```

5. Изменения внесения/отслеживание

## ВАЖНО

- После внесения изменения в код кидайте их в github в новую ветку! (как? смотреть ниже)
- После внесения изменения распишите их в файле CHANGELOG.md (английская и русская версия) (как? смотреть пример в файле CHANGELOG.md)

## GitHUB изменения

- Перейдите в главный католог это будет (flicksfi/) И выполняйте действия в строгом порядке.

```bash
git add . (если высветились сообщения ниже то еще раз ниписать git add .)

git commit -m "название изменений(на английском), если делали по issues до в скобках указать номер - пример ниже"

git commit -m "added functions (#32)"

git push origin v0.1.2
```
