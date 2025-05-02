package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)



type InsertQuery struct {
	Table string
	Data  map[string]interface{}
}

type UpdateQuery struct {
	Table     string
	Data      map[string]interface{}
	Condition map[string]interface{}
}

type DeleteQuery struct {
	Table     string
	Condition map[string]interface{}
}

type SelectQuery struct {
	Table     string
	Columns   []string               // какие колонки выбирать
	Condition map[string]interface{} // условия выборки
}


//
func scanRowToMapFromRows(rows pgx.Rows) (map[string]interface{}, error) {
	defer rows.Close()

	if !rows.Next() {
		return nil, pgx.ErrNoRows
	}

	values, err := rows.Values()
	if err != nil {
		return nil, err
	}

	fields := rows.FieldDescriptions()
	result := make(map[string]interface{})
	for i, f := range fields {
		result[string(f.Name)] = values[i]
	}
	return result, nil
}


//
func Insert(q InsertQuery) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var cols, placeholders []string
	var values []interface{}
	i := 1

	for col, val := range q.Data {
		cols = append(cols, col)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		values = append(values, val)
		i++
	}

	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING *",
		q.Table,
		strings.Join(cols, ", "),
		strings.Join(placeholders, ", "),
	)

	log.Printf("[DB] Executing INSERT: %s | Values: %v", sql, values)

	rows, err := Pool.Query(ctx, sql, values...)
	if err != nil {
		log.Printf("[DB] INSERT ERROR: %v", err)
		return nil, err
	}

	result, err := scanRowToMapFromRows(rows)
	if err != nil {
		log.Printf("[DB] INSERT SCAN ERROR: %v", err)
	}
	return result, err
}


//
func Update(q UpdateQuery) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var setParts, whereParts []string
	var values []interface{}
	i := 1

	for col, val := range q.Data {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", col, i))
		values = append(values, val)
		i++
	}

	for col, val := range q.Condition {
		whereParts = append(whereParts, fmt.Sprintf("%s = $%d", col, i))
		values = append(values, val)
		i++
	}

	sql := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s RETURNING *",
		q.Table,
		strings.Join(setParts, ", "),
		strings.Join(whereParts, " AND "),
	)

	log.Printf("[DB] Executing UPDATE: %s | Values: %v", sql, values)

	rows, err := Pool.Query(ctx, sql, values...)
	if err != nil {
		log.Printf("[DB] UPDATE ERROR: %v", err)
		return nil, err
	}

	result, err := scanRowToMapFromRows(rows)
	if err != nil {
		log.Printf("[DB] UPDATE SCAN ERROR: %v", err)
	}
	return result, err
}


//
func Delete(q DeleteQuery) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var whereParts []string
	var values []interface{}
	i := 1

	for col, val := range q.Condition {
		whereParts = append(whereParts, fmt.Sprintf("%s = $%d", col, i))
		values = append(values, val)
		i++
	}

	sql := fmt.Sprintf(
		"DELETE FROM %s WHERE %s RETURNING *",
		q.Table,
		strings.Join(whereParts, " AND "),
	)

	log.Printf("[DB] Executing DELETE: %s | Values: %v", sql, values)

	rows, err := Pool.Query(ctx, sql, values...)
	if err != nil {
		log.Printf("[DB] DELETE ERROR: %v", err)
		return nil, err
	}

	result, err := scanRowToMapFromRows(rows)
	if err != nil {
		log.Printf("[DB] DELETE SCAN ERROR: %v", err)
	}
	return result, err
}


// Функция для выполнения SELECT-запроса и получения данных в виде списка строк
func Select(q SelectQuery) ([][]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Формируем SQL-запрос для SELECT
	var columns string
	if len(q.Columns) > 0 {
		columns = strings.Join(q.Columns, ", ")
	} else {
		columns = "*"
	}

	var whereParts []string
	var values []interface{}
	i := 1

	for col, val := range q.Condition {
		whereParts = append(whereParts, fmt.Sprintf("%s = $%d", col, i))
		values = append(values, val)
		i++
	}

	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s",
		columns,
		q.Table,
		strings.Join(whereParts, " AND "),
	)

	// Выполняем запрос
	rows, err := Pool.Query(ctx, sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Собираем результат
	var result [][]string
	for rows.Next() {
		// Чтение значений строки
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		// Преобразуем значения строки в список строк
		var row []string
		for _, value := range values {
			// Преобразуем каждое значение в строку
			switch v := value.(type) {
			case string:
				row = append(row, v)
			case []byte:
				row = append(row, string(v))
			case int, int32, int64:
				row = append(row, fmt.Sprintf("%d", v))
			case float32, float64:
				row = append(row, fmt.Sprintf("%f", v))
			case bool:
				row = append(row, fmt.Sprintf("%t", v))
			case time.Time:
				row = append(row, v.Format(time.RFC3339))
			default:
				row = append(row, fmt.Sprintf("%v", v))
			}
		}
		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}




