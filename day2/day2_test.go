package day2_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day2"
)

// 2
// 2
// 2
// 5
// 2
func TestGetScoreSpec(t *testing.T) {
	testCases := map[string]struct {
		content string
		want    int
	}{
		"strategy": {
			content: `
A Y
B X
C Z
`,
			want: 15,
		},

		"strategy2": {
			content: `C Y
C Y
C Y
B Z
C Y
`,
			want: 17,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day2.GetScoreSpec(tc.content)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("content mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// 6
// 6
// 6
// 9
// 6
func TestGetScore(t *testing.T) {
	testCases := map[string]struct {
		content string
		want    int
	}{
		"strategy": {
			content: `
A Y
B X
C Z
`,
			want: 12,
		},

		"strategy2": {
			content: `C Y
C Y
C Y
B Z
C Y
`,
			want: 33,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day2.GetScore(tc.content)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("content mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
