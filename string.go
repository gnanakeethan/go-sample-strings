package gosample

import (
	"strconv"
	"strings"
)

func testValidity(s string) (bool, []string) {
	splits := strings.Split(s, "-")
	lastPosition := 0
	for i, split := range splits {
		_, err := strconv.Atoi(split)
		if err != nil && i%2 == 0 {
			return false, nil
		} else if err == nil && i%2 == 1 {
			return false, nil
		}
		lastPosition = i
	}
	if lastPosition%2 == 0 {
		return false, nil
	}
	return true, splits
}
