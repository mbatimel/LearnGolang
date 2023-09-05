package main

import (
	"fmt"
	"strings"
)

// Battery представляет собой тип, удовлетворяющий интерфейсу fmt.Stringer.
type Battery struct {
	charge int
}

// NewBattery создает новый объект Battery на основе введенной строки.
func NewBattery(s string) *Battery {
	charge := strings.Count(s, "1")
	return &Battery{charge: charge}
}

// String реализует интерфейс fmt.Stringer для типа Battery.
func (b *Battery) String() string {
	discharge := 10 - b.charge
	return "[" + strings.Repeat(" ", discharge) + strings.Repeat("X", b.charge) + "]"
}

func main() {
	var input string
	fmt.Scanln(&input)

	// Создаем объект batteryForTest на основе введенной строки.
	
	// batteryForTest := NewBattery(input)

	// ... (здесь ваш код продолжается и вызывает функцию, принимающую batteryForTest)
}