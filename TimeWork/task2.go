
package main

import (
	"fmt"
	"time"
)

func main() {
	// Чтение строки даты и времени с стандартного ввода
	var year, month, day, hour, minute, second int
	fmt.Scanf("%d-%d-%d %d:%d:%d", &year, &month, &day, &hour, &minute, &second)

	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)

	// Если время больше или равно 13:00, добавляем один день к дате
	if t.Hour() >= 13 {
		t = t.Add(24 * time.Hour)
	}

	// Вывод результата
	fmt.Printf("%04d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
