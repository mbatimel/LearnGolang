package main

import (
 "fmt"
 "strings"
 "strconv"
 "os"
 "bufio"
 "io"
)

func main() {
 text, err := bufio.NewReader(os.Stdin).ReadString('\n')
 if err != nil && err != io.EOF {
  fmt.Println("Error!!!!!!!!!!!!", err)

}
text = strings.Trim(text, "\n")
 parts := strings.Split(text, ";")

 if len(parts) != 2 {
  fmt.Println("Invalid input format. Expected two numbers separated by ';'.")
  return
 }

 num1, err := parseNumber(parts[0])
 if err != nil {
  fmt.Println("Error parsing number 1:", err)
  return
 }

 num2, err := parseNumber(parts[1])
 if err != nil {
  fmt.Println("Error parsing number 2:", err)
  return
 }

 result := num1 / num2
 fmt.Printf("%.4f\n", result)
}

func parseNumber(input string) (float64, error) {
 input = strings.Replace(input, ",", ".", 1)
 input = strings.Replace(input, " ", "", -1)
 return strconv.ParseFloat(input, 64)
}