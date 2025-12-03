//go:build integration

package integration

import (
	"database/sql"
	"os"
	"testing"

	"your-project/db" // Импортируем ваш пакет

	_ "github.com/lib/pq" // или другой драйвер БД
)

// testDB инициализирует подключение к тестовой БД
func testDB(t *testing.T) *sql.DB {
	connStr := os.Getenv("TEST_DB_CONN_STR")
	if connStr == "" {
		// fallback для локальной разработки
		connStr = "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	return db
}

// setupTestData подготавливает тестовые данные
func setupTestData(t *testing.T, db *sql.DB) {
	// Очищаем таблицу перед тестом
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("Failed to clean users table: %v", err)
	}

	// Вставляем тестовые данные
	_, err = db.Exec(`
		INSERT INTO users (id, name) VALUES 
		(1, 'Alice'),
		(2, 'Bob'),
		(3, 'Charlie')
	`)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}
}

func TestGetUser_Integration(t *testing.T) {
	db := testDB(t)
	defer db.Close()

	// Подготавливаем тестовые данные
	setupTestData(t, db)

	tests := []struct {
		name      string
		userID    int
		wantUser  *db.User
		wantError bool
		errorType error
	}{
		{
			name:   "existing user",
			userID: 1,
			wantUser: &db.User{
				ID:   1,
				Name: "Alice",
			},
			wantError: false,
		},
		{
			name:      "non-existing user",
			userID:    999,
			wantUser:  nil,
			wantError: true,
			errorType: sql.ErrNoRows,
		},
		{
			name:   "another existing user",
			userID: 2,
			wantUser: &db.User{
				ID:   2,
				Name: "Bob",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Вызываем тестируемую функцию
			gotUser, err := db.GetUser(db, tt.userID)

			// Проверяем ошибки
			if (err != nil) != tt.wantError {
				t.Errorf("GetUser() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if tt.wantError {
				// Проверяем тип ошибки для случая, когда ошибка ожидается
				if err != tt.errorType {
					t.Errorf("GetUser() error = %v, want %v", err, tt.errorType)
				}
				return
			}

			// Проверяем результат
			if gotUser == nil {
				t.Error("GetUser() returned nil user, but no error was expected")
				return
			}

			if gotUser.ID != tt.wantUser.ID {
				t.Errorf("GetUser() user ID = %v, want %v", gotUser.ID, tt.wantUser.ID)
			}

			if gotUser.Name != tt.wantUser.Name {
				t.Errorf("GetUser() user Name = %v, want %v", gotUser.Name, tt.wantUser.Name)
			}
		})
	}
}
