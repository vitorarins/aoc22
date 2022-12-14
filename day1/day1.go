package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Sum(calories []string) (int, error) {
	result := 0

	for _, calorie := range calories {
		cal, err := strconv.ParseInt(calorie, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("failed to sum %+v: %w", calories, err)
		}

		result += int(cal)
	}

	return result, nil
}

func InsertToRank(rank []int, calories int) []int {

	if len(rank) < 3 {
		rank[len(rank)] = calories

		return rank
	}

	sort.Ints(rank)

	if calories > rank[0] {
		rank[0] = calories
	}

	return rank
}

func SumRank(rank []int) int {

	result := 0

	for _, r := range rank {
		result += r
	}

	return result
}

func GetFattiesElf(content string) (int, error) {
	elves := strings.Split(content, "\n\n")

	fattiest := 0

	for _, elf := range elves {

		elf = strings.TrimSpace(elf)
		cals := strings.Split(elf, "\n")

		sumCals, err := Sum(cals)
		if err != nil {
			return 0, fmt.Errorf("failed to parse %+v: %w", elf, err)
		}

		if sumCals > fattiest {
			fattiest = sumCals
		}
	}

	return fattiest, nil
}

func GetThreeFattiest(content string) (int, error) {
	elves := strings.Split(content, "\n\n")

	rank := make([]int, 3)

	for _, elf := range elves {

		elf = strings.TrimSpace(elf)
		cals := strings.Split(elf, "\n")

		sumCals, err := Sum(cals)
		if err != nil {
			return 0, fmt.Errorf("failed to parse %+v: %w", elf, err)
		}

		rank = InsertToRank(rank, sumCals)
	}

	return SumRank(rank), nil
}
