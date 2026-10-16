# short_link_increment7

# Задание по треку «Сервис сокращения URL»

Добавьте в код сервера новый эндпоинт `POST /api/shorten`, который будет принимать в теле запроса JSON-объект `{"url":"<some_url>"}` и возвращать в ответ объект `{"result":"<short_url>"}`.

## Пример запроса

```http
POST http://localhost:8080/api/shorten HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "url": "https://practicum.yandex.ru"
}
```

## Пример ответа

```http
HTTP/1.1 201 OK
Content-Type: application/json
Content-Length: 30

{
  "result": "http://localhost:8080/EwHXdJfB"
}
```

## Тесты

Не забудьте добавить тесты на новый эндпоинт, как и на предыдущие.

## Библиотеки

При реализации задействуйте одну из распространённых библиотек:

- [`encoding/json`](https://pkg.go.dev/encoding/json)
- [`github.com/mailru/easyjson`](https://github.com/mailru/easyjson)
- [`github.com/pquerna/ffjson`](https://github.com/pquerna/ffjson)
