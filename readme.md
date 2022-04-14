# Otus Highload Architect homework 10

СЕрвис счетчиков

## Run

    docker-compose up // запуск БД

    cp .env.example .env

    ./main

## .env configuration

     DB_HOST=127.0.0.1
     DB_PORT=33161
     DB_USER=otus
     DB_PASSWORD=otus
     DB_NAME=otus
     HTTP_PORT=7001


## Описание

### Функции системы

- Изменение счетчиков пользователя:

  Пример запроса:

      POST http://localhost:7001/api/counters/1
      
      {
          "UnreadMessagesCountDelta": 1
      }

  Пример ответа:
      
       {
          "UnreadMessagesCount": 1,
          "UserId": 1
       }      

- Чтение счетчиков пользователя:
   
   Пример запроса:

       GET http://localhost:7001/api/counters/1    

   Пример ответа:

       {
          "UnreadMessagesCount": 0,
          "UserId": 1
       }

