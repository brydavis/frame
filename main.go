package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Frame map[string][]string

func main() {
	file, err := os.Open("crime.csv")
	// file, err := ioutil.ReadFile("crime.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	raw, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	df := Frame{}
	for k, head := range raw[0] {
		df[head] = []string{}
		// fmt.Println(k, head)
		for _, row := range raw[1:] {
			df[head] = append(df[head], row[k])
		}
	}

	fmt.Println(df.Row(4))
}

func (df Frame) Row(num int) (row []string) {
	for _, each := range df {
		row = append(row, each[num])
	}
	return
}
