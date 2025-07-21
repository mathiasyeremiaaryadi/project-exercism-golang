package main

import (
	"errors"
	"fmt"
)

// Error khusus untuk scale malfunction
var ErrScaleMalfunction = errors.New("scale malfunction")

// Interface WeightFodder
type WeightFodder interface {
	FodderAmount() (float64, error)
}

// Implementasi error SillyNephewError
type SillyNephewError struct {
	NumberCows int
}

func (s SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.NumberCows)
}

// Implementasi fungsi utama
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if fodder < 0 {
		if err == nil || err == ErrScaleMalfunction {
			return 0.0, errors.New("negative fodder")
		}
		return 0.0, err
	}

	if err != nil {
		if err == ErrScaleMalfunction && fodder > 0 {
			return (fodder * 2) / float64(cows), nil
		}
		return 0.0, err
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0.0, SillyNephewError{NumberCows: cows}
	}

	return fodder / float64(cows), nil
}

// Mock struct untuk implementasi WeightFodder
type Fodder struct {
	amount float64
	err    error
}

func (f Fodder) FodderAmount() (float64, error) {
	return f.amount, f.err
}

var (
	twentyFodderNoError                 = Fodder{amount: 20.0, err: nil}
	twentyFodderWithErrScaleMalfunction = Fodder{amount: 20.0, err: ErrScaleMalfunction}
	negativeFiveFodder                  = Fodder{amount: -5.0, err: nil}
)

func main() {
	fodder, err := DivideFood(twentyFodderNoError, 10)
	fmt.Println(fodder, err) // 2.0 <nil>

	fodder, err = DivideFood(twentyFodderWithErrScaleMalfunction, 10)
	fmt.Println(fodder, err) // 4.0 <nil>

	fodder, err = DivideFood(negativeFiveFodder, 10)
	fmt.Println(fodder, err) // 0.0 negative fodder

	fodder, err = DivideFood(twentyFodderNoError, 0)
	fmt.Println(fodder, err) // 0.0 division by zero

	fodder, err = DivideFood(twentyFodderNoError, -5)
	fmt.Println(fodder, err) // 0.0 silly nephew, there cannot be -5 cows
}
