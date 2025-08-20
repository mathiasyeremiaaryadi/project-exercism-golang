package main

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record
	for _, record := range in {
		if predicate(record) {
			out = append(out, record)
		}
	}

	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return record.Day >= p.From && record.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	out := Filter(in, ByDaysPeriod(p))

	var total float64
	for _, v := range out {
		total += v.Amount
	}

	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filteredCategory := Filter(in, ByCategory(c))
	if len(filteredCategory) == 0 {
		return 0, errors.New("unkown category " + c)
	}

	total := TotalByPeriod(filteredCategory, p)

	return total, nil
}

// Day1Records only returns true for records that are from day 1
func Day1Records(r Record) bool {
	return r.Day == 1
}

func main() {
	// p := DaysPeriod{From: 1, To: 31}

	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
	}

	fmt.Println(Filter(records, Day1Records))
	// =>
	// [
	//   {Day: 1, Amount: 15, Category: "groceries"}
	// ]

	records = []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	period := DaysPeriod{From: 1, To: 15}

	Filter(records, ByDaysPeriod(period))
	// =>
	// [
	//   {Day: 1, Amount: 15, Category: "groceries"},
	//   {Day: 11, Amount: 300, Category: "utility-bills"},
	//   {Day: 12, Amount: 28, Category: "groceries"},
	// ]

	records = []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	fmt.Println(Filter(records, ByCategory("groceries")))
	// =>
	// [
	//   {Day: 1, Amount: 15, Category: "groceries"},
	//   {Day: 12, Amount: 28, Category: "groceries"},
	// ]

	records = []Record{
		{Day: 15, Amount: 16, Category: "entertainment"},
		{Day: 32, Amount: 20, Category: "groceries"},
		{Day: 40, Amount: 30, Category: "entertainment"},
	}

	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}

	fmt.Println(TotalByPeriod(records, p1))
	// => 16

	fmt.Println(TotalByPeriod(records, p2))
	// => 50

	p1 = DaysPeriod{From: 1, To: 30}
	p2 = DaysPeriod{From: 31, To: 60}

	records = []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	fmt.Println(CategoryExpenses(records, p1, "entertainment"))
	// => 0, error(unknown category entertainment)

	fmt.Println(CategoryExpenses(records, p1, "rent"))
	// => 1300, nil

	fmt.Println(CategoryExpenses(records, p2, "rent"))
	// => 0, nil
}
