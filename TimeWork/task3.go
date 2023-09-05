
package main

import (
	"fmt"
	
	"time"
)

func main() {
	var year, month, day, hour, minute, second int
	var year2, month2, day2, hour2, minute2, second2 int
	fmt.Scanf("%d.%d.%d %d:%d:%d,%d.%d.%d %d:%d:%d", &day, &month, &year, &hour, &minute, &second, &day2,& month2, &year2, &hour2, &minute2, &second2)

	t1 := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	t2 := time.Date(year2, time.Month(month2), day2, hour2, minute2, second2, 0, time.UTC)
	
	var duration time.Duration
	if t1.Before(t2) {
		duration = t2.Sub(t1)
	} else {
		duration = t1.Sub(t2)
	}
	
	// Выводим результат
	fmt.Println(duration)

	
}
