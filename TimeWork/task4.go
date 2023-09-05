
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// Константа с датой
	const baseUnixTime int64 = 1589570165

	// Читаем строку с продолжительностью периода
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Используем регулярные выражения для извлечения минут и секунд
	re := regexp.MustCompile(`(\d+) мин\. (\d+) сек\.`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 3 {
		fmt.Println("Некорректный ввод")
		return
	}

	minutes, err1 := strconv.Atoi(matches[1])
	seconds, err2 := strconv.Atoi(matches[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка при преобразовании чисел")
		return
	}

	// Преобразуем продолжительность в тип Duration
	duration := time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second

	// Преобразуем Unix-Time в Time и добавляем к нему продолжительность
	resultTime := time.Unix(baseUnixTime, 0).In(time.UTC).Add(duration)

	// Выводим результат в формате UnixDate
	fmt.Println(resultTime.Format(time.UnixDate))
}
