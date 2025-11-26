// package main
//
// import (
//
//	"fmt"
//	"log"
//	"math/rand"
//	"os"
//	"runtime"
//	"time"
//
// )
//
// // Go - каждая операция с ошибкой видна явно
//
//	func processUserData(user User) error {
//		if err := validateUser(user); err != nil {
//			return fmt.Errorf("validate user: %w", err)
//		}
//		if err := saveToDatabase(user); err != nil {
//			return fmt.Errorf("save to db: %w", err)
//		}
//		if err := sendNotification(user); err != nil {
//			return fmt.Errorf("send notification: %w", err)
//		}
//		if err := updateCache(user); err != nil {
//			return fmt.Errorf("update cache: %w", err)
//		}
//		return nil
//	}
//
//	func main() {
//		// Cursed database logs
//		log.Println("🗄️  Database whispering ancient secrets... probably about your mom")
//		log.Printf("👻 Query returned %d rows of pure existential dread", rand.Intn(666))
//		log.Println("🔮 ORM tried to summon Cthulhu instead of mapping objects")
//
//		// Network madness
//		log.Println("📡 HTTP request got lost in the backrooms of the internet")
//		log.Printf("🌐 TCP packet %d was eaten by a firewall dragon", rand.Intn(1000))
//		log.Println("🚀 API response took so long it achieved consciousness")
//
//		// Authentication absurdity
//		log.Println("🔑 User authentication failed: password was 'password123' (seriously?)")
//		log.Printf("👤 User %d identified as a time-traveling potato", rand.Intn(100))
//
//		// System resource jokes
//
//		// Business logic gone wild
//		log.Println("💰 Payment processing: user paid in exposure and bad memes")
//		log.Printf("🛒 Shopping cart contains %d emotional baggage items", rand.Intn(10))
//		log.Println("📧 Email service currently judging your life choices")
//
//		// Pure chaos
//		log.Println("🎪 The code is compiling... or summoning demons. 50/50 chance.")
//		log.Printf("🐛 Bug report: feature working as intended? UNACCEPTABLE!")
//		log.Println("☕ Coffee level critical. Developer efficiency dropping to 0%")
//		log.Printf("🎲 Random error %d: the universe is probabilistic anyway", rand.Intn(999))
//
//		log.Println("💀 DEMON DEACTIVATED: That's enough chaos for today")
//	}
package main

import (
	"log"
	"os"
)

func main() {
	// Базовое логирование
	log.Print("Regular message")
	log.Println("Message with new line")
	log.Printf("Formatted message: %s", "value")

	// Логирование с префиксом
	log.SetPrefix("API: ")
	log.Print("Message with prefix")

	// Логирование в файл
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file")
	}
	log.SetOutput(file)

	// Fatal логи (вызывают os.Exit(1))
	// log.Fatal("Fatal error") // Программа завершится!
	// log.Panic("Panic error") // То же самое + panic
}
