package main

import "fmt"

// Определяем структуру Car
type Car struct {
	Brand string
	Model string
	Year  int
	Price float64
}

func main() {
	// Создаем экземпляр структуры Car
	myCar := Car{
		Brand: "Toyota",
		Model: "Camry",
		Year:  2022,
		Price: 25000.50,
	}

	// Выводим данные автомобиля
	fmt.Println("Данные автомобиля:")
	fmt.Printf("Марка: %s\n", myCar.Brand)
	fmt.Printf("Модель: %s\n", myCar.Model)
	fmt.Printf("Год выпуска: %d\n", myCar.Year)
	fmt.Printf("Цена: $%.2f\n", myCar.Price)

	// Альтернативный вывод (через %+v)
	fmt.Println("\nВывод через fmt.Printf(%+v):")
	fmt.Printf("%+v\n", myCar)
}
