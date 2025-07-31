package main

import "fmt"

func main() {
	// Создаем карту: ключ (string) - имя, значение (int) - возраст
	users := map[string]int{
		"Алексей": 28,
		"Мария":   24,
		"Иван":    31,
	}

	// 1. Вывод всех пользователей (перебор карты)
	fmt.Println("Все пользователи:")
	for name, age := range users {
		fmt.Printf("- %s: %d лет\n", name, age)
	}

	// 2. Получение возраста по ключу
	fmt.Println("\nВозраст Марии:", users["Мария"], "лет") // Прямой доступ

	// 3. Проверка наличия ключа
	if age, exists := users["Петр"]; exists {
		fmt.Println("Возраст Петра:", age)
	} else {
		fmt.Println("Петр не найден в списке") // Выведется это
	}

	// 4. Добавление нового пользователя
	users["Ольга"] = 29
	fmt.Println("\nПосле добавления Ольги:")
	fmt.Printf("%+v\n", users) // Вывод всей карты

	// 5. Удаление пользователя
	delete(users, "Иван")
	fmt.Println("\nПосле удаления Ивана:")
	for name, age := range users {
		fmt.Printf("- %s: %d лет\n", name, age)
	}
}
