package handlers

import "github.com/gofiber/fiber/v2"

//	GET
func TestGoals1(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("GET UID:"+req.UID+";")
}

//	POST
func TestGoals2(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("POST UID:"+req.UID+";")
}

//	PUT
func TestGoals3(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("PUT UID:"+req.UID+";")
}

