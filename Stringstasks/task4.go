


package main

import (
 "fmt"

)

func main() {
 var input, newtext string

 fmt.Scan(&input)
 rs := []rune(input)
 
 for i := range rs {
  if i % 2 == 1 {
    newtext+=string(rs[i])
  }
}
fmt.Print(newtext)


}




