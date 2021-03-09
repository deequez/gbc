package main

import (
	"compress/bzip2"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	rev, err := revenue("./data/taxi-1k.csv.bz2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rev)
}

// Design: consider passing io.Reader instead of a string.
func revenue(csvFile string) (float64, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	bz := bzip2.NewReader(file)
	rdr := csv.NewReader(bz)
	total, lnum := 0.0, 0 // state variables
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			break
		}
		lnum++
		if err != nil {
			return 0, fmt.Errorf("%s:%d: %s", csvFile, lnum, err)
		}
		// Skip header line
		if lnum == 1 {
			continue
		}
		amount, err := strconv.ParseFloat(record[len(record)-2], 64)
		if err != nil {
			return 0, fmt.Errorf("%s:%d: %s", csvFile, lnum, err)
		}
		total += amount
	}

	return total, nil
}
