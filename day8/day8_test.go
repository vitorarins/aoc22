package day8_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day8"
)

func TestGetVisibleTrees(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"trees": {
			input: `
30373
25512
65332
33549
35390
`,
			want: 21,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day8.GetVisibleTrees(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("trees mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestParseInputTrees(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  [][]int
	}{
		"trees": {
			input: `
30373
25512
65332
33549
35390
`,
			want: [][]int{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 9, 0},
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day8.ParseInputTrees(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("trees mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetScenic(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"trees": {
			input: `
30373
25512
65332
33549
35390
`,
			want: 8,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day8.GetScenic(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("trees mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
