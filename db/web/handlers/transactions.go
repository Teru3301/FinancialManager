package handlers

import "github.com/gofiber/fiber/v2"

//	Зaпрос списка транзакций
func GetTransactions(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("Metod:GET; UID:"+req.UID+"; Token:"+req.Token+";")

	//	id пользователя

	//	запрос к бд

	//	отправка списка транзакций [id, сумма, тип, описание, дата, время, категория]
}

//	Добавление новой транзакции
func AddTransactions(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("Metod:POST; UID:"+req.UID+"; Token:"+req.Token+";")

	//	сумма
	//	тип зачисление/списание
	//	дата
	//	время
	//	описание (не обязательно)
	//	id транзакции (мб генерируется автомаитчески и глобально)
	//	id пользователя
	//	категория

	//	запрос к бд

	//	отправка результата добавления (скорее всего всегда TRUE)
}

//	Обновление транзакции
func UpdateTransactions(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("Metod:PUT; UID:"+req.UID+"; Token:"+req.Token+";")

	//	id транзакции
	//	id пользователя
	//	новая сумма
	//	новое описание
	//	новая дата
	//	новое время
	//	новая категория

	//	проверка существования транзакции
	//	проверка владельца
	//	запрос к бд на обновление

	//	отправка результата обновления (TRUE, FALSE - если не пройдена проверка или данные не валидны)
}

//	Удаление транзакции
func DeleteTransactions(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("Metod:DELETE; UID:"+req.UID+"; Token:"+req.Token+";")

	//	id транзакции
	//	id пользователя

	//	проверка существования транзакции
	//	проверка владельца
	//	запрос к бд на удаление

	//	отправка результата удаления (TRUE, FALSE - если проверка не пройдена)
}

