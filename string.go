package gosample

import (
	"errors"
	"strconv"
	"strings"
)

func testValidity(s string) (bool, []string) {
	splits := strings.Split(s, "-")
	lastPosition := 0 // ensure last item is a string
	for i, split := range splits {
		_, err := strconv.Atoi(split)
		if err != nil && i%2 == 0 { // The input is invalid if mod 0 split is not a number
			return false, nil
		} else if err == nil && i%2 == 1 { // The input is invalid if mod 1 split is a number
			return false, nil
		}
		lastPosition = i
	}
	if lastPosition%2 == 0 {
		return false, nil
	}
	return true, splits
}

func averageNumber(s string) (error, float64) {
	valid, splits := testValidity(s)
	if !valid {
		return errors.New("Invalid input string : " + s), 0
	}
	sum := 0
	for i, split := range splits {
		if i%2 == 0 { // only convert even indexed splits (0,2, 4, 6, etc)
			number, _ := strconv.Atoi(split)
			sum += number
		}
	}
	return nil, float64(sum) / float64(len(splits)/2) // split length / 2 as half of strings are numbers
}
