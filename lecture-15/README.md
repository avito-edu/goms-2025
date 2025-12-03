# Conspectus: DB stuff

## Итро

**Почему Go хорош для работы с БД:**
- **Производительность**: Статическая типизация и компиляция в нативный код
- **Простота**: Чистый синтаксис, стандартная библиотека `database/sql`
- **Статика**: Предотвращение ошибок на этапе компиляции
- **Конкурентность**: Горутины и каналы для эффективной работы с БД

**Популярные БД с Go:**
- **SQL**: PostgreSQL, MySQL, SQLite
- **NoSQL**: MongoDB, Redis
- **База**: PostgreSQL - наиболее популярен в Go-экосистеме

**Подходы к работе с БД:**
1. **Нативный SQL** (`database/sql`) - контроль, производительность
2. **ORM** (GORM, ent) - удобство, абстракция
3. **Query Builders** (jet, squirrel) - компромисс между SQL и ORM

---

## Работа с `database/sql`

### Подключение к БД
```go
import (
    "database/sql"
    _ "github.com/lib/pq" // Драйвер PostgreSQL
)

dsn := "postgres://user:pass@localhost/dbname?sslmode=disable"
db, err := sql.Open("postgres", dsn)
```

**DSN (Data Source Name) форматы:**
- PostgreSQL: `postgres://user:pass@host:port/dbname?param=value`
- MySQL: `user:pass@tcp(host:port)/dbname`
- SQLite: `file:test.db?cache=shared`

**Драйверы:**
- PostgreSQL: `pgx`, `pq`
- MySQL: `go-mysql-driver`
- SQLite: `go-sqlite3`

### Управление пулом подключений
```go
db.SetMaxOpenConns(25)      // Максимум открытых соединений
db.SetMaxIdleConns(5)       // Соединения в пуле бездействия
db.SetConnMaxLifetime(5*time.Minute) // Время жизни соединения
```

**Предотвращение утечек:**
- Всегда закрывать `Rows` и `Stmt`
- Использовать `defer` для гарантированного закрытия
- Мониторинг количества соединений

### Выполнение запросов
```go
// Для одной записи
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = $1", 1).Scan(&name)

// Для выборки
rows, err := db.Query("SELECT id, name FROM users")
defer rows.Close()
for rows.Next() {
    rows.Scan(&id, &name)
}

// Для INSERT/UPDATE/DELETE
result, err := db.Exec("INSERT INTO users(name) VALUES($1)", "John")
```

**Scan и NULL:**
```go
var name sql.NullString
err := rows.Scan(&name)
if name.Valid {
    // Использовать name.String
}
```

### Обработка ошибок
```go
err := db.QueryRow("SELECT ...").Scan(&value)
if err != nil {
    if err == sql.ErrNoRows {
        // Обработка отсутствия данных
    }
    // Другие ошибки
}
```

---

## Паттерны и практики

### Транзакции
```go
tx, err := db.BeginTx(ctx, &sql.TxOptions{
    Isolation: sql.LevelReadCommitted,
})
defer tx.Rollback()

// Выполнение операций в транзакции
_, err = tx.Exec("INSERT ...")
if err != nil {
    return err
}

err = tx.Commit()
```

**Уровни изоляции:**
- `LevelReadUncommitted`
- `LevelReadCommitted` (чаще всего)
- `LevelRepeatableRead`
- `LevelSerializable`

**Savepoints:**
```go
tx.Exec("SAVEPOINT sp1")
// Откат к savepoint
tx.Exec("ROLLBACK TO SAVEPOINT sp1")
```

### Подготовленные запросы (Prepared Statements)
```go
stmt, err := db.Prepare("SELECT name FROM users WHERE id = $1")
defer stmt.Close()

// Многократное использование
for _, id := range ids {
    stmt.QueryRow(id).Scan(&name)
}
```

**Когда использовать:**
- ✅ Многократное выполнение одинаковых запросов
- ❌ Разовые запросы (оверхед)
- ❌ Динамические запросы (меняется структура)

### Контекст
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Запрос с таймаутом
rows, err := db.QueryContext(ctx, "SELECT ...")
```

**Преимущества:**
- Отмена долгих запросов
- Таймауты для операций
- Распространение метаданных

---

## ORM: GORM

### Основы
```go
import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name  string
    Email string `gorm:"uniqueIndex"`
}

// Автомиграция
db.AutoMigrate(&User{})

// CRUD операции
db.Create(&User{Name: "John"})
db.First(&user, "name = ?", "John")
db.Model(&user).Updates(User{Name: "Jane"})
db.Delete(&user)
```

### Хуки
```go
func (u *User) BeforeCreate(tx *gorm.DB) error {
    u.CreatedAt = time.Now()
    return nil
}
```

### Raw SQL в GORM
```go
// Когда использовать:
// 1. Сложные JOIN
// 2. Оптимизированные запросы
// 3. Специфичные функции БД
db.Raw("SELECT * FROM users WHERE age > ?", 18).Scan(&users)
```

---

## Дополнительно

### sqlx
```go
import "github.com/jmoiron/sqlx"

type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
}

// Автоматическое сканирование в структуры
var users []User
err := sqlx.Select(db, &users, "SELECT * FROM users")

// Именованные параметры
_, err := db.NamedExec("INSERT INTO users (name) VALUES (:name)", 
    map[string]interface{}{"name": "John"})
```

### Миграции
**golang-migrate:**
```bash
# Создание миграции
migrate create -ext sql -dir migrations create_users_table

# Применение
migrate -path migrations -database postgres://... up

# Откат
migrate -path migrations -database postgres://... down
```

**Файл миграции:**
```sql
-- 001_create_users.up.sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

-- 001_create_users.down.sql
DROP TABLE users;
```

### Тестирование
**Docker для тестов:**
```go
import "github.com/ory/dockertest"

resource, _ := pool.Run("postgres", "latest", []string{
    "POSTGRES_PASSWORD=secret",
    "POSTGRES_DB=testdb",
})
defer pool.Purge(resource)
```

**SQLite в памяти:**
```go
db, err := sql.Open("sqlite3", ":memory:")
// Быстрые юнит-тесты без внешних зависимостей
```

**Моки с go-sqlmock:**
```go
db, mock, _ := sqlmock.New()
defer db.Close()

mock.ExpectQuery("SELECT (.+) FROM users").
    WithArgs(1).
    WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
        AddRow(1, "John"))
```

---

## Ключевые рекомендации

1. **Начинайте с `database/sql`** для понимания основ
2. **Используйте контекст** для управления временем жизни запросов
3. **Закрывайте ресурсы** (Rows, Stmt, Tx)
4. **Обрабатывайте ошибки** на каждом шаге
5. **Выбирайте инструменты по потребностям**:
    - Проект до 10 таблиц → `database/sql`/`sqlx`
    - Сложная бизнес-логика → ORM
    - Критичная производительность → нативный SQL
6. **Пишите тесты** с использованием моков или тестовых БД
7. **Используйте миграции** для управления схемой БД
8. **Мониторьте пул соединений** в продакшене

## Ресурсы для углубления
- [Документация database/sql](https://golang.org/pkg/database/sql/)
- [GORM документация](https://gorm.io/docs/)
- [SQLX библиотека](https://github.com/jmoiron/sqlx)
- [Awesome Go - Database](https://awesome-go.com/database/)