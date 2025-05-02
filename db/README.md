# Структура базы данных
[<img src="https://raw.githubusercontent.com/Teru3301/FinancialManager/dev/db/other/plan.png"/>]()

# Запуск модуля:
## Перед запуском убедитесь что у вас установлены:
- PostgreSQL 
- Go-lang версии 1.24.2 или выше
## Запуск
- Перейдите в директорию модуля
```bash
cd FinancialManager/db
```
- Запустите модуль
```go
go run .
```

При отсутствии пользователя или базы данных, будет выполняться попытка их создания. Если они есть, выполнится попытка подключения.
После успешного подключения к БД и пользователю, запустится регистрация API модуля, затем модуль начнёт прослушивать ```3000``` порт.


### Тестовые эндпоинты
- http://localhost:3000/database/goals?uid=1
- http://localhost:3000/database/groups?uid=1
- http://localhost:3000/database/transactions?uid=1

# Предоставляемое модулем API

### Каждый запрос должен содержать в заголовке JWT токен

## Транзакции

| Метод | Эндпоинт               | Описание                                             | Параметры в URL | Принимаемый JSON | Возвращаемый JSON |
| ----- | ---------------------- | ---------------------------------------------------- | --------------- | ---------------- | --- |
| GET   | /database/transactions | Возвращает список всех транзакций пользователя       | uid             | -                | [tid, category, date_time, description, money, type] |
| POST  | /database/transaction  | Добавляет новую транзакцию                           | uid             | money, date_time, type, description, category | - |
| PUT   | /database/transaction  | Изменение транзакции                                 | uid, tid        | money, date_time, type, description, category | - |
| Delete| /database/transaction  | Удаление транзакции                                  | uid, tid        | -                | - |

## Запросы для тестов

#### POST
- добавление транзакции
```bash
curl -X POST http://localhost:3000/database/transaction\?uid\=1 \
  -H "Content-Type: application/json" \
  -d '{
    "money": 30.5,
    "type": "income",
    "category": "степендия",
    "description": "...",
    "date_time": "2026-05-02T18:45:00"
  }'
```
#### PUT
- обновление транзакции
```bash
curl -X PUT http://localhost:3000/database/transaction\?uid\=1\&tid\=1 \
  -H "Content-Type: application/json" \
  -d '{
    "money": 3500.75,
    "type": "expense",
    "category": "учеба",
    "description": "оплата учебных материалов",
    "date_time": "2026-06-01T12:30:00"
  }'
```
#### DELETE
- удаление транзакции
```bash
curl -X DELETE http://localhost:3000/database/transaction\?uid\=1\&tid\=1 
```
#### GET 
- получение списка транзакций
```bash
curl http://localhost:3000/database/transactions\?uid\=1
```
