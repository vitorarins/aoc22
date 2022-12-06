package day3

import (
	"strings"
)

func GetSharedItem(rucksack string) string {
	halfLen := len(rucksack) / 2
	firstHalf := rucksack[:halfLen]
	secondHalf := rucksack[halfLen:]

	for _, c := range firstHalf {
		char := string(c)
		if strings.Contains(secondHalf, char) {
			return char
		}
	}

	return ""
}

func GetSharedItems(rucksacks []string) []string {

	result := []string{}

	for _, rs := range rucksacks {
		item := GetSharedItem(rs)
		result = append(result, item)
	}

	return result
}

func HasBadge(rucksacks []string, badge string) bool {
	rs := rucksacks[0]

	if !strings.Contains(rs, badge) {
		return false
	}

	if len(rucksacks) == 1 {
		return true
	}

	rucksacks = rucksacks[1:]

	return HasBadge(rucksacks, badge)
}

func GetSharedBadge(rucksacks []string) string {
	firstRs := rucksacks[0]

	for _, b := range firstRs {
		badge := string(b)
		if HasBadge(rucksacks, badge) {
			return badge
		}
	}

	return ""
}

func GetSharedBadges(rucksacks []string) []string {

	groupLen := 3
	rsCounter := 0
	limit := len(rucksacks) / groupLen

	result := []string{}

	for i := 0; i < limit; i++ {

		start := i * groupLen
		end := (rsCounter * groupLen) + groupLen

		group := rucksacks[start:end]

		badge := GetSharedBadge(group)
		result = append(result, badge)

		rsCounter++
	}

	return result
}

func GetSumPrio(sharedItems []string) int {

	lowOff := 96
	highOff := 38
	result := 0

	for _, c := range sharedItems {
		char := int(c[0])
		prio := char - highOff
		if char > lowOff {
			prio = char - lowOff
		}

		result += prio
	}

	return result
}

func GetSumPrioItems(input string) (int, error) {

	input = strings.TrimSpace(input)

	rucksacks := strings.Split(input, "\n")

	sharedItems := GetSharedItems(rucksacks)

	sumPrio := GetSumPrio(sharedItems)

	return sumPrio, nil
}

func GetSumPrioBadges(input string) (int, error) {

	input = strings.TrimSpace(input)

	rucksacks := strings.Split(input, "\n")

	sharedBadges := GetSharedBadges(rucksacks)

	sumPrio := GetSumPrio(sharedBadges)

	return sumPrio, nil
}
