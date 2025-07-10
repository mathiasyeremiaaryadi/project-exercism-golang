package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}

	bill[item] += units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billQuantity, isItemExist := bill[item]
	unitQuantity, isUnitExist := units[unit]

	if !isItemExist || !isUnitExist || unitQuantity > billQuantity {
		return false
	} else if billQuantity == unitQuantity {
		delete(bill, item)
	} else {
		bill[item] -= unitQuantity
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, ok := bill[item]; !ok {
		return 0, false
	}

	return bill[item], true
}

func main() {
	units := Units()
	fmt.Println(units)
	// Output: map[...] with entries like ("dozen": 12)

	bill := NewBill()
	fmt.Println(bill)
	// Output: map[]

	bill = NewBill()
	units = Units()
	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	// Output: true (since dozen is a valid unit)

	bill = NewBill()
	units = Units()
	ok = RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	// Output: false (because there are no carrots in the bill)

	bill = map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := GetItem(bill, "carrot")
	fmt.Println(qty)
	// Output: 12

	fmt.Println(ok)
	// Output: true
}
