package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type StudentsStruct struct{
	ID int64
	Number string
	Year int64
	Students []map[string]interface{}
}

type StudentsAverage struct{
	Average float64
}

func main() {
    var stud StudentsStruct
	var aver StudentsAverage
	ScanStudent(&stud)
    aver.AverageRainting(&stud)
	PrintStudent(&aver)
}
func ScanStudent(st* StudentsStruct) {
	dec := json.NewDecoder(os.Stdin)
   
	err := dec.Decode(&st)
	if err == io.EOF {
		
		log.Fatal(err)
		
	}
	if err != nil {
		log.Fatal(err)
		
	}
}
func  PrintStudent(stud *StudentsAverage) {
	data, err := json.MarshalIndent(stud, "", "    ")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s",data)

}
func (avr *StudentsAverage) AverageRainting(st *StudentsStruct) {
	var totalRating float64
	var count int
   
	for _, stud := range st.Students {
	 ratings, ok := stud["Rating"].([]interface{})
	 if !ok {
	  log.Println("Failed to convert Rating to []interface{}")
	  continue
	 }
   
	 for _, rate := range ratings {
	  rateFloat, ok := rate.(float64)
	  if !ok {
	   log.Println("Failed to convert rate to float64")
	   continue
	  }
	  totalRating += rateFloat
	  count++
	 }
	}
	if count == 0 {
	 avr.Average = 0
	 return
	}
	avr.Average = totalRating / float64(len(st.Students))
   }
