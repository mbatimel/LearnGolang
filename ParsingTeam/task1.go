package main

import (
	"strconv"
	"unicode"
)

func adding(txt1, txt2 string) int64{
	txt11 := formatintString(txt1)
	txt22 := formatintString(txt2)
   
	 res1, err := strconv.ParseInt(txt11, 10, 64)
	 if err != nil { // не забываем проверить ошибку
		   panic(err)
	 }
	 res2, err := strconv.ParseInt(txt22,10,64)
	 if err != nil { // не забываем проверить ошибку
		   panic(err)
	 }
   
	 return res1+res2
   } 
   func formatintString(s string) string {
	 var newS string  
	   rs := []rune(s)
	 for v := range rs {
	   if unicode.IsDigit(rs[v]) {
		 newS += string(rs[v])
	   }
	 }
	 return newS
   }