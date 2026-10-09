package sorting

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
	return fmt.Sprintf("This is a box containing the number %d.0", nb.Number())
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
	switch t := fnb.(type) {
	case FancyNumber:
		number, _ := strconv.Atoi(t.Value())
		return number
	default:
		return 0
	}

}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fancyNumber := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %d.0", fancyNumber)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	var answer string
	switch v := i.(type) {
	case int:
		answer = DescribeNumber(float64(v))
	case float64:
		answer = DescribeNumber(float64(v))
	case NumberBox:
		answer = DescribeNumberBox(v)
	case FancyNumberBox:
		answer = DescribeFancyNumberBox(v)
	default:
		answer = "Return to sender"
	}
	return answer
}
