# Тестовое задание для KODE (REST API сервис)

### Особенности:
- REST API интерфейс, для передачи используется `json`
- Для логирования используется библиотека [logrus](https://github.com/sirupsen/logrus)
- Для роутера используется [gorilla](https://github.com/gorilla/mux)
- Live-reload с помощью [air](https://github.com/air-verse/air)
- В качестве БД используется Postgresql (запускается в контейнере)
- Переменные окружения хранятся в файле `.env`
- Добавляемые заметки проверяются на орфографические ошибки сервисом [Яндекс.Спеллер](https://yandex.ru/dev/speller)
- web-сервер запускается на порту `:8080`, бд - на порту `:5432`
- Для аутентификации используется `X-Session-Token`

### Тестирование через Postman:
В файле, расположенном docs содержится коллекция Postman

### Принимаемые X-Session-Token (пока что хардкод):
- Alex: wn29ANSMq39
- Jane: k01NSmp2WQA
- Kate: 14mAf93Aq1f

### Endpoints:
- POST /note - добавить заметку
  - Тело запроса:
    - note - текст заметки
- GET /note - вывести все заметки

### Запуск сервиса:
```
docker compose build
docker compose up
```
