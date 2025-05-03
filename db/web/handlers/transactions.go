package handlers

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"log"
	"financialmanager/db/db"
	"strconv"
)


//	Запрос списка транзакций
/*
*	Возвращает JSON со списком транзакций пользователя по UID и подробной информацией по каждой
*/
func GetTransactions(c *fiber.Ctx) error {
	req, err := Parse(c)
	if err != nil {
		return c.Status(400).SendString("Invalid request: " + err.Error())
	}

	// Поиск транзакций
	rows, err := db.Select(db.SelectQuery{
		Table:   "transactions",
		Columns: []string{"tid", "money", "type", "category", "date_time", "description"},
		Condition: map[string]interface{}{
			"uid": req.UID,
		},
	})
	if err != nil {
		log.Fatalf("Ошибка при выполнении SELECT: %v", err)
	}

	var transactions []map[string]interface{}

	for _, row := range rows {
		tid, _ := strconv.Atoi(row[0])
		money, _ := strconv.ParseFloat(row[1], 64)	
		transaction := map[string]interface{}{
			"tid":         tid,
			"money":       money,
			"type":        row[2],   // type
			"category":    row[3],   // category
			"date_time":   row[4],   // date_time
			"description": row[5],   // description
		}
		transactions = append(transactions, transaction)
	}

	return c.JSON(transactions)
}


//	Добавление новой транзакции
/*
*	Принимает JSON с транзакцией сохраняет её в БД
*/
func AddTransaction(c *fiber.Ctx) error {
	req, err := Parse(c)
	if err != nil {
		return c.Status(400).SendString("Invalid request: " + err.Error())
	}

	_, err = db.Insert(db.InsertQuery{
		Table: "transactions",
		Data: map[string]interface{}{
			"uid":         req.UID,
			"money":       req.Money,
			"date_time":   time.Now(),
			"type":        req.Transtype,
			"category":    req.Category,
			"description": req.Description,
		},
	})

	if err != nil {
		return c.Status(500).SendString("DB insert error: " + err.Error())
	}

	return c.Status(201).SendString("Transaction added successfully")
}


//	Обновление транзакции
/*
*	Принимает JSON и обновляет каждое поле существующей транзакции значением из него
*/
func UpdateTransaction(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err
	}
	
	updateData := map[string]interface{}{
		"money":       req.Money,
		"category":    req.Category,
		"description": req.Description,
		"date_time":   req.Date_time,
		"type": req.Transtype,
	}

	// Создаем запрос на обновление
	updateQuery := db.UpdateQuery{
		Table:   "transactions",
		Data:    updateData,
		Condition: map[string]interface{}{
			"tid": req.TID,
			"uid": req.UID,
		},
	}

	_, err = db.Update(updateQuery)
	if err != nil {
		return c.Status(500).SendString("Error updating transaction: " + err.Error())
	}

	return c.Status(200).SendString("Updating success")
}


//	Удаление транзакции
/*
*	Удаляет транзакцию по TID
*/
func DeleteTransaction(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err
	}

	deleteQuery := db.DeleteQuery{
		Table: "transactions",
		Condition: map[string]interface{}{
			"tid": req.TID,
			"uid": req.UID,
		},
	}

	_, err = db.Delete(deleteQuery)
	if err != nil {
		return c.Status(500).SendString("Error deleting transaction: " + err.Error())
	}

	return c.Status(200).SendString("Transaction deleted successfully")
}


