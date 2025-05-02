package db

import (
	"context"
	"fmt"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)


// CreateDB — создает базу данных и пользователя, если они ещё не существуют
func CreateDB() error {
	ctx := context.Background()

	adminPool, err := AdminConnect(ctx)
	if err != nil {
		return err
	}
	defer adminPool.Close()

	if err := CreateUserDB(ctx, adminPool); err != nil {
		return err
	}

	if err := CreateDatabase(ctx, adminPool); err != nil {
		return err
	}

	if err := grantPrivileges(ctx, adminPool); err != nil {
		return err
	}

	log.Println("DB and user successfully created")
	return nil
}


//	Подключение как админ для создания БД и пользователя
func AdminConnect(ctx context.Context) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(adminConnectionString)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфигурации администратора: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к административной БД: %w", err)
	}

	return pool, nil
}


//	Создание пользователя
func CreateUserDB(ctx context.Context, pool *pgxpool.Pool) error {
	query := fmt.Sprintf(`
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT FROM pg_catalog.pg_user WHERE usename = '%s'
			) THEN
				CREATE USER %s WITH PASSWORD '%s';
			END IF;
		END $$;`, DBUser, DBUser, DBPassword)

	_, err := pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("Error creating user: %w", err)
	}
	return nil
}


//	Создание БД
func CreateDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	query := fmt.Sprintf(`
		CREATE DATABASE %s
		OWNER %s
		ENCODING 'UTF8'
		LC_COLLATE 'C'
		LC_CTYPE 'C'
		TEMPLATE template0;`, DBName, DBUser)

	_, err := pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("Error creating DB: %w", err)
	}
	return nil
}


//	Выдача прав на БД
func grantPrivileges(ctx context.Context, pool *pgxpool.Pool) error {
	query := fmt.Sprintf(`
		GRANT ALL PRIVILEGES ON DATABASE %s TO %s;`, DBName, DBUser)

	_, err := pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("Error when issuing rights: %w", err)
	}
	return nil
}


