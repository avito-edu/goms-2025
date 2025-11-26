# Conspectus: Advanced-ErrorHandling

## 1. Обработка ошибок в Go

### Почему нет try-catch
- **Явная обработка** — ошибки видны в коде, нет "скрытых" исключений
- **Простота** — код легче читать и отлаживать
- **Контроль** — разработчик сам решает, как обрабатывать каждую ошибку

### error как стандартный механизм
```go
func ReadFile(filename string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return data, nil
}
```

### nil vs ошибка
- `nil` — операция завершилась успешно
- `error` — возникла проблема, требующая обработки

## 2. Кастомные ошибки

### Создание пользовательских ошибок
```go
// Через структуры
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Через errors.New
var ErrNotFound = errors.New("resource not found")
```

### Оборачивание ошибок
```go
import "errors"

if err != nil {
    return fmt.Errorf("reading config: %w", err)
}

// Проверка типа ошибки
if errors.Is(err, os.ErrNotExist) {
    // Обработка
}

// Извлечение конкретного типа
var myErr *MyError
if errors.As(err, &myErr) {
    // Используем myErr
}
```

## 3. Panic и Recover

### Что такое panic
- **Panic** — критическая СИТУАЦИЯ, останавливающая программу
- **Recover** — функция для перехвата panic (работает только в defer)

```go
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    
    // Код, который может вызвать panic
    panic("something went wrong")
}
```

### Когда использовать
- **Error** — ожидаемые ошибки (файл не найден, сетевые проблемы)
- **Panic** — критические ситуации, девочка/мальчик в дейтинге реджектнули, депрессия в 0 лет (инварианты нарушены, невозможные состояния)

## 4. Логирование

### В(л)ажность логирования
- Отладка и мониторинг приложения
- Аудит действий пользователей
- Анализ производительности

### Встроенный логгер
```go
import "log"

log.Println("Info message")
log.Printf("User %s logged in", username)
```

**Недостатки:**
- Нет уровней логирования
- Ограниченная кастомизация
- Нет структурированного вывода

## 5. Structured Logging

### Преимущества structured logging
- Машинно-читаемый формат (JSON)
- Легко фильтровать и анализировать
- Богатый контекст

### Популярные библиотеки

#### Logrus
```go
import log "github.com/sirupsen/logrus"

log.WithFields(log.Fields{
    "user": "john",
    "ip":   "192.168.1.1",
}).Info("User logged in")
```

#### Zap
```go
import "go.uber.org/zap"

logger, _ := zap.NewProduction()
defer logger.Sync()

logger.Info("User logged in",
    zap.String("user", "john"),
    zap.String("ip", "192.168.1.1"),
)
```

#### Zerolog
```go
import "github.com/rs/zerolog/log"

log.Info().
    Str("user", "john").
    Str("ip", "192.168.1.1").
    Msg("User logged in")
```

### Уровни логирования
- **DEBUG** — отладочная информация
- **INFO** — общая информация о работе
- **WARN** — предупреждения
- **ERROR** — ошибки, требующие внимания
- **FATAL** — критические ошибки (завершение программы)

## Ключевые выводы
1. Используйте error для ожидаемых ошибок, panic — для критических
2. Оборачивайте ошибки для сохранения контекста
3. Выбирайте structured logging для production-приложений
4. Используйте appropriate уровни логирования для разных ситуаций