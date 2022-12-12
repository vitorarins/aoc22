package day12_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day12"
)

func TestGetPathSteps(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"path0": {
			input: `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`,
			want: 31,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day12.GetPathSteps(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("path steps mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetPathStepsManyStarts(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"path0": {
			input: `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`,
			want: 29,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day12.GetPathStepsManyStarts(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("path steps mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
