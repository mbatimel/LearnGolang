package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)


func ReadDataInfile(path string) error{
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	buf := bufio.NewReader(file)

	defer file.Close()

	r := csv.NewReader(buf)
	
	for i := 0; i < 1; i++{
		row, err := r.Read()
		if err != nil && err != io.EOF {
			return err
		}
		rows := strings.Split(row[i], ";")
		for j := 0; j< len(rows); j++ {
			if rows[j] == "0"{
				fmt.Print(j)
			}
		}
	}


	return nil
}

func main() {
	const root = "./task.data"
	if err:= ReadDataInfile(root); err != nil {
		fmt.Print(err)
	}
	
}