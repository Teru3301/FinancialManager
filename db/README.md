# Структура базы данных
[<img src="https://raw.githubusercontent.com/Teru3301/FinancialManager/dev/db/other/plan.png"/>]()

# Запуск модуля:
- перейдите в директорию модуля
```bash
cd FinancialManager/db
```
- запустите
```go
go run .
```

пока есть только 9 тестовых эндпоинта

### Для GET
http://localhost:3000/database/goals
http://localhost:3000/database/groups
http://localhost:3000/database/transactions

### Для POST адреса те же но для получения ответа нужно скопировать текст и вставить в строку поиска в ручную
```html
data:text/html,<body onload="document.forms[0].submit()"><form action='http://localhost:3000/database/goals?uid=1' method='post'><input type='hidden' name='json' value='{}'></form></body>
```
```html
data:text/html,<body onload="document.forms[0].submit()"><form action='http://localhost:3000/database/groups?uid=1' method='post'><input type='hidden' name='json' value='{}'></form></body>
```
```html
data:text/html,<body onload="document.forms[0].submit()"><form action='http://localhost:3000/database/transactions?uid=1' method='post'><input type='hidden' name='json' value='{}'></form></body>
```

### Как сделть PUT запрос я не разобрался

