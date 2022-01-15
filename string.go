package gosample

import (
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
