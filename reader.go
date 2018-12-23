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
	initializedHeader := false
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if initializedHeader {
			row := make([]Row, len(line))
			for i := 0; i < len(line); i++ {
				row[i] = Row{Header: table.Headers[i], Value: line[i]}
				print(row[i].Value)
			}
			print("\n\n")
		} else {
			table.Headers = make([]Header, len(line))
			for i := 0; i < len(line); i++ {
				table.Headers[i] = Header{line[i]}
			}
			initializedHeader = true
		}

	}
	return table
}
