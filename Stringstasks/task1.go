package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

 text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
 text = strings.TrimSpace(text) 

 if isValid(text) {
  fmt.Println("Right")
 } else {
  fmt.Println("Wrong")
 }
}

func isValid(text string) bool {
 rs := []rune(text)
 lenght := len(rs)
 if !unicode.IsUpper(rs[0]) {
  return false
 }

 if !strings.HasSuffix(string(rs[lenght-1]), ".") {
  return false
 }

 return true
}