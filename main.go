package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const (
	EXCELNAME  = "./assets/_prototype.xlsx"
	SHEET      = "Sheet1"
	OUTPUTNAME = "./assets/output.txt"
)

func main() {
	f, err := excelize.OpenFile(EXCELNAME)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows(SHEET)
	if err != nil {
		fmt.Println(err)
		return
	}

	var sumCol, val int
	res := make(map[string]int)
	var indexName = ""
	for index, row := range rows {
		if len(row) == 2 {
			val, _ = strconv.Atoi(row[1])
		} else if len(row) == 3 {
			val, _ = strconv.Atoi(row[2])
		} else {
			fmt.Println("some row don't have data")
			return
		}

		if sumCol == 0 {
			sumCol = val
			indexName = rows[index][0]
		}

		if row[0] == "" || row[0] == indexName {
			sumCol += val
			res[indexName] = sumCol
		} else {
			sumCol = val
			indexName = row[0]
			res[indexName] = val
		}
	}

	write, err := os.Create(OUTPUTNAME)
	if err != nil {
		log.Fatal(err)
	}

	for index, val := range res {
		num := strconv.Itoa(val)
		if _, err := write.WriteString(index + " " + num + "\n"); err != nil {
			log.Fatal(err)
		}
	}
	write.Close()
	fmt.Println("process run execute successfully...")
}
