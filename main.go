package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Frame map[string][]string

func ReadCSV(filename string) Frame {
	file, err := os.Open(filename)
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
		for _, row := range raw[1:] {
			df[head] = append(df[head], row[k])
		}
	}
	return df
}

////////////////////////////////////////////////////////////////////////////////
// Returns a row of data based on a given row number.       				  //
// For example, calling df.Row(x) would return a slice of data from that row. //
func (df Frame) Row(num int) (row []string) {
	for _, each := range df {
		row = append(row, each[num])
	}
	return
}

////////////////////////////////////////////////////////////////////////////////
// Returns a set row of data based on a range of row numbers.       		  //
// For example, calling df.Rows(x,y) would...								  //
func (df Frame) Rows(x, y int) (rows [][]string) {
	for ; x < y; x++ {
		var row []string
		for _, each := range df {
			row = append(row, each[x])
		}
		rows = append(rows, row)
	}
	return
}

////////////////////////////////////////////////////////////////////////////////
// Returns a set row of data based on a range of row numbers.       		  //
// For example, calling df.Rows(x,y) would...								  //
func (df Frame) Top(num int) {
	fmt.Println(df.Rows(0, num))
}

func (df Frame) Describe() Frame {
	// Needs updating for descriptive stats
	return df
}

func (df Frame) Headers() (heads []string) {
	for k, _ := range df {
		heads = append(heads, k)
	}
	return
}

func (df Frame) Rename(current, future string) Frame {
	df[future] = df[current]
	delete(df, current)
	return df
}

// func (df Frame) Headers() map[int]string {
// 	// headers := []string{}
// 	headers := map[int]string{}
// 	var i int
// 	for k, _ := range df {
// 		headers[i] = k
// 		i += 1
// 	}
// 	return headers
// }

func main() {
	df := ReadCSV("crime.csv")
	df.Rename("Year", "YEAR")
	// fmt.Println(df.Headers())
	fmt.Println(df["YEAR"])

}
