package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	c := 0
	var err error = nil
	for n != 1 {
		c++
		if n < 1 {
			err = errors.New("Interger is smaller 1")
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}
	return c, err
}
