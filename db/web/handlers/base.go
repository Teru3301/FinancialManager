package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"fmt"
)

type BaseRequest struct {
	Token string `json:"token"`
	UID   string `json:"uid"`
	GID   string `json:"gid"`
	TID   string `json:"tid"`
	GoalID   string `json:"pid"`
	Other string `json:"other"`
	// другие общие поля
}

func Parse(c *fiber.Ctx) (*BaseRequest, error) {
    log.Printf("Incoming %s request: %s", c.Method(), c.Path())

    var req BaseRequest

    // Пробуем получить токен из заголовка Authorization (Bearer token)
    authHeader := c.Get("Authorization")
    if authHeader != "" {
        parts := strings.Split(authHeader, " ")
        if len(parts) == 2 && parts[0] == "Bearer" {
            req.Token = parts[1]
        }
    }

    req.UID = c.Query("uid")
    req.Other = c.Query("other")
    if req.Token == "" {
        req.Token = c.Query("token")
    }
	if err := c.BodyParser(&req); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "invalid request format")
    }

	//	Валидация токена
	if IsValidToken(req.Token) {
		fmt.Println("С токеном всё ок")
	} else {
		fmt.Println("Че с токеном?")
	}

    // Проверка прочих обязательных полей
    if req.UID == "" {
        fmt.Println("Че с юидом?")
    }

    return &req, nil
}

func IsValidToken(toker string) (bool) {
	//	токен отсутствует
	//	токен истёк
	//	токен не валидный
	return true
}
