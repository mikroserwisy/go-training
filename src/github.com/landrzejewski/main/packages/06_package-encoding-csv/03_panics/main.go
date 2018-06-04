package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type state struct {
	id               int
	name             string
	abbreviation     string
	censusRegionName string
}

func parseState(columns map[string]int, record []string) (*state, error) {
	panic("not implemented")
}

func main() {
	f, err := os.Open("../state_table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	columns := make(map[string]int)

	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		if rowCount == 0 {
			for idx, column := range record {
				columns[column] = idx
			}
		} else {
			state, err := parseState(columns, record)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(state)
		}
	}
}
