package handlers

import "github.com/gofiber/fiber/v2"

//	GET
func TestGroups1(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("GET UID:"+string(req.UID)+";")
}

//	POST
func TestGroups2(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("POST UID:"+string(req.UID)+";")
}

//	PUT
func TestGroups3(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	return c.SendString("PUT UID:"+string(req.UID)+";")
}

