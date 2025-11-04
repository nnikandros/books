package database

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeParsing(t *testing.T) {

	dateString := "2021-03-11"
	time1, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		time1, err = time.Parse("2006-01-02", dateString)
		if err != nil {
			t.Fatal("cannot parse using either layout:", err)
		}
	}

	fmt.Println(time1)

}

func TestBookModelDates(t *testing.T) {
	// year-month-day
	dateString := "2021-12-30"
	// time1, err := time.Parse(time.RFC3339, dateString)
	time1, err := time.Parse(time.DateOnly, dateString)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if err != nil {
		t.Fatalf("parsing %v", err)
	}
	fmt.Println(time1.Month())

	// b := BookModel{PublicationDate: time1}

	fmt.Println(time1.Format(time.DateOnly))

}
