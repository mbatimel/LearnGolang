package main

import (
 "fmt"

)

func main() {
 var input, newtext string
 fmt.Scan(&input)
 rs := []rune(input)
 for i := range rs {
  count := 0
  for j := range rs {
    if rs[i] == rs[j]{
      count++
    }
  }
  if count == 1 {
    newtext+=string(rs[i])
  }
}
fmt.Print(newtext)

}


