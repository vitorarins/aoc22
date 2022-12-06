package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func GetSections(pair string) (map[int]struct{}, error) {

	result := make(map[int]struct{})

	n := strings.Split(pair, "-")

	start, err := strconv.Atoi(n[0])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(n[1])
	if err != nil {
		return nil, err
	}

	for i := start; i < end+1; i++ {
		result[i] = struct{}{}
	}

	return result, nil
}

type CheckBadFn func(map[int]struct{}, map[int]struct{}) bool

func copyMap(input map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for k, v := range input {
		result[k] = v
	}

	return result
}

func IsContained(firstSections, secondSections map[int]struct{}) bool {

	contains := copyMap(secondSections)

	for k := range firstSections {
		delete(contains, k)
	}

	if len(contains) == 0 {
		return true
	}

	contains = copyMap(firstSections)

	for k := range secondSections {
		delete(contains, k)
	}

	return len(contains) == 0
}

func IsOverlap(firstSections, secondSections map[int]struct{}) bool {

	for k := range firstSections {
		if _, ok := secondSections[k]; ok {
			return true
		}
	}

	return false
}

func IsBadAssign(assign string, checkBad CheckBadFn) (bool, error) {

	pairs := strings.Split(assign, ",")

	firstPair := pairs[0]
	secondPair := pairs[1]

	firstSections, err := GetSections(firstPair)
	if err != nil {
		return false, fmt.Errorf("failed to get sections of first pair %q: %w", firstPair, err)
	}

	secondSections, err := GetSections(secondPair)
	if err != nil {
		return false, fmt.Errorf("failed to get sections of second pair %q: %w", secondPair, err)
	}

	return checkBad(firstSections, secondSections), nil
}

func GetBadAssigns(assigns []string, checkBad CheckBadFn) ([]string, error) {

	badAssigns := []string{}

	for _, assign := range assigns {
		bad, err := IsBadAssign(assign, checkBad)
		if err != nil {
			return nil, fmt.Errorf("failed to check if bad assign %q: %w", assign, err)
		}

		if bad {
			badAssigns = append(badAssigns, assign)
		}
	}

	return badAssigns, nil
}

func GetBadAssignsAmnt(input string) (int, error) {

	input = strings.TrimSpace(input)

	assigns := strings.Split(input, "\n")

	badAssigns, err := GetBadAssigns(assigns, IsContained)
	if err != nil {
		return 0, fmt.Errorf("faild to get bad assigns: %w", err)
	}

	return len(badAssigns), nil
}

func GetOverlap(input string) (int, error) {

	input = strings.TrimSpace(input)

	assigns := strings.Split(input, "\n")

	badAssigns, err := GetBadAssigns(assigns, IsOverlap)
	if err != nil {
		return 0, fmt.Errorf("faild to get bad assigns: %w", err)
	}

	return len(badAssigns), nil
}
