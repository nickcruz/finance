package main

import (
	"encoding/csv"
	"log"

	"github.com/nickcruz/finance/message"
)

func convertToProtos(reader *csv.Reader) *message.AccountTransactions {
	accountTransactions := &message.AccountTransactions{}
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(records) < 2 {
		log.Fatal("Need more than one header and row in CSV.")
	}
	for i, record := range records {
		if i == 0 {
			// The first row of the CSV file should always be the field.
			accountTransactions.Field = initFields(record)
			accountTransactions.Transaction = make([]*message.Transaction, len(records)-1)
			continue
		}
		accountTransactions.Transaction[i-1] = &message.Transaction{Value: record}
	}
	return accountTransactions
}

func initFields(record []string) []*message.Field {
	fields := make([]*message.Field, len(record))
	for i, value := range record {
		fields[i] = &message.Field{Name: &value}
	}
	return fields
}
