package main

import (
 "fmt"
 "strings"

)

func main() {
 var input, input2 string
 fmt.Scan(&input, &input2)
 fmt.Print(strings.Index(input, input2),)

}


