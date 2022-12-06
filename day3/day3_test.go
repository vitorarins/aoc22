package day3_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day3"
)

func TestGetSharedItem(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  string
	}{
		"rucksack0": {
			input: "vJrwpWtwJgWrhcsFMMfFFhFp",
			want:  "p",
		},
		"rucksack1": {
			input: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			want:  "L",
		},
		"rucksack2": {
			input: "PmmdzqPrVvPwwTWBwg",
			want:  "P",
		},
		"rucksack3": {
			input: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",

			want: "v",
		},
		"rucksack4": {
			input: "ttgJtRGJQctTZtZT",

			want: "t",
		},
		"rucksack5": {
			input: "CrZsJsPPZsGzwwsLwLmpwMDw",
			want:  "s",
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day3.GetSharedItem(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum of priority items mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSharedItems(t *testing.T) {
	testCases := map[string]struct {
		input []string
		want  []string
	}{
		"rucksacks": {
			input: []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			want: []string{"p", "L", "P", "v", "t", "s"},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day3.GetSharedItems(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum of priority items mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSumPrio(t *testing.T) {
	testCases := map[string]struct {
		input []string
		want  int
	}{
		"sharedItems": {
			input: []string{"p", "L", "P", "v", "t", "s"},
			want:  157,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day3.GetSumPrio(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum of priority items mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSumPrioItems(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"rucksacks": {
			input: `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`,
			want: 157,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day3.GetSumPrioItems(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum of priority items mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHasBadge(t *testing.T) {
	testCases := map[string]struct {
		input []string
		badge string
		want  bool
	}{
		"rucksacks0": {
			input: []string{
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			badge: "v",
			want:  false,
		},
		"rucksacks1": {
			input: []string{
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			badge: "J",
			want:  false,
		},
		"rucksacks2": {
			input: []string{
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			badge: "r",
			want:  true,
		},
		"rucksacks3": {
			input: []string{
				"nddNNMMPNBnBNnBTQSShlSHghlDHBr",
				"VcccVmqJsJsjlTmzTDggmHHT",
				"VqLtFCqFJfVtVjsNgPNNMMWNwgtNvn",
			},
			badge: "g",
			want:  true,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day3.HasBadge(tc.input, tc.badge)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("shared badges mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSharedBadge(t *testing.T) {
	testCases := map[string]struct {
		input []string
		want  string
	}{
		"rucksacks0": {
			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			want: "r",
		},
		"rucksacks1": {
			input: []string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			want: "Z",
		},
		"rucksacks2": {
			input: []string{
				"nddNNMMPNBnBNnBTQSShlSHghlDHBr",
				"VcccVmqJsJsjlTmzTDggmHHT",
				"VqLtFCqFJfVtVjsNgPNNMMWNwgtNvn",
			},
			want: "g",
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day3.GetSharedBadge(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("shared badges mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSharedBadges(t *testing.T) {
	testCases := map[string]struct {
		input []string
		want  []string
	}{
		"rucksacks0": {
			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			want: []string{"r", "Z"},
		},
		"rucksacks1": {
			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
				"nddNNMMPNBnBNnBTQSShlSHghlDHBr",
				"VcccVmqJsJsjlTmzTDggmHHT",
				"VqLtFCqFJfVtVjsNgPNNMMWNwgtNvn",
			},
			want: []string{"r", "Z", "g"},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day3.GetSharedBadges(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("shared badges mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSumPrioBadges(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"rucksacks": {
			input: `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`,
			want: 70,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day3.GetSumPrioBadges(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum of priority items mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
