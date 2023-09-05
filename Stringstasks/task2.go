package main

import (
 "fmt"

 "unicode/utf8"
)

func main() {
 var input string
 fmt.Scan(&input)

 if isPalindrome(input) {
  fmt.Println("Палиндром")
 } else {
  fmt.Println("Нет")
 }
}

func isPalindrome(s string) bool {
 length := utf8.RuneCountInString(s)
 rs:= []rune(s)

 for i := range rs {


    if rs[i] != rs[length-i-1] {
   return false
  }
 }
 return true

}

