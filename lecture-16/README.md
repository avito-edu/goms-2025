# Conspectus: Testing

# Введение в тестирование в Go

Тестирование — важная часть разработки программного обеспечения. Оно помогает убедиться, что код работает корректно, уменьшает количество ошибок и облегчает поддержку проекта. В Go тестирование встроено в стандартную библиотеку, что делает его простым и удобным.

---

## Почему тестирование важно?

Тестирование позволяет:
- Находить ошибки на ранних этапах.
- Упрощать рефакторинг (если тесты проходят, значит, код работает).
- Документировать поведение функций (тесты показывают, как должен использоваться код).
- Ускорять разработку (автоматические тесты быстрее ручной проверки).

**Пример без тестирования:**
```go
// main.go
package main

func Add(a, b int) int {
    return a + b
}

func main() {
    result := Add(2, 3)
    println(result) // 5
}
```

Если мы случайно изменим Add (например, на return a - b), ошибка останется незамеченной.

## Unit-тесты (тестирование отдельных функций)
Unit-тесты проверяют работу отдельных функций или методов. В Go они пишутся в файлах с суффиксом _test.go.
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

```

Тесты:
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}
```

Запуск теста:
```bash
go test -v
```

##Integration-тесты (проверка взаимодействия компонентов)

```go
// db.go
package db

import "database/sql"

type User struct {
    ID   int
    Name string
}

func GetUser(db *sql.DB, id int) (*User, error) {
    var user User
    err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```

Тесты:
```go
// db.go
package db

import "database/sql"

type User struct {
    ID   int
    Name string
}

func GetUser(db *sql.DB, id int) (*User, error) {
    var user User
    err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```

## End-to-End тесты (проверка системы целиком)
E2E-тесты проверяют работу всей системы, например, HTTP-сервера.

```go
// server.go
package main

import (
    "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

func main() {
    http.HandleFunc("/", HelloHandler)
    http.ListenAndServe(":8080", nil)
}
```

Тесты:

```go
// server_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()

    HelloHandler(rr, req)

    if rr.Body.String() != "Hello, World!" {
        t.Errorf("Unexpected response: %s", rr.Body.String())
    }
}

func TestServer(t *testing.T) {
    // Запускаем сервер в тесте
    srv := httptest.NewServer(http.HandlerFunc(HelloHandler))
    defer srv.Close()

    // Делаем запрос к серверу
    resp, err := http.Get(srv.URL)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    // Проверяем ответ
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }
}
```

# Table-Driven тесты в Go

## Что такое table-driven тестирование?

Table-Driven тестирование — это подход, при котором тестовые случаи (входные данные и ожидаемые результаты) описываются в виде таблицы (обычно слайса структур), а затем один и тот же тест прогоняется для всех этих случаев в цикле.

**Пример простого теста без table-driven:**
```go
func TestAdd(t *testing.T) {
    if Add(1, 2) != 3 {
        t.Error("1 + 2 != 3")
    }
    if Add(0, 0) != 0 {
        t.Error("0 + 0 != 0")
    }
    if Add(-1, 1) != 0 {
        t.Error("-1 + 1 != 0")
    }
}
```
Тот же тест, но с table-driven подходом:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {100, -50, 50},
    }

    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf("Add(%d, %d) = %d, expected %d", 
                test.a, test.b, result, test.expected)
        }
    }
}
```

## Почему table-driven удобнее простого набора if?
### 1. Компактность и читаемость
Все тест-кейсы собраны в одном месте, их легко добавлять/изменять.

**Добавление нового случая:**
```go
tests := []struct {
    a, b     int
    expected int
}{
    {1, 2, 3},
    {0, 0, 0},
    {-1, 1, 0},
    // Новый случай:
    {10, -5, 5}, // Добавили одной строкой
}
```

### 2. Единообразие ошибок
Все ошибки выводятся в одинаковом формате, что упрощает анализ.

**Сравнение выводов:**

```go
// Простой тест:
add_test.go:10: 1 + 2 != 3
add_test.go:13: 0 + 0 != 0

// Table-Driven:
add_test.go:15: Add(1, 2) = 4, expected 3
add_test.go:15: Add(0, 0) = 1, expected 0
```

### 3. Возможность добавить описание тестов
Можно расширить структуру полем name для пояснения:

```go
func TestMultiply(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"Positive numbers", 2, 3, 6},
        {"With zero", 5, 0, 0},
        {"Negative", -2, 4, -8},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Multiply(test.a, test.b)
            if result != test.expected {
                t.Errorf("got %d, want %d", result, test.expected)
            }
        })
    }
}
```

### 4. Легко тестировать edge-кейсы
Можно систематически проверять граничные условия:

```go
func TestSafeDivide(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
        hasError bool
    }{
        {4, 2, 2, false},
        {1, 0, 0, true},  // Деление на ноль
        {-9, 3, -3, false},
    }

    for _, test := range tests {
        result, err := SafeDivide(test.a, test.b)
        if test.hasError {
            if err == nil {
                t.Errorf("Expected error for %d/%d", test.a, test.b)
            }
        } else {
            if result != test.expected {
                t.Errorf("%d/%d: got %d, want %d", 
                    test.a, test.b, result, test.expected)
            }
        }
    }
}
```

### 5. Поддержка субтестов (subtests)
Каждый случай можно запускать отдельно через t.Run():

```go
func TestParseDate(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected time.Time
        hasError bool
    }{
        {"Valid date", "2023-01-15", time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC), false},
        {"Invalid format", "15-01-2023", time.Time{}, true},
        {"Empty string", "", time.Time{}, true},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := ParseDate(test.input)
            if test.hasError {
                if err == nil {
                    t.Error("Expected error, got nil")
                }
            } else {
                if !result.Equal(test.expected) {
                    t.Errorf("got %v, want %v", result, test.expected)
                }
            }
        })
    }
}
```
**Преимущество:** При запуске go test -v вы увидите каждый подтест отдельно:

### 6. Легко пропускать определенные тесты
Можно добавлять флаги для сложных или долгих тестов:

```go
func TestComplexCalculation(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected float64
        skip     bool
    }{
        {"Simple case", 1, 1.0, false},
        {"Edge case", 1000, 123.45, true}, // Пропускаем в обычных прогонах
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            if test.skip && !testing.Short() {
                t.Skip("Skipping in short mode")
            }
            result := ComplexCalculation(test.input)
            if math.Abs(result-test.expected) > 0.01 {
                t.Errorf("got %.2f, want %.2f", result, test.expected)
            }
        })
    }
}
```

Запуск без пропусков:
```shell
go test -v
```

Запуск только быстрых тестов:
```shell
go test -v -short
```

# Использование моков и стабов в тестах Go

## Разница между моками и стабами

**Стабы (Stubs)** - простые заглушки, которые возвращают предопределённые значения:
```go
type UserRepositoryStub struct{}

func (u *UserRepositoryStub) GetUser(id int) (*User, error) {
    return &User{ID: 1, Name: "Test User"}, nil // Всегда возвращает одного и того же пользователя
}
```

**Моки (Mocks)** - более умные заглушки, которые дополнительно проверяют, как они были вызваны:

```go
type UserRepositoryMock struct {
    mock.Mock
}

func (u *UserRepositoryMock) GetUser(id int) (*User, error) {
    args := u.Called(id) // Фиксируем факт вызова
    return args.Get(0).(*User), args.Error(1) // Возвращаем то, что настроили в тесте
}

```
Ключевые различия:
- Стабы только возвращают данные
- Моки ещё и проверяют ожидания (был ли вызов, с какими параметрами, сколько раз)

## Как мокировать внешние зависимости

Пример без моков (проблема - реальное обращение к БД):
```go
func TestGetUserEmail(t *testing.T) {
    db := realDatabaseConnection() // Плохо для unit-теста!
    email, err := GetUserEmail(db, 1)
    if err != nil {
        t.Fatal(err)
    }
    if email != "user@example.com" {
        t.Errorf("unexpected email: %s", email)
    }
}
```

Решение с моком:
```go
func TestGetUserEmail(t *testing.T) {
    // 1. Создаём мок
    dbMock := new(DatabaseMock)
    
    // 2. Настраиваем ожидания
    dbMock.On("GetUser", 1).Return(&User{Email: "user@example.com"}, nil)
    
    // 3. Тестируем
    email, err := GetUserEmail(dbMock, 1)
    if err != nil {
        t.Fatal(err)
    }
    if email != "user@example.com" {
        t.Errorf("unexpected email: %s", email)
    }
    
    // 4. Проверяем, что мок был вызван как ожидалось
    dbMock.AssertExpectations(t)
}

```
## Использование testify/mock для упрощения моков
Библиотека testify/mock предоставляет удобный способ создания моков.

Установка:

```shell
go get github.com/stretchr/testify
Пример мока для интерфейса:
```
```go
// Интерфейс, который мы хотим замокать
type Mailer interface {
    Send(email string, body string) error
}

// MockMailer реализует интерфейс Mailer
type MockMailer struct {
    mock.Mock
}

func (m *MockMailer) Send(email string, body string) error {
    args := m.Called(email, body)
    return args.Error(0)
}

func TestNotificationService(t *testing.T) {
    // Создаём мок
    mailer := new(MockMailer)
    
    // Настраиваем ожидание
    mailer.On("Send", "user@example.com", "Hello").Return(nil)
    
    // Тестируем
    service := NewNotificationService(mailer)
    err := service.SendWelcomeEmail("user@example.com")
    assert.NoError(t, err)
    
    // Проверяем вызовы
    mailer.AssertExpectations(t)
}
```

## Использование mockery для генерации моков
mockery автоматически генерирует моки на основе интерфейсов.

Установка:
```shell
go install github.com/vektra/mockery/v2@latest
```
Генерация мока для интерфейса:
```shell
mockery --name=Mailer --output=mocks --case=underscore
```

Пример использования сгенерированного мока:

```go
func TestNotificationService_GeneratedMock(t *testing.T) {
    // Используем сгенерированный мок
    mailer := &mocks.Mailer{}
    
    mailer.On("Send", "user@example.com", "Welcome").Return(nil)
    
    service := NewNotificationService(mailer)
    err := service.SendWelcomeEmail("user@example.com")
    require.NoError(t, err)
    
    mailer.AssertExpectations(t)
}
```

## Практический пример: тестирование сервиса с моком БД

```go
// user_service.go
type UserService struct {
    repo UserRepository
}

func (s *UserService) GetUserName(id int) (string, error) {
    user, err := s.repo.GetUser(id)
    if err != nil {
        return "", err
    }
    return user.Name, nil
}

// user_service_test.go
func TestUserService_GetUserName(t *testing.T) {
    // 1. Создаём мок репозитория
    repoMock := new(UserRepositoryMock)
    
    // 2. Настраиваем ожидаемый вызов
    testUser := &User{ID: 1, Name: "Test User"}
    repoMock.On("GetUser", 1).Return(testUser, nil)
    
    // 3. Создаём сервис с моком
    service := &UserService{repo: repoMock}
    
    // 4. Вызываем метод
    name, err := service.GetUserName(1)
    
    // 5. Проверяем результат
    assert.NoError(t, err)
    assert.Equal(t, "Test User", name)
    
    // 6. Проверяем, что мок был вызван
    repoMock.AssertExpectations(t)
    
    // 7. Проверяем, что не было неожиданных вызовов
    repoMock.AssertNumberOfCalls(t, "GetUser", 1)
}

```

## Когда использовать моки, а когда стабы
**Используйте моки, когда нужно:**
- Проверить, что метод был вызван с определёнными параметрами
- Убедиться в определённой последовательности вызовов
- Проверить количество вызовов метода

**Используйте стабы, когда нужно:**
- Просто подменить реализацию для изоляции теста
- Вернуть заранее известные данные
- Заменить медленные или нестабильные зависимости


# Написание тестов для конкурентного кода в Go

Конкурентный код в Go (с использованием горутин и каналов) требует особого подхода к тестированию. В этой статье разберём:
- Как тестировать горутины
- Как выявлять гонки данных (data race)
- Использование `sync.WaitGroup` и `sync.Mutex` в тестах

---

## Особенности тестирования горутин

Горутины выполняются асинхронно, поэтому тесты должны дожидаться их завершения. Иначе тест может завершиться раньше, чем горутина выполнит свою работу.

### Пример 1: Тест без ожидания горутины (проблемный вариант)

```go
func ProcessAsync(data int, resultChan chan int) {
	go func() {
		resultChan <- data * 2
	}()
}

func TestProcessAsync_Bad(t *testing.T) {
	resultChan := make(chan int)
	ProcessAsync(5, resultChan)

	// Тест завершится, не дожидаясь результата!
	if result := <-resultChan; result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}
```
**Проблема**: В реальном коде горутина может не успеть выполниться.

### Пример 2: Правильный тест с ожиданием

```go
func TestProcessAsync_Good(t *testing.T) {
	resultChan := make(chan int, 1) // Буферизованный канал
	ProcessAsync(5, resultChan)

	select {
	case result := <-resultChan:
		if result != 10 {
			t.Errorf("Expected 10, got %d", result)
		}
	case <-time.After(1 * time.Second): // Таймаут на случай зависания
		t.Error("Timeout waiting for result")
	}
}
```

---

## Как выявлять гонки данных (data race)

Гонки данных возникают, когда несколько горутин одновременно обращаются к одной переменной, и хотя бы одна из них изменяет её.

### Пример 3: Код с гонкой данных

```go
var counter int

func Increment() {
	counter++
}

func TestIncrement_Race(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go Increment()
	}

	time.Sleep(1 * time.Second) // Наивная попытка дождаться завершения
	t.Logf("Counter: %d", counter) // Результат непредсказуем!
}
```
**Проблема**: Тест может выводить разные значения `counter` из-за гонки.

### Способ 1: Запуск теста с детектором гонок

Добавьте флаг `-race` при запуске тестов:
```bash
go test -race ./...
```
Детектор сообщит о проблеме:
```
WARNING: DATA RACE
Write at 0x00000123456 by goroutine 7:
  main.Increment()
```

### Способ 2: Использование `sync.Mutex`

```go
var (
	counter int
	mu      sync.Mutex
)

func IncrementSafe() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func TestIncrementSafe(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			IncrementSafe()
		}()
	}
	wg.Wait()

	if counter != 1000 {
		t.Errorf("Expected 1000, got %d", counter)
	}
}
```

---

## Использование `sync.WaitGroup` и `sync.Mutex` в тестах

### Пример 4: `WaitGroup` для ожидания горутин

```go
func ProcessBatch(data []int, resultChan chan int) {
	var wg sync.WaitGroup
	for _, num := range data {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultChan <- n * 2
		}(num)
	}
	wg.Wait()
	close(resultChan)
}

func TestProcessBatch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int, len(data))

	ProcessBatch(data, resultChan)

	var results []int
	for res := range resultChan {
		results = append(results, res)
	}

	expected := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, got %v", expected, results)
	}
}
```

### Пример 5: Тестирование с `Mutex`

Допустим, у нас есть кэш с конкурентным доступом:

```go
type Cache struct {
	mu    sync.Mutex
	items map[string]string
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.items[key]
	return val, ok
}
```

Тест:

```go
func TestCache_ConcurrentAccess(t *testing.T) {
	cache := &Cache{items: make(map[string]string)}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			cache.Set(key, "value")
		}(i)
	}

	wg.Wait()

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := cache.Get(key); !ok || val != "value" {
			t.Errorf("Key %s not found or incorrect value", key)
		}
	}
}
```

---

## Полезные советы

1. **Всегда используйте `-race`**
   ```bash
   go test -race ./...
   ```

2. **Избегайте `time.Sleep` в тестах**  
   Вместо этого используйте `WaitGroup` или каналы.

3. **Тестируйте разные сценарии**
    - Конкурентные запись и чтение
    - Очень большая нагрузка (1000+ горутин)
    - Ошибки и паники в горутинах

4. **Пример теста с паникой**
   ```go
   func TestGoroutinePanic(t *testing.T) {
       var wg sync.WaitGroup
       wg.Add(1)

       go func() {
           defer func() {
               if r := recover(); r == nil {
                   t.Error("Expected panic, got none")
               }
               wg.Done()
           }()
           panic("expected")
       }()

       wg.Wait()
   }
   ```

# Тестирование HTTP-серверов и клиентов в Go

В Go стандартная библиотека предоставляет мощные инструменты для тестирования HTTP-серверов и клиентов. В этой статье разберём, как тестировать обработчики API и как использовать пакет `httptest` для мокирования HTTP-запросов.

---

## Как тестировать обработчики API

Обработчики (handlers) в Go — это функции, которые принимают `http.ResponseWriter` и `http.Request`, обрабатывают запрос и возвращают ответ. Чтобы протестировать их, можно:
1. Создать тестовый HTTP-запрос.
2. Передать его в обработчик.
3. Проверить ответ.

### Пример 1: Простой обработчик

Допустим, у нас есть такой обработчик:

```go
package main

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
```

Тест для него:

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}

	expectedBody := "Hello, World!"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, rec.Body.String())
	}
}
```

### Пример 2: Обработчик с параметрами URL

Допустим, обработчик принимает параметр `name`:

```go
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name is required"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + name + "!"))
}
```

Тест:

```go
func TestGreetHandler(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"Valid name", "/greet?name=Alice", http.StatusOK, "Hello, Alice!"},
		{"Empty name", "/greet", http.StatusBadRequest, "Name is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.url, nil)
			rec := httptest.NewRecorder()

			GreetHandler(rec, req)

			if rec.Code != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, rec.Code)
			}

			if rec.Body.String() != tt.expectedBody {
				t.Errorf("Expected body '%s', got '%s'", tt.expectedBody, rec.Body.String())
			}
		})
	}
}
```

---

## Использование `httptest` для мокирования HTTP-запросов

Часто HTTP-клиенты взаимодействуют с внешними API, и в тестах нужно эмулировать их поведение. Пакет `net/http/httptest` позволяет создать фейковый сервер (`httptest.Server`), который можно использовать вместо реального API.

### Пример 3: Тестирование HTTP-клиента

Допустим, у нас есть клиент, который делает запрос к API:

```go
func GetUserInfo(apiURL string) (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
```

Тест с мок-сервером:

```go
func TestGetUserInfo(t *testing.T) {
	// Создаём тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "Alice", "age": 30}`))
	}))
	defer server.Close() // Важно закрыть сервер после теста

	// Вызываем функцию, передавая URL тестового сервера
	result, err := GetUserInfo(server.URL)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := `{"name": "Alice", "age": 30}`
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
```

### Пример 4: Тестирование ошибок сервера

Проверим, как клиент обрабатывает ошибки:

```go
func TestGetUserInfo_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err := GetUserInfo(server.URL)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	expectedErr := "HTTP error: 500"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
	}
}
```

### Пример 5: Мокирование задержки ответа

Иногда API отвечает медленно, и нужно проверить, как клиент обрабатывает таймауты:

```go
func TestGetUserInfo_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Имитируем долгий ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаём клиент с таймаутом
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	// Переопределяем клиент (в реальном коде лучше использовать интерфейсы)
	oldClient := http.DefaultClient
	defer func() { http.DefaultClient = oldClient }()
	http.DefaultClient = client

	_, err := GetUserInfo(server.URL)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}
```

# Тестирование работы с файловой системой в Go

Тестирование кода, взаимодействующего с файловой системой, представляет особые сложности. В этой статье мы рассмотрим:
- Проблемы тестирования файлового ввода-вывода
- Как использовать библиотеку `afero` для мокирования файловой системы

---

## Проблемы тестирования файлового ввода-вывода

### Основные проблемы:

1. **Зависимость от состояния файловой системы**
    - Тесты могут влиять друг на друга, изменяя одни и те же файлы
    - Тесты могут давать разные результаты на разных машинах

2. **Побочные эффекты**
    - Тесты могут создавать/удалять файлы, что нежелательно

3. **Проблемы производительности**
    - Работа с реальной файловой системой медленнее, чем с моками

### Пример 1: Проблемный тест

```go
func CountLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func TestCountLines(t *testing.T) {
	// Создаём временный файл
	err := os.WriteFile("test.txt", []byte("line1\nline2\nline3"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.txt") // Не всегда срабатывает при панике

	count, err := CountLines("test.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if count != 3 {
		t.Errorf("Expected 3 lines, got %d", count)
	}
}
```

**Проблемы**:
- Тест зависит от файловой системы
- Может не очистить файл при панике
- Не изолирован от других тестов

---

## Как мокировать файловую систему с `afero`

Библиотека `afero` предоставляет абстракцию файловой системы, позволяющую использовать как реальную ФС, так и моки.

### Установка:
```bash
go get github.com/spf13/afero
```

### Пример 2: Рефакторинг с использованием `afero`

```go
import (
	"github.com/spf13/afero"
)

func CountLines(fs afero.Fs, filename string) (int, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}
```

### Пример 3: Тест с моком файловой системы

```go
func TestCountLinesWithAfero(t *testing.T) {
	// Создаём мок файловой системы
	fs := afero.NewMemMapFs()

	// Создаём файл в памяти
	err := afero.WriteFile(fs, "test.txt", []byte("line1\nline2\nline3"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	count, err := CountLines(fs, "test.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if count != 3 {
		t.Errorf("Expected 3 lines, got %d", count)
	}
}
```

### Пример 4: Тестирование ошибок

```go
func TestCountLines_FileNotExists(t *testing.T) {
	fs := afero.NewMemMapFs()
	_, err := CountLines(fs, "nonexistent.txt")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}
```

---

## Продвинутые сценарии

### Пример 5: Тестирование записи в файл

```go
func WriteData(fs afero.Fs, filename string, data []byte) error {
	file, err := fs.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func TestWriteData(t *testing.T) {
	fs := afero.NewMemMapFs()
	testData := []byte("test data")

	err := WriteData(fs, "output.txt", testData)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Проверяем, что файл создан и содержит правильные данные
	exists, err := afero.Exists(fs, "output.txt")
	if err != nil || !exists {
		t.Error("File was not created")
	}

	content, err := afero.ReadFile(fs, "output.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if !bytes.Equal(content, testData) {
		t.Errorf("Expected %s, got %s", testData, content)
	}
}
```

### Пример 6: Тестирование работы с директориями

```go
func ListFiles(fs afero.Fs, dir string) ([]string, error) {
	return afero.ReadDirNames(fs, dir)
}

func TestListFiles(t *testing.T) {
	fs := afero.NewMemMapFs()

	// Создаём структуру директорий и файлов
	fs.MkdirAll("testdir/subdir", 0755)
	afero.WriteFile(fs, "testdir/file1.txt", []byte{}, 0644)
	afero.WriteFile(fs, "testdir/file2.txt", []byte{}, 0644)
	afero.WriteFile(fs, "testdir/subdir/file3.txt", []byte{}, 0644)

	files, err := ListFiles(fs, "testdir")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []string{"file1.txt", "file2.txt", "subdir"}
	if !reflect.DeepEqual(files, expected) {
		t.Errorf("Expected %v, got %v", expected, files)
	}
}
```

---

## Альтернативные подходы

### 1. Использование временных файлов

Если всё же нужно работать с реальной ФС:

```go
func TestWithTempFile(t *testing.T) {
	// Создаём временный файл
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Удаляем после теста

	content := []byte("test content")
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Тестируем на реальной ФС
	fs := afero.NewOsFs()
	count, err := CountLines(fs, tmpfile.Name())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 line, got %d", count)
	}
}
```

### 2. Использование интерфейсов

Можно создать собственный интерфейс файловой системы:

```go
type FileSystem interface {
	Open(name string) (File, error)
	Create(name string) (File, error)
	// другие методы
}

type File interface {
	io.Reader
	io.Writer
	io.Closer
}
```

# Тестирование работы с временем в Go

Работа с временем в тестах представляет особую сложность, так как функции вроде `time.Now()` и `time.Sleep()` делают тесты недетерминированными. Рассмотрим, как правильно тестировать временнозависимый код.

## Проблемы тестирования временнозависимого кода

### Почему `time.Now()` делает тесты нестабильными

Основные проблемы:
- Реальное время постоянно меняется
- Тесты могут давать разные результаты в разное время суток
- Проверки на конкретные временные значения хрупкие

**Плохой пример:**
```go
func IsMorning() bool {
    hour := time.Now().Hour()
    return hour >= 5 && hour < 12
}

func TestIsMorning(t *testing.T) {
    if !IsMorning() {
        t.Error("Expected morning time")
    }
}
```
Этот тест будет падать ночью и вечером!

## Как тестировать функции, зависящие от времени

### 1. Использование интерфейсов для инъекции зависимостей

Лучший подход - сделать время явной зависимостью:

```go
type Clock interface {
    Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time {
    return time.Now()
}

func IsMorning(clock Clock) bool {
    hour := clock.Now().Hour()
    return hour >= 5 && hour < 12
}
```

Теперь можно тестировать с mock-часами:

```go
type MockClock struct {
    fixedTime time.Time
}

func (m MockClock) Now() time.Time {
    return m.fixedTime
}

func TestIsMorning(t *testing.T) {
    tests := []struct {
        time time.Time
        want bool
    }{
        {time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC), true},   // утро
        {time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC), false},  // день
    }

    for _, tt := range tests {
        clock := MockClock{fixedTime: tt.time}
        got := IsMorning(clock)
        if got != tt.want {
            t.Errorf("IsMorning(%v) = %v, want %v", tt.time, got, tt.want)
        }
    }
}
```

### 2. Тестирование time.Sleep

Для тестирования функций с задержками:

```go
type Sleeper interface {
    Sleep(time.Duration)
}

type RealSleeper struct{}

func (RealSleeper) Sleep(d time.Duration) {
    time.Sleep(d)
}

func Countdown(sleeper Sleeper, n int) {
    for i := n; i > 0; i-- {
        sleeper.Sleep(1 * time.Second)
    }
}
```

Тест с mock-задержкой:

```go
type SpySleeper struct {
    calls int
    durations []time.Duration
}

func (s *SpySleeper) Sleep(d time.Duration) {
    s.calls++
    s.durations = append(s.durations, d)
}

func TestCountdown(t *testing.T) {
    sleeper := &SpySleeper{}
    Countdown(sleeper, 3)
    
    if sleeper.calls != 3 {
        t.Errorf("expected 3 sleeps, got %d", sleeper.calls)
    }
    
    expected := []time.Duration{1*time.Second, 1*time.Second, 1*time.Second}
    if !reflect.DeepEqual(sleeper.durations, expected) {
        t.Errorf("expected sleeps %v, got %v", expected, sleeper.durations)
    }
}
```

## Готовые решения для работы со временем

### Пакет clockwork

Библиотека [clockwork](https://github.com/jonboulle/clockwork) предоставляет готовую реализацию:

```go
import "github.com/jonboulle/clockwork"

func TestWithClockwork(t *testing.T) {
    // Создаем fake clock
    clock := clockwork.NewFakeClock()
    
    // Устанавливаем конкретное время
    clock.Set(time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC))
    
    // Проверяем утреннее время
    if !IsMorning(clock) {
        t.Error("Expected morning time")
    }
    
    // Можем "переводить" время вперед
    clock.Advance(2 * time.Hour)
    if IsMorning(clock) {
        t.Error("Expected not morning after advance")
    }
}
```

### Пакет testify/suite с временными моками

Для комплексных тестов:

```go
import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/mock"
)

type TimeMock struct {
    mock.Mock
}

func (m *TimeMock) Now() time.Time {
    args := m.Called()
    return args.Get(0).(time.Time)
}

type TimeTestSuite struct {
    suite.Suite
    timeMock *TimeMock
}

func (s *TimeTestSuite) SetupTest() {
    s.timeMock = new(TimeMock)
}

func (s *TimeTestSuite) TestIsMorning() {
    morningTime := time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC)
    s.timeMock.On("Now").Return(morningTime)
    
    s.True(IsMorning(s.timeMock))
    s.timeMock.AssertExpectations(s.T())
}

func TestTimeSuite(t *testing.T) {
    suite.Run(t, new(TimeTestSuite))
}
```

## Рекомендации по тестированию временнозависимого кода

1. **Избегайте прямых вызовов `time.Now()`** в бизнес-логике
2. **Используйте интерфейсы** для работы со временем
3. **Для простых случаев** достаточно мок-структур
4. **Для сложных сценариев** используйте готовые библиотеки вроде clockwork
5. **Тестируйте граничные случаи**:
    - Переход через полночь
    - Летнее время
    - Разные часовые пояса

Пример теста граничного случая:

```go
func TestMidnightTransition(t *testing.T) {
    clock := MockClock{
        fixedTime: time.Date(2023, 1, 1, 23, 59, 59, 0, time.UTC),
    }
    
    // Проверяем поведение перед полуночью
    if IsMorning(clock) {
        t.Error("23:59 should not be morning")
    }
    
    // Переводим время на 1 секунду вперед
    clock.fixedTime = clock.fixedTime.Add(1 * time.Second)
    
    // Теперь полночь - еще не утро
    if IsMorning(clock) {
        t.Error("00:00 should not be morning")
    }
    
    // Переводим на 5 часов вперед
    clock.fixedTime = clock.fixedTime.Add(5 * time.Hour)
    
    // 5 утра - уже утро
    if !IsMorning(clock) {
        t.Error("05:00 should be morning")
    }
}
```

Правильный подход к тестированию времени делает ваши тесты:
- **Надёжными** - не зависят от реального времени
- **Детерминированными** - всегда одинаковый результат
- **Поддерживаемыми** - легко изменять и расширять
