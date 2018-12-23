package finance

import (
	"encoding/csv"
	"log"
)

// Table representing CSV file.
type Table struct {
	Headers []Header
	Rows    [][]Row
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
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(records) < 2 {
		log.Fatal("Need more than one header and row in CSV.")
	}
	for i, record := range records {
		if i == 0 {
			// The first row of the CSV file should always be the header.
			table.Headers = initHeaders(record)
			table.Rows = make([][]Row, len(records)-1)
			continue
		}
		table.Rows[i-1] = initRow(record, table.Headers)
	}
	return table
}

func initHeaders(record []string) []Header {
	headers := make([]Header, len(record))
	for i, value := range record {
		headers[i] = Header{value}
	}
	return headers
}

func initRow(record []string, headers []Header) []Row {
	rows := make([]Row, len(record))
	for i, value := range record {
		rows[i] = Row{Header: headers[i], Value: value}
	}
	return rows
}
