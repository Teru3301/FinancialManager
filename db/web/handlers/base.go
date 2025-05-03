package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
	"strconv"
	"fmt"
)

type BaseRequest struct {
	Token		string	`json:"token"`	//	токен
	UID			int 	`json:"uid"`		//	id пользователя
	GID			int 	`json:"gid"`		//	id группы
	TID         int		`json:"tid"`
	GoalID		int 	`json:"goalid"`	//	id цели
	Money       float64 `json:"money"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`	//	описание
	Category	string	`json:"category"`	//	категория (продукты/одежда/развлечения...)
	Date_time	string	`json:"date_time"`
	Date_time_finish	string	`json:"date_time_finish"`
	Transtype	string	`json:"type"`	//	зачисление/списание
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

	//	парсинг данных из URL
	uidStr := c.Query("uid")
	if uidStr != "" {
		if uidInt, err := strconv.Atoi(uidStr); err == nil {
			req.UID = uidInt
		} else {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid UID format")
		}
	}
	tidStr := c.Query("tid")
	if tidStr != "" {
		if tidInt, err := strconv.Atoi(tidStr); err == nil {
			req.TID = tidInt
		} else {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid UID format")
		}
	}
	goalidStr := c.Query("goalid")
	if goalidStr != "" {
		if goalidInt, err := strconv.Atoi(goalidStr); err == nil {
			req.GoalID = goalidInt
		} else {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid UID format")
		}
	}
	//	парсинг содержимого JSON
	if c.Method() != fiber.MethodGet && c.Method() != fiber.MethodDelete {
        if err := c.BodyParser(&req); err != nil {
            return nil, fiber.NewError(fiber.StatusBadRequest, "invalid request format")
        }
    }

	//	Валидация токена
	if IsValidToken(req.Token) {
		fmt.Println("С токеном всё ок")
	} else {
		fmt.Println("Че с токеном?")
	}

    // Проверка прочих обязательных полей
    if req.UID == 0 {
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
