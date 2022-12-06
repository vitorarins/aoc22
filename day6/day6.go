package day6

import (
	"fmt"
	"strings"
)

func IsValidMark(signal string) bool {
	track := make(map[rune]bool)

	for _, c := range signal {
		if _, ok := track[c]; ok {
			return false
		}

		track[c] = true
	}

	return true
}

func getStartOf(input string, qtyMarker int) (int, error) {
	input = strings.TrimSpace(input)

	limit := len(input) - qtyMarker

	for i := 0; i < limit; i++ {

		signal := input[i : i+qtyMarker]
		if IsValidMark(signal) {
			return i + qtyMarker, nil
		}
	}

	return len(input), fmt.Errorf("could not find start of marker with 'qtyMarker': %d", qtyMarker)
}

func GetFirstMark(input string) (int, error) {
	qtyMarker := 4

	return getStartOf(input, qtyMarker)
}

func GetStartOfMsg(input string) (int, error) {
	qtyMarker := 14

	return getStartOf(input, qtyMarker)
}
