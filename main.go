package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Frame map[string][]string
type RawData [][]string

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
	return DataFrame(raw)
}

////////////////////////////////////////////////////////////////////////////////
// CREATEFRAME
// 														       				  //
// 																			  //
////////////////////////////////////////////////////////////////////////////////
func DataFrame(raw [][]string) Frame {
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
// ROW
// Returns a row of data based on a given row number.       				  //
// For example, calling df.Row(x) would return a slice of data from that row. //
////////////////////////////////////////////////////////////////////////////////
func (df Frame) Row(num int) (row []string) {
	for _, each := range df {
		row = append(row, each[num])
	}
	return
}

////////////////////////////////////////////////////////////////////////////////
// ROWS
// Returns a set row of data based on a range of row numbers.       		  //
// For example, calling df.Rows(x,y) would...								  //
////////////////////////////////////////////////////////////////////////////////
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
// TOP														        		  //
// Returns a set row of data based on a range of row numbers.
// For example, calling df.Rows(x,y) would...								  //
////////////////////////////////////////////////////////////////////////////////
func (df Frame) Top(num int) Frame {
	raw := df.Rows(0, num)
	fmt.Println(raw)
	return DataFrame(raw)
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

func (df1 Frame) Copy() Frame {
	// currently shares memory
	// goal: return frame with new memory
	df2 := df1
	return df2
}

func (df Frame) Raw() { // [][]string
	headers := df.Headers()
	fmt.Println(headers)
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

func (df Frame) Length() {
	fmt.Println("")
}

func main() {
	df1 := ReadCSV("crime.csv")
	// df2 := df1.Copy()
	// df2.Rename("Year", "Smoochies")
	// fmt.Println(df1.Headers(), df2.Headers())
	// fmt.Println(df["YEAR"])
	// fmt.Println(df1.Raw())

	// var raw [][]string
	// headers := []string{}
	// for k, v := range df1 {
	// 	fmt.Println(k, v)
	// 	headers = append(headers, k)
	// 	raw = append(raw, v)
	// }

	fmt.Println(len(df1["Year"]))

}
