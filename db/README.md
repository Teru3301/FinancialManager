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

| Метод | Эндпоинт               | Описание                                             | Параметры в URL | Параметры в JSON |
| ----- | ---------------------- | ---------------------------------------------------- | --------------- | ---------------- |
| GET   | /database/transactions | Возвращает список всех транзакций пользователя       | uid             | -                |
| POST  | /database/transaction  | Добавляет новую транзакцию                           | uid             | money, date_time, type, description, category |
| PUT   | /database/transaction  | Изменение транзакции                                 | uid, tid        | money, date_time, type, description, category |
| Delete| /database/transaction  | Удаление транзакции                                  | uid, tid        | -                |

## Группы

| Метод | Эндпоинт               | Описание                                             | Параметры в URL | Параметры в JSON |
| ----- | ---------------------- | ---------------------------------------------------- | --------------- | ---------------- |
| GET   | /database/groups       | Возвращает список всех групп в которых состоит пользователь | uid      | -                |
| POST  | /database/group        | Создаёт новую группу                                 | uid, gid        | name, description|
| PUT   | /database/group        | Изменяет данные группы                               | uid, gid        | name, description|
| Delete| /database/group        | Удаляет группу                                       | uid, gid        | -                |

## Цели

| Метод | Эндпоинт               | Описание                                             | Параметры в URL | Параметры в JSON |
| ----- | ---------------------- | ---------------------------------------------------- | --------------- | ---------------- |
| GET   | /database/goals        | Возвращает список целей пользователя                 | uid             | -                |
| POST  | /database/goals        | Создаёт новую пользовательскую цель                  | uid             | name, description, money, datestart, datefinish |
| PUT   | /database/goals        | Изменяет пользовательскую цель                       | uid, goalid     | name, description, money, datestart, datefinish |
| Delete| /database/goals        | Удаляет цель                                         | uid, goalid     | -                |



