package db

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Конфигурационные параметры БД
const (
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUser     = "kuzya"
	dbPassword = "gju67gjyg"
	dbName     = "grpprj"
	sslMode    = "disable" // "disable" для разработки, "require" для production
)

var (
	// Pool - пул соединений с основной БД
	Pool *pgxpool.Pool
	
	// connectionString - строка подключения к основной БД
	connectionString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)
	
	// adminConnectionString - строка подключения с правами админа (для создания БД)
	adminConnectionString = fmt.Sprintf("postgres://postgres:your_postgres_password@%s:%s/postgres?sslmode=%s",
		dbHost, dbPort, sslMode)
)

// Init инициализирует пул соединений и при необходимости создает БД
func Init() error {
	log.Println("Initializing database...")
	
	// Сначала пытаемся подключиться к существующей БД
	err := tryConnect()
	if err == nil {
		return nil // Подключение успешно, БД существует
	}
	
	log.Printf("Database connection failed, attempting to create: %v", err)
	
	// Если подключение не удалось, создаем БД и пользователя
	if err := CreateDB(); err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}
	
	// Пытаемся подключиться снова
	return tryConnect()
}

// tryConnect пытается подключиться к существующей БД
func tryConnect() error {
	var err error
	Pool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return fmt.Errorf("unable to create pool: %v", err)
	}
	
	conn, err := Pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("connection test failed: %v", err)
	}
	defer conn.Release()
	
	log.Println("Successfully connected to database")
	return nil
}

// CreateDB создает базу данных и пользователя
func CreateDB() error {
	// Подключаемся к стандартной БД postgres с правами админа
	config, err := pgxpool.ParseConfig(adminConnectionString)
	if err != nil {
		return fmt.Errorf("failed to parse admin config: %v", err)
	}
	
	// Создаем временное соединение
	tempPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("failed to connect to template database: %v", err)
	}
	defer tempPool.Close()
	
	// Создаем пользователя (если не существует)
	_, err = tempPool.Exec(context.Background(), 
		fmt.Sprintf("DO $$ "+
			"BEGIN "+
			"  IF NOT EXISTS (SELECT FROM pg_catalog.pg_user WHERE usename = '%s') THEN "+
			"    CREATE USER %s WITH PASSWORD '%s'; "+
			"  END IF; "+
			"END $$;", dbUser, dbUser, dbPassword))
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	
	// Создаем базу данных (если не существует)
	_, err = tempPool.Exec(context.Background(), 
		fmt.Sprintf("CREATE DATABASE %s OWNER %s ENCODING 'UTF8' LC_COLLATE 'C' LC_CTYPE 'C' TEMPLATE template0",
			dbName, dbUser))
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}
	
	// Даем все права пользователю на базу данных
	_, err = tempPool.Exec(context.Background(), 
		fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s", dbName, dbUser))
	if err != nil {
		return fmt.Errorf("failed to grant privileges: %v", err)
	}
	
	log.Println("Database and user created successfully")
	return nil
}

// CloseDB закрывает пул соединений
func CloseDB() {
	if Pool != nil {
		Pool.Close()
		log.Println("DB connection pool closed")
	}
}
