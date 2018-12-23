package finance

import (
	"encoding/csv"
	"io"
	"log"
)

// Table representing CSV file.
type Table struct {
	Headers []Header
	Rows    []Row
}

// Header in CSV table.
type Header struct {
	Name string
}

// Row in a CSV Table.
type Row struct {
	Header Header
	Value  string
}

// CreateTable creates a *Table given a *csv.Reader by parsing the lines it reads.
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
