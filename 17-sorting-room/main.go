package main

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	convertedNb := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", convertedNb)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if _, ok := fnb.(FancyNumber); ok {
		i, err := strconv.Atoi(fnb.Value())
		if err != nil {
			return 0
		}

		return i
	}

	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	convertedFn := float64(ExtractFancyNumber(fnb))
	return fmt.Sprintf("This is a fancy box containing the number %.1f", convertedFn)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	}

	return "Return to sender"
}

// ✅ Implementasi struct tambahan:

type simpleNumberBox struct {
	value int
}

func (s simpleNumberBox) Number() int {
	return s.value
}

type AnotherFancyNumber struct {
	n string
}

func (a AnotherFancyNumber) Value() string {
	return a.n
}

type numberBoxContaining struct {
	f float64
}

func (n numberBoxContaining) Number() int {
	return int(n.f)
}

func main() {
	fmt.Println(DescribeNumber(-12.345))
	// Output: This is the number -12.3

	fmt.Println(DescribeNumberBox(simpleNumberBox{12}))
	// Output: This is a box containing the number 12.0

	fmt.Println(ExtractFancyNumber(FancyNumber{"10"}))
	// Output: 10
	fmt.Println(ExtractFancyNumber(AnotherFancyNumber{"4"}))
	// Output: 0

	fmt.Println(DescribeFancyNumberBox(FancyNumber{"10"}))
	// Output: This is a fancy box containing the number 10.0
	fmt.Println(DescribeFancyNumberBox(AnotherFancyNumber{"4"}))
	// Output: This is a fancy box containing the number 0.0

	fmt.Println(DescribeAnything(numberBoxContaining{12.345}))
	// Output: This is a box containing the number 12.3
	fmt.Println(DescribeAnything("some string"))
	// Output: Return to sender
}
