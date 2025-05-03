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


# Предоставляемое модулем API

### Каждый запрос должен содержать в заголовке JWT токен

## Транзакции

| Метод | Эндпоинт               | Описание                                             | Параметры в URL | Принимаемый JSON | Возвращаемый JSON |
| ----- | ---------------------- | ---------------------------------------------------- | --------------- | ---------------- | --- |
| GET   | /database/transactions | Возвращает список всех транзакций пользователя       | uid             | -                | [tid, category, date_time, description, money, type] |
| POST  | /database/transaction  | Добавляет новую транзакцию                           | uid             | money, date_time, type, description, category | - |
| PUT   | /database/transaction  | Изменение транзакции                                 | uid, tid        | money, date_time, type, description, category | - |
| Delete| /database/transaction  | Удаление транзакции                                  | uid, tid        | -                | - |

### Примеры запросов

#### POST
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
```bash
curl -X DELETE http://localhost:3000/database/transaction\?uid\=1\&tid\=1 
```
#### GET 
```bash
curl http://localhost:3000/database/transactions\?uid\=1
```

## Цели

| Метод | Эндпоинт        | Описание                                  | Параметры в URL | Принимаемый JSON | Возвращаемый JSON |
| ----- | --------------- | ----------------------------------------- | --------------- | ---------------- | --- |
| GET   | /database/goals | Возвращает список всех целей пользователя | uid             | -                | [goalid, money, goal_money, name, description, date_time_start, date_time_finish] |
| POST  | /database/goal  | Добавляет новую цуль                      | uid             | goal_money, name, description, date_time_start, date_time_finish | - |
| PUT   | /database/goal  | Изменение цели                            | uid, goalid     | goal_money, name, description, date_time_start, date_time_finish | - |
| Delete| /database/goal  | Удаление цули                             | uid, goalid     | -                | - |

### Примеры запросов

#### POST
```bash
curl -X POST http://localhost:3000/database/goal\?uid\=1 \
  -H "Content-Type: application/json" \
  -d '{
    "money": 10000000.0,
    "name": "Дом",
    "description": "Вилла в дубаях",
    "date_time": "2025-12-31T23:59:00",
    "date_time_finish": "2025-12-31T23:59:00"
}'
```
#### PUT
```bash
curl -X PUT http://localhost:3000/database/goal\?uid\=1\&goalid\=1 \
  -H "Content-Type: application/json" \
  -d '{
    "money": 99999999.0,
    "name": "Квартира",
    "description": "Сарайчик на окраине Москвы",
    "date_time": "2025-12-31T23:59:00",
    "date_time_finish": "2025-12-31T23:59:00"
}'
```
#### DELETE
```bash
curl -X DELETE http://localhost:3000/database/goal\?uid\=1\&goalid\=1 
```
#### GET
```bash
curl http://localhost:3000/database/goals\?uid\=1
```


