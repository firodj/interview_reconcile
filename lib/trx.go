package lib

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

type InternalTrx struct {
	Amount float32
}

type ExternalTrx struct {
	Bank   string
	Amount float32
}

func InternalReader(in string, fromDate string, toDate string) error {
	trx := make([]InternalTrx, 0)

	err := CsvReader(in, func(record []string) {
		dates := strings.Split(record[3], "T")
		if dates[0] < fromDate || dates[0] > toDate {
			return
		}

		trxID := record[0]
		amount64, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			println("Error parsing amount for trxID", trxID)
		}
		amount := float32(amount64)

		if record[2] == "DEBIT" {
			amount = -amount
		}

		trx = append(trx, InternalTrx{
			Amount: amount,
		})
	})

	return err
}

func ExternalReader(in string, fromDate string, toDate string) error {
	trx := make([]ExternalTrx, 0)

	err := CsvReader(in, func(record []string) {
		date := record[2]
		if date < fromDate || date > toDate {
			return
		}

		unqID := record[0]
		banks := strings.Split(unqID, "-")

		amount64, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			println("Error parsing amount for unqID", unqID)
		}
		amount := float32(amount64)

		trx = append(trx, ExternalTrx{
			Bank:   banks[0],
			Amount: amount,
		})
	})

	return err
}

func CsvReader(in string, add func(record []string)) error {
	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		add(record)
	}

	return nil
}
