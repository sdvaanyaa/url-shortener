# 🌐 URL Shortener 

Сервис для создания коротких ссылок на Go. Поддерживает сохранение, редирект и удаление ссылок с авторизацией и SQLite-хранилищем.

## Основные возможности
- **Сокращение ссылок**: Генерация коротких алиасов для URL (**POST**).
- **Редирект**: Перенаправление по коротким ссылкам (**GET**).
- **Удаление**: Удаление ссылок по алиасу (**DELETE**).
- **Авторизация**: Basic Auth для защищённых операций.
- **Хранилище**: SQLite.
- **Логирование**: Настраиваемые уровни логов с цветным выводом.

## Технологии
- **Роутинг**: `github.com/go-chi/chi/v5`
- **Хранилище**: `github.com/mattn/go-sqlite3`
- **Конфигурация**: `github.com/ilyakaznacheev/cleanenv`
- **Логи**: `log/slog` с кастомными обработчиками
- **Валидация**: `github.com/go-playground/validator/v10`

## Структура проекта
- `cmd/url-shortener/main.go` — точка входа.
- `internal/config/` — загрузка конфигурации из YAML.
- `internal/http-server/` — HTTP-обработчики и middleware.
- `internal/storage/` — взаимодействие с SQLite.
- `internal/lib/` — утилиты (логи, API, генерация строк).

## 🌐 API Примеры

### Создание ссылки
**POST** `http://localhost:8082/url`  
**Аутентификация:** Basic Auth  
**Пример запроса:**
```json
{
  "url": "https://music.yandex.ru/",
  "alias": "yandmusic"
}
``` 
* Если alias не указан, он будет сгенерирован автоматически
 
**Пример ответа:**
```json
{
  "status": "OK",
  "alias": "yandmusic"  
}
```   
  
**Запрос с помощью curl:**
```shell
curl -X POST "http://localhost:8082/url" \
    -u username:password \
    -H "Content-Type: application/json" \
    -d '{"url": "https://music.yandex.ru/", "alias": "yandmusic"}'
```
<br>

### Редирект по ссылке
**GET** `http://localhost:8082/yandmusic`  

**Запрос с помощью curl:**
```shell
curl -L "http://localhost:8082/yandmusic"
```  
  
**Редирект на `https://music.yandex.ru/`**

<br>

### Удаление ссылки

**DELETE** `http://localhost:8082/yandmusic`  
**Аутентификация:** Basic Auth  

**Запрос с помощью curl:**
```shell
curl -X DELETE "http://localhost:8082/yandmusic" \
    -u username:password
```

**Пример ответа:**
```json
{
  "status": "OK"
}
```