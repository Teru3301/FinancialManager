package db

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

//	Конфигурационные параметры БД
//	В дальнейшем нужно будет вынести в отдельный файл и добавить его в .gitignore
var (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "fimmanuser"
	DBPassword = "superduperpassword123"
	DBName     = "finmandb"
	SSLMode    = "disable" // "disable" — для разработки, "require" — для продакшена
	connectionString = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		DBUser, DBPassword, DBHost, DBPort, DBName, SSLMode,
	)
	adminConnectionString = fmt.Sprintf(
		"postgres://postgres:your_postgres_password@%s:%s/postgres?sslmode=%s",
		DBHost, DBPort, SSLMode,
	)
	Pool *pgxpool.Pool
)


// Init — инициализирует соединение с БД, при необходимости создает базу и пользователя
func Init() error {
	log.Println("Initializing database...")

	err := tryConnect()
	if err == nil {
		CreateTables()
		return nil
	}

	log.Printf("Database connection failed, attempting to create: %v", err)
	if err := CreateDB(); err != nil {
		CreateTables()
		return fmt.Errorf("Failed to create database: %v", err)
	}

	return tryConnect()
}


// tryConnect — проверяет возможность подключения к основной базе данных
func tryConnect() error {
	ctx := context.Background()

	var err error
	Pool, err = pgxpool.New(ctx, connectionString)

	conn, err := Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	defer conn.Release()

	log.Println("Connection to DB success")

	return nil
}


// CloseDB закрывает пул соединений
func CloseDB() {
	if Pool != nil {
		Pool.Close()
		log.Println("DB connection pool closed")
	}
}



