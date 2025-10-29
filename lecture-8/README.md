# Conspectus: Project Layout

## Что такое Go Modules и зачем они нужны
Go Modules — это система управления зависимостями в языке Go, появившаяся в версии 1.11. Она позволяет:
- Работать вне `GOPATH`.
- Фиксировать версии зависимостей.
- Упрощать импорт и управление пакетами.

## Проблема GOPATH в старых версиях Go
Ранее все проекты размещались в `GOPATH`, что вызывало проблемы:
- Все зависимости скачивались в одно место, что мешало их версионированию.
- Невозможно было легко работать с несколькими версиями одной зависимости.

## Основные команды для работы с модулями

```sh
# Инициализация модуля
$ go mod init mymodule

# Добавление зависимости
$ go get example.com/pkg@latest

# Удаление неиспользуемых зависимостей
$ go mod tidy

# Обновление всех зависимостей
$ go get -u ./...
```

## Структура файла `go.mod` и `go.sum`
Файл `go.mod` описывает модуль и его зависимости:
```go
module mymodule

go 1.25

require (
    example.com/pkg v1.2.3
)
```
Файл `go.sum` содержит контрольные суммы зависимостей.

## Как создавать и импортировать собственные пакеты

```sh
# Создадим структуру проекта
% mkdir -p mymodule/pkg/mypackage
% touch mymodule/pkg/mypackage/mypackage.go
```

Пример кода для `mypackage.go`:
```go
package mypackage

import "fmt"

func SayHello() {
    fmt.Println("Hello from mypackage!")
}
```

Использование пакета в `main.go`:
```go
package main

import "mymodule/pkg/mypackage"

func main() {
    mypackage.SayHello()
}
```

## Пакеты и их структура
Принято разделять код по логическим модулям:
```
project/
├── cmd/           # Исполняемые файлы
├── internal/      # Внутренние пакеты
├── pkg/           # Внешние пакеты
├── go.mod         # Файл модуля
├── go.sum         # Контрольные суммы
```

## Правила именования пакетов
- Имена должны быть короткими и осмысленными.
- Избегайте избыточных слов (`mypackageutils` → `utils`).
- Используйте единственное число (`models` → `model`).

## Видимость переменных, функций и типов (экспорт и импорт)
- **Экспортируемые** (public) объекты начинаются с заглавной буквы:
  ```go
  package mypackage
  
  func PublicFunc() {}
  ```
- **Неэкспортируемые** (private) объекты начинаются со строчной буквы:
  ```go
  package mypackage
  
  func privateFunc() {}
  ```

## Поиск пакетов на [pkg.go.dev](https://pkg.go.dev)
Можно искать сторонние библиотеки и просматривать их документацию.

## Пример использования внешнего пакета
```go
package main

import (
    "fmt"
	
    "github.com/google/uuid"
)

func main() {
    id := uuid.New()
    fmt.Println("Generated UUID:", id)
}
```

## Flat Structure: плюсы и минусы

### Что такое Flat Structure?
Flat Structure — это организация кода, при которой все файлы проекта находятся в одном уровне или минимально разнесены по директориям.

### Преимущества Flat Structure:
✅ Простота — легко ориентироваться в небольших проектах.  
✅ Быстрое добавление новых файлов без необходимости пересмотра структуры.  
✅ Удобство для скриптов и небольших утилит.

### Недостатки Flat Structure:
❌ Плохая масштабируемость — сложнее поддерживать большие проекты.  
❌ Возможность пересечения имен файлов.  
❌ Нарушение принципа разделения ответственности (код разных слоев смешан).

### Пример Flat Structure:
```
myapp/
├── main.go
├── config.go
├── database.go
├── handler.go
├── service.go
├── repository.go
```

Такой подход может быть удобен для маленьких программ, но быстро становится неуправляемым в масштабных проектах.

---

## Standard Project Layout

**Standard Project Layout** — это устоявшаяся структура проектов Go, рекомендованная сообществом и используемая в продакшн-проектах.

### Преимущества Standard Project Layout:
✅ Четкое разделение ответственности.  
✅ Масштабируемость — легко добавлять новые компоненты.  
✅ Удобство тестирования и поддержки кода.  
✅ Следование best practices, понятных другим разработчикам.

### Пример структуры:
```
myapp/
├── cmd/             # Точки входа в приложение
│   ├── app/         # Основное приложение
│   │   ├── main.go
│   ├── worker/      # Вторичный сервис (например, background worker)
│   │   ├── main.go
├── internal/        # Внутренние пакеты (нельзя импортировать извне)
│   ├── config/      # Конфигурация
│   ├── database/    # Логика работы с базой данных
│   ├── service/     # Бизнес-логика
│   ├── repository/  # Репозитории (работа с БД)
│   ├── handler/     # HTTP-хендлеры
├── pkg/             # Общие пакеты, которые можно переиспользовать
├── api/             # API-спецификации (Swagger, Protobuf)
├── configs/         # Конфигурационные файлы
├── scripts/         # Скрипты для деплоя, миграций и т. д.
├── go.mod           # Модульный файл Go
├── go.sum           # Контрольные суммы зависимостей
```

---

## Разделение бизнес-логики и слоев

### Основные слои в проекте

1. **Handler (Контроллеры)** — обработка входящих HTTP-запросов.
2. **Service (Сервисный слой)** — бизнес-логика.
3. **Repository (Репозиторий)** — взаимодействие с базой данных.

### Почему важно разделять слои?
- Повышает читаемость кода.
- Уменьшает связанность компонентов.
- Упрощает тестирование.

### Пример кода с разделением слоев

#### `handler/user.go` (Контроллер)
```go
package handler

import (
    "net/http"
    "myapp/internal/service"
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := h.service.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}
```

#### `service/user.go` (Сервисный слой)
```go
package service

import "myapp/internal/repository"

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) GetUser(id string) (string, error) {
    return s.repo.FindById(id)
}
```

#### `repository/user.go` (Репозиторий)
```go
package repository

type UserRepository struct {}

func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

func (r *UserRepository) FindById(id string) (string, error) {
    return "John Doe", nil
}
```

Этот подход делает код чистым, модульным и легко тестируемым.

---

### Итоги
- Flat Structure удобна для маленьких проектов, но не масштабируется.
- Standard Project Layout — лучший вариант для продакшн-проектов.
- Разделение слоев (Handler, Service, Repository) улучшает читаемость и тестируемость кода.

## Как правильно именовать пакеты в Go

### Основные рекомендации:
1. **Используйте осмысленные и короткие имена**
    - ❌ `mypackageutils` → ✅ `utils`
    - ❌ `mymodulehelpers` → ✅ `helpers`

2. **Не дублируйте имя пакета в названиях его содержимого**
   ```go
   // Плохо
   package config
   func ConfigLoad() {}

   // Хорошо
   package config
   func Load() {}
   ```

3. **Следуйте общепринятым названиям**
    - `config`, `logger`, `auth`, `handler`, `repository`

## Использование `internal` пакетов для ограничения видимости

Папка `internal/` запрещает импорт ее содержимого вне основного модуля.

```sh
project/
│── internal/
│   ├── db/
│   │   ├── connection.go
│── cmd/
│── main.go
```

Пример использования:
```go
package db

import "fmt"

func connect() {
    fmt.Println("Connecting to DB...")
}
```
Вне `internal/db`, импорт `project/internal/db` вызовет ошибку.

## Циклические зависимости: как их обнаружить и исправить

**Ошибка:**
- `package A` импортирует `package B`, а `package B` импортирует `package A`

**Как исправить:**
1. Вынести общие зависимости в третий пакет
2. Использовать интерфейсы

```go
package storage

type Storage interface {
    Save(data string)
}
```

## Использование интерфейсов для уменьшения связности

Интерфейсы позволяют зависеть от абстракций, а не конкретных реализаций:

```go
type Repository interface {
    GetUser(id int) string
}
```

Вместо конкретной структуры:
```go
type PostgresRepo struct{}
func (p PostgresRepo) GetUser(id int) string {
    return "User"
}
```

## Неправильное использование глобальных переменных в пакетах

Глобальные переменные могут привести к непредсказуемому поведению:
```go
package config
var DBConnection string
```

### Как исправить:
Использовать `sync.Once` для инициализации:
```go
package config
import "sync"
var (
    dbConnection string
    once sync.Once
)
func GetDBConnection() string {
    once.Do(func() {
        dbConnection = "Initialized"
    })
    return dbConnection
}
```

## Ошибки при работе с видимостью (exported/unexported)

В Go идентификаторы, начинающиеся с заглавной буквы, экспортируются:
```go
// Публичная функция
func PublicFunction() {}
// Приватная функция
func privateFunction() {}
```

### Ошибка:
- Попытка импортировать `privateFunction` в другом пакете

## Игнорирование семантического версионирования в модулях

Всегда фиксируйте версию зависимостей:
```sh
$ go get example.com/mypackage@v1.2.3
```

## Использование устаревших или небезопасных пакетов

Проверяйте зависимости:
```sh
$ go list -m -u all
```
Анализируйте уязвимости:
```sh
$ go list -m -json all | govulncheck
```

## Refs.

- [Effective Go — document gives tips for writing clear, idiomatic Go code](https://go.dev/doc/effective_go)
- [Golden config for golangci-lint v2.5.0 — very strict config](https://gist.github.com/maratori/47a4d00457a92aa426dbd48a18776322)
- [Organizing a Go module — based truth](https://go.dev/doc/modules/layout)
- [Standard Go Project Layout — lies from people of java world that’ve become standard](https://github.com/golang-standards/project-layout)
