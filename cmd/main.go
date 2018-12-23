package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// Usage: go run main.go file1.csv file2.csv file3.csv ...
func main() {
	flag.Parse()
	files := flag.Args()
	fmt.Println("files:", files)

	csvFile, _ := os.Open(files[0])
	csvReader := csv.NewReader(bufio.NewReader(csvFile))

	accountTransactions := convertToProtos(csvReader)
	fmt.Println("Account:\n", accountTransactions.String())
}
