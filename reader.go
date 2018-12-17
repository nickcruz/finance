package finance

import (
	"encoding/csv"
	"io"
	"log"
)

type Table struct {
	Headers []Header
	Rows    []Row
}

type Header struct {
	Name string
}

type Row struct {
	Header Header
	Value  string
}

func CreateTable(reader *csv.Reader) *Table {
	table := &Table{}
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
	}
	return table
}
