package main

import (
 "fmt"
  "unicode"
)

func main() {
 var input string
 var res bool = true
 fmt.Scan(&input)
 rs := []rune(input)
  if len(rs) < 5 {
    fmt.Print("Wrong password")
}else{
 for i := range rs {
  
    if unicode.Is(unicode.Latin, rs[i]){
      continue
    }else if unicode.IsDigit(rs[i]){
      continue
    }else{
      res = false
      break
    }
  }


 if !res{
  fmt.Print("Wrong password")
 }else{
  fmt.Print("Ok")
 }
}


}


