package main

import (
	"database/sql"
	"time"
)

const dsn = "postgres://user:pass@localhost:5432/app?sslmode=disable"

func main() {
	// БЕЗ пула - дорогое создание для каждого запроса
	for i := 0; i < 1000; i++ {
		db, _ := sql.Open("pgx", dsn) // Медленно!
		// выполнение запроса
		db.Close() // Закрытие соединения
	}

	// С пулом - переиспользуем существующие соединения
	db := connect() // Пул создается здесь
	for i := 0; i < 1000; i++ {
		// Используем соединение из пула
		db.Query("SELECT...") // Быстро!
	}

	db.SetMaxOpenConns(25)           // Макс. одновременных соединений
	db.SetMaxIdleConns(10)           // Макс. соединений в режиме ожидания
	db.SetConnMaxLifetime(time.Hour) // Время жизни соединения

	x := sql.ErrNoRows

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // rollback - на всякий случай

	// ...ИСПОЛНЯЕМ SQL...

	if err := tx.Commit(); err != nil {
		return err
	}
}
