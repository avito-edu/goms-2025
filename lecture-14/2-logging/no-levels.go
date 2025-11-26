package main

import (
	"log"
	"os"
	"time"
)

var (
	orderID, userID, requestID, err = "", "", "", ""
	user, ip, product, qty          = "", "", "", ""
	hugeObject                      = `", "", "", "`
)

func main() {
	// NO levels DEBUG, INFO, WARN, ERROR
	log.Print("What level is this? Debug? Info? Error?") // Unclear!

	// Have to reinvent the wheel
	log.Print("[ERROR] Database error")
	log.Print("[INFO] Server started")
	log.Print("[DEBUG] Variable x = 5")

	// Проблемы такого подхода:

	// ❌ Нет фильтрации по уровням
	log.Print("[DEBUG] Very verbose debug information") // Всегда выводится!
	log.Print("[DEBUG] Another debug message")          // Даже в production!

	// ❌ Нет контроля в runtime
	// Нельзя динамически изменить уровень логирования
	// Всегда логируется ВСЁ

	// ❌ Сложный парсинг логов
	// Приходится анализировать текст вместо структурированных полей

	// ❌ Нет компиляционной оптимизации
	// Все строки форматируются, даже если уровень отключен

	// Bad - just text
	log.Printf("User error %s for order %s: %v", userID, orderID, err)

	// Good would be to have structured logs:
	// {
	//   "level": "error",
	//   "user_id": "123",
	//   "order_id": "456",
	//   "error": "database connection failed"
	// }

	// ❌ Hard to parse and analyze
	log.Printf("User %s from IP %s ordered product %s with quantity %d", user, ip, product, qty)
	// How to find all orders for specific product? Need complex regex!

	// ❌ Inconsistent format
	log.Printf("Order %s by user %s failed: %s", orderID, userID, err)          // Version 1
	log.Printf("Failed order: user=%s order=%s error=%v", userID, orderID, err) // Version 2
	log.Printf("ERROR: %s (user: %s, order: %s)", err, userID, orderID)         // Version 3

	// ❌ Difficult to add context
	// Want to add request_id, trace_id, session_id to all logs?
	log.Printf("[%s] User %s error: %v", requestID, userID, err)
	log.Printf("[%s] Processing order %s", requestID, orderID)
	// Very verbose and error-prone!

	// 3. Limited configuration

	// ❌ Can't change log level at runtime
	log.SetOutput(os.Stdout) // That's about it for configuration

	// ❌ No easy way to change format
	// Want JSON in production and text in development? Not possible!

	// ❌ Can't disable specific parts
	// Can't disable database logs but keep HTTP logs

	// ❌ No granular control
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Very limited flag options

	// 4. No context

	// ❌ No request-scoped logging
	/*
		func handleRequest(w
		http.ResponseWriter, r * http.Request) {
			// Can't automatically include request_id in all logs
			log.Printf("Processing request") // Where's the context?
			processOrder()
			log.Printf("Request completed") // Still no context!
		}

		func processOrder() {
			log.Printf("Processing order") // Completely disconnected from request
		}
	*/

	// ❌ No way to add common fields to all logs
	// Want to add service_name, version, environment to every log? Tough luck!

	// 5. Performance problems
	var expensiveOperation = func() string { return orderID }
	// ❌ No compile-time optimization
	log.Printf("Debug info: %s", expensiveOperation())
	// expensiveOperation() runs ALWAYS, even when debug is disabled

	// ❌ String formatting always happens
	for i := 0; i < 1000000; i++ {
		log.Printf("Processing item %d", i) // Formatting happens every time!
	}

	// ❌ No level checking before formatting
	debug := false
	if debug {
		log.Printf("Very verbose debug: %v", hugeObject) // Still formats the string!
	}

	// 6. Limited log rotation

	// ❌ Files grow indefinitely
	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE, 0666)
	log.SetOutput(file)
	// File will eventually fill the disk!

	// ❌ No automatic rotation by size/time
	// Have to implement manual rotation or use external tools

	// ❌ No compression of old logs
	// Old logs take up same space as new ones

	// Additional problems mentioned:

	// ❌ No runtime control
	log.Printf("This will always be logged") // No way to disable without code changes

	// ❌ Hard to parse
	log.Printf("Error: %s at %s by user %s", err, time.Now(), userID)
	// Good luck building queries on this!

	// ❌ No compile optimization
	// All log calls are always included in binary, even if unused
}
