package main

import (
	//"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	// "os"            // пакет используется для проверки ответа, не удаляйте его
)



func main() {
	value1, value2, operation := readTask() 
	if CheckValue(value1) == false  {
		PrintErr(value1)
	}else if CheckValue(value2) == false {
		PrintErr(value2)
	}else if CheckOpertion(operation) == false {
		fmt.Printf("неизвестная операция")
	}else{
	switch operation {
	case "+":
		fmt.Printf("%.4f",value1.(float64)+value2.(float64))
		case "-":
			fmt.Printf("%.4f",value1.(float64)-value2.(float64))

		case "*":
			fmt.Printf("%.4f",value1.(float64)*value2.(float64))

		case "/":
			fmt.Printf("%.4f",value1.(float64)/value2.(float64))

	}
}

}
 func readTask()(val1 interface{}, val2 interface{}, oper interface{}) {
	fmt.Scanf("%s%s%s",&val1, &val2, &oper)
	return val1, val2, oper
 }
 func CheckValue(val interface{}) bool{
	if _, ok := val.(float64); ok {
		return true
	}
	return false
 }
 func CheckOpertion(val interface{}) bool{
	if val ==  "+" || val ==  "-" || val == "*" || val == "/"{
		return true
	}
	return false
 }
 func PrintErr(val interface{}){
	switch i := val.(type){
	case int:
		fmt.Printf("value=%d: %T",val.(int),i)
	case string:
		fmt.Printf("value=%s: %T",val.(string),i)
	
	case bool:
		fmt.Printf("value=%t: %T",val.(bool),i)
	}

 }
 

