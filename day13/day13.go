package day13

import (
	"sort"
	"strconv"
	"strings"
)

func ParseChar(items []string, char rune, innerList int, listItem, normalItem []rune) ([]string, int, []rune, []rune) {
	switch {
	case char == '[':
		innerList++
		listItem = append(listItem, char)
	case char == ']':
		innerList--
		listItem = append(listItem, char)
	case char == ',':

		if innerList > 0 {
			listItem = append(listItem, char)
		} else {
			items, listItem, normalItem = AppendItem(items, listItem, normalItem)
		}

		return items, innerList, listItem, normalItem
	case innerList > 0:

		listItem = append(listItem, char)

		return items, innerList, listItem, normalItem
	default:
		normalItem = append(normalItem, char)
	}

	return items, innerList, listItem, normalItem
}

func AppendItem(items []string, listItem, normalItem []rune) ([]string, []rune, []rune) {
	if len(listItem) > 0 {
		items = append(items, string(listItem))
		listItem = []rune{}
	} else {
		items = append(items, string(normalItem))
		normalItem = []rune{}
	}

	return items, listItem, normalItem
}

func ParseList(list string) []string {
	list = strings.TrimSuffix(list, "]")
	list = strings.TrimPrefix(list, "[")

	if len(list) == 0 {
		return []string{""}
	}

	items := []string{}

	innerList := 0
	listItem := []rune{}
	normalItem := []rune{}

	for _, char := range list {
		items, innerList, listItem, normalItem = ParseChar(items, char, innerList, listItem, normalItem)
	}

	items, listItem, normalItem = AppendItem(items, listItem, normalItem)

	return items
}

func CompareListPackets(left, right string) int {
	leftList := ParseList(left)
	rightList := ParseList(right)

	diffLen := 0

	if len(leftList) != len(rightList) {
		diffLen = len(rightList) - len(leftList)
	}

	iter := len(rightList)
	if diffLen > 0 {
		iter = len(leftList)
	}

	for i := 0; i < iter; i++ {
		comparison := ComparePackets(leftList[i], rightList[i])
		if comparison != 0 {
			return comparison
		}
	}

	return diffLen
}

func ComparePackets(left, right string) int {

	switch {
	case len(left) == 0 && len(right) == 0:

		return 0

	case len(left) == 0 && len(right) != 0:

		return 1

	case len(left) != 0 && len(right) == 0:

		return -1

	case left[0] == '[' && right[0] == '[':

		return CompareListPackets(left, right)

	case left[0] == '[' && right[0] != '[':

		right = "[" + right + "]"

		return ComparePackets(left, right)

	case left[0] != '[' && right[0] == '[':

		left = "[" + left + "]"

		return ComparePackets(left, right)

	case left[0] != '[' && right[0] != '[':

		leftNr, _ := strconv.Atoi(left)
		rightNr, _ := strconv.Atoi(right)

		return rightNr - leftNr

	}

	return 0
}

func GetRightOrderSum(input string) int {

	input = strings.TrimSpace(input)

	pairsNotes := strings.Split(input, "\n\n")

	correctPairs := []int{}

	for i, pairStr := range pairsNotes {

		packets := strings.Split(pairStr, "\n")

		if ComparePackets(packets[0], packets[1]) >= 0 {
			correctPairs = append(correctPairs, i+1)
		}
	}

	sum := 0
	for _, index := range correctPairs {
		sum += index
	}

	return sum
}

func GetDecoderKey(input string) int {

	input = strings.TrimSpace(input)

	pairsNotes := strings.Split(input, "\n\n")

	packets := []string{"[[2]]", "[[6]]"}

	for _, pairStr := range pairsNotes {

		curPackets := strings.Split(pairStr, "\n")
		packets = append(packets, curPackets...)
	}

	sort.Slice(packets, func(i, j int) bool {
		return ComparePackets(packets[i], packets[j]) >= 0
	})

	mult := 0
	for i, packet := range packets {
		if packet == "[[2]]" {
			mult = i + 1
		}

		if packet == "[[6]]" {
			mult *= i + 1
		}
	}

	return mult
}
