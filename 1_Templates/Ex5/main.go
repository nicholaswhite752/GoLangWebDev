package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
	"text/template"
)

type Record struct {
	Date time.Time
	Open float64
	High float64
	Low float64
	Close float64
	Volume float64
	AdjClose float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	records := prs("table.csv")

	err := tpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}

}


func prs(filePath string) []Record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		high, _ := strconv.ParseFloat(row[2], 64)
		low, _ := strconv.ParseFloat(row[3], 64)
		close, _ := strconv.ParseFloat(row[4], 64)
		volume, _ := strconv.ParseFloat(row[5], 64)
		adjClose, _ := strconv.ParseFloat(row[6], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
			High: high,
			Low: low,
			Close: close,
			Volume: volume,
			AdjClose: adjClose,
		})
	}

	return records

}