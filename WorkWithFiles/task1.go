package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func mywalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	if strings.Contains(info.Name(), "file") {
		if err := ReadDataInfile(path); err != nil {
			fmt.Printf("Error reading file")
		}
		return nil
	} else {
		return nil
	}
}
func ReadDataInfile(path string) error{
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(file)
	defer file.Close()
	r := csv.NewReader(buf)
	data, err := r.ReadAll()
	if err != nil {
		return err
	}
	if len(data) == 10{
		fmt.Print(data[4][2])
	}

	
	
	return nil
}

func main() {
	const root = "./task"
	if err := filepath.Walk(root, mywalkFunc); err != nil {
		fmt.Printf("Error : %v ", err)
	}
}
