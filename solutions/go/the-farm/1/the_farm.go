package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fooderCalculator FodderCalculator, cows int) (float64, error) {
	totalFood, err := fooderCalculator.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}
	multiplier, err := fooderCalculator.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return totalFood * multiplier / float64(cows), nil

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fooderCalculator FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fooderCalculator, cows)
	} else {
		return 0.0, errors.New("invalid number of cows")
	}
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (invalidCowsError *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", invalidCowsError.cows, invalidCowsError.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) err {
	if cows > 0 {
		return nil
	} else if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	} else {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
