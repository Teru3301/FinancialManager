package db

import (
	"context"
	"fmt"
	"log"
)


// ColumnSchema описывает отдельную колонку таблицы
type ColumnSchema struct {
	Name string
	Type string // SQL тип: INTEGER, TEXT, etc.
}

// TableSchema описывает таблицу и её колонки
type TableSchema struct {
	Name    string
	Columns []ColumnSchema
}

//	транзакции
var transactionsTable = TableSchema{
	Name: "transactions",
	Columns: []ColumnSchema{
		{Name: "tid", Type: "SERIAL PRIMARY KEY"},
		{Name: "uid", Type: "INTEGER NOT NULL"},
		{Name: "money", Type: "NUMERIC(10,2) NOT NULL"},
		{Name: "time", Type: "TIMESTAMP NOT NULL"},
		{Name: "descroption", Type: "TEXT"},
		{Name: "type", Type: "VARCHAR(50) NOT NULL"},
	},
}

func CreateTables() error {
	ctx := context.Background()
	
	if err := EnsureTableExists(ctx, transactionsTable); err != nil {
		log.Fatalf("DB schema check failed: %v", err)
	}

	return nil
}

// EnsureTableExists проверяет существование таблицы и нужных колонок. Создает таблицу при необходимости.
func EnsureTableExists(ctx context.Context, schema TableSchema) error {
	conn, err := Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	// Проверка существования таблицы
	var exists bool
	checkQuery := `
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_schema = 'public' AND table_name = $1
		);
	`
	if err := conn.QueryRow(ctx, checkQuery, schema.Name).Scan(&exists); err != nil {
		return fmt.Errorf("error checking table existence: %w", err)
	}

	if !exists {
		log.Printf("Table '%s' does not exist. Creating...", schema.Name)

		// Формируем SQL для создания таблицы
		columnDefs := ""
		for i, col := range schema.Columns {
			def := fmt.Sprintf("%s %s", col.Name, col.Type)
			if i > 0 {
				columnDefs += ", "
			}
			columnDefs += def
		}

		createQuery := fmt.Sprintf(`CREATE TABLE %s (%s);`, schema.Name, columnDefs)
		_, err := conn.Exec(ctx, createQuery)
		if err != nil {
			return fmt.Errorf("error creating table '%s': %w", schema.Name, err)
		}

		log.Printf("Table '%s' created successfully.", schema.Name)
		return nil
	} else {
		log.Printf("Table '%s' already exist.", schema.Name)
	}

	// Проверка колонок
	columnSet := make(map[string]bool)
	columnsQuery := `
		SELECT column_name FROM information_schema.columns WHERE table_name = $1;
	`
	rows, err := conn.Query(ctx, columnsQuery, schema.Name)
	if err != nil {
		return fmt.Errorf("error querying columns: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return fmt.Errorf("error scanning column name: %w", err)
		}
		columnSet[name] = true
	}

	for _, col := range schema.Columns {
		if !columnSet[col.Name] {
			return fmt.Errorf("missing column '%s' in table '%s'", col.Name, schema.Name)
		}
	}

	log.Printf("Table '%s' with required columns exists.", schema.Name)
	return nil
}



