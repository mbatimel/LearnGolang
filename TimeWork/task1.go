
package main

import (
	"fmt"
	"time"
)

func main() {
	// Чтение строки даты и времени с стандартного ввода
	var inputDate string
	fmt.Scanln(&inputDate)

	// Парсинг строки в объект time.Time
	t, err := time.Parse(time.RFC3339, inputDate)
	if err != nil {
		fmt.Println("Ошибка при разборе даты:", err)
		return
	}

	// Форматирование объекта time.Time в строку UnixDate
	outputDate := t.Format(time.UnixDate)

	// Вывод результата
	fmt.Println(outputDate)
}