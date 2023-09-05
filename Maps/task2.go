package main

import "fmt"

func main() {
	// Дано
	groupCity := map[int][]string{
		//10:  //[]string{...},
		//100: // []string{...},
		//1000:// []string{...},
	}
	cityPopulation := map[string]int{
		// пример данных
		"Город1": 110,
		"Город2": 55,
		"Город3": 1200,
		// ...
	}

	// Создаем временный объект map для городов из groupCity[100]
	tempCities := make(map[string]bool)
	for _, city := range groupCity[100] {
		tempCities[city] = true
	}

	// Пройти по всем городам из cityPopulation и проверить, содержатся ли они во временном объекте map
	for city := range cityPopulation {
		if _, ok := tempCities[city]; !ok {
			delete(cityPopulation, city)
		}
	}

	// Выводим для проверки
	fmt.Println(cityPopulation)
}
