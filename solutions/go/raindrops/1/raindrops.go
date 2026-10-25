package raindrops

import "fmt"

func Convert(number int) string {
	result := ""
	if number%3 == 0 {
		result = fmt.Sprintf("%sPling", result)
	}
	if number%5 == 0 {
		result = fmt.Sprintf("%sPlang", result)
	}
	if number%7 == 0 {
		result = fmt.Sprintf("%sPlong", result)
	}
	if result == "" {
		result = fmt.Sprint(number)
	}
	return result
}
