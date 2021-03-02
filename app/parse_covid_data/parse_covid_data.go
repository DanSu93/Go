package parse_covid_data

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	maxRows = flag.Int("max-rows", 5, "Maximum number of lines to read")
	columns = flag.String("columns", "", "Comma-separated list of columns to print")
)

func parse() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("The number of command line arguments is invalid. " +
			"Make sure you have given the file name.")
	}

	fileName := strings.TrimSpace(flag.Arg(0))

	if _, err := os.Stat(fileName); err != nil {
		log.Fatalf("Can't check if the file exists: %v", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Can't open file:%v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Can't close file:%v", err)
		}
	}()

	if *maxRows < 1 || *maxRows > 100 {
		log.Fatalf("The max-rows value must be between 1 and 100")
	}

	if len(*columns) == 0 {
		log.Fatalf("No columns set to read from file")
	}
	columns := strings.Split(*columns, ",")

	colNumbers := make(map[string]int, len(columns))
	reader := csv.NewReader(file)
	fileCols, err := reader.Read()
	if err != nil {
		log.Fatalf("Cannot read the header line of the csv file:%v", err)
	}

	for i, col := range fileCols {
		for j, c := range columns {
			if col == c {
				colNumbers[c] = i
				columns = append(columns[:j], columns[j+1:]...)
			}
		}
	}

	if len(columns) > 0 {
		log.Fatalf("No columns found in file:%v", columns)
	}

	for i := 1; i <= *maxRows; i++ {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file:%v", err)
		}

		fmt.Printf("---------- String %d --------- \n", i)
		for key, col := range colNumbers {
			fmt.Printf("%s:\t%s\n", key, row[col])
		}
		fmt.Printf("-----------------------------\n\n")
	}
}
