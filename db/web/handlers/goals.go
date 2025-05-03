package handlers

import (
	"github.com/gofiber/fiber/v2"
	"financialmanager/db/db"
	"strconv"
	"log"
)


//	GET
func GetGoals(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	//return c.SendString("GET UID:"+string(req.UID)+";")



	rows, err := db.Select(db.SelectQuery{
		Table:   "goals",
		Columns: []string{"goalid", "money", "goal_money","name", "description", "date_time_start", "date_time_finish"},
		Condition: map[string]interface{}{
			"uid": req.UID,
		},
	})
	if err != nil {
		log.Fatalf("Ошибка при выполнении SELECT: %v", err)
	}

	var goals []map[string]interface{}

	for _, row := range rows {
		goalid, _ := strconv.Atoi(row[0])
		money, _ := strconv.ParseFloat(row[1], 64)	
		goal_money, _ := strconv.ParseFloat(row[2], 64)	
		goal := map[string]interface{}{
			"goalid": goalid,
			"money": money,
			"goal_money": goal_money,
			"name": row[3],
			"description": row[4],
			"date_time": row[5],
			"date_time_finish": row[6],
		}
		goals = append(goals, goal)
	}

	return c.JSON(goals)
}


//	POST
func AddGoal(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}

	_, err = db.Insert(db.InsertQuery{
		Table: "goals",
		Data: map[string]interface{}{
			"uid":				req.UID,
			"money":			0.0,
			"goal_money":		req.Money,
			"name":				req.Name,
			"description":		req.Description,
			"date_time_start":	req.Date_time,
			"date_time_finish":	req.Date_time_finish,
		},
	})

	if err != nil {
		return c.Status(500).SendString("DB insert error: " + err.Error())
	}

	return c.Status(201).SendString("Goal added successfully")
}


//	PUT
func UpdateGoal(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}
	
	//	Логика тут
	//return c.SendString("PUT UID:"+string(req.UID)+";")

	_, err = db.Update(db.UpdateQuery{
		Table:   "goals",
		Data:    map[string]interface{}{
			"goal_money":	req.Money,
			"name":			req.Name,
			"description":	req.Description,
			"date_time_start":	req.Date_time,
			"date_time_finish":	req.Date_time_finish,
		},
		Condition: map[string]interface{}{
			"goalid": req.GoalID,
			"uid": req.UID,
		},
	})
	if err != nil {
		return c.Status(500).SendString("Error updating transaction: " + err.Error())
	}

	return c.Status(200).SendString("Updating success")
}


//	DELETE
func DeleteGoal(c *fiber.Ctx) error {
	req, err := Parse(c)
    if err != nil {
        return err	//	ошибка если какие-то данные не валидные в т.ч. токен
	}


	deleteQuery := db.DeleteQuery{
		Table: "goals",
		Condition: map[string]interface{}{
			"goalid": req.GoalID,
			"uid": req.UID,
		},
	}

	_, err = db.Delete(deleteQuery)
	if err != nil {
		return c.Status(500).SendString("Error deleting goal: " + err.Error())
	}

	return c.Status(200).SendString("Goal deleted successfully")
}
