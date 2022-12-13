package day13_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day13"
)

func TestParseList(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  []string
	}{
		"packet0": {
			input: "",
			want:  []string{""},
		},
		"packet1": {
			input: "[1,1,3,1,1]",
			want:  []string{"1", "1", "3", "1", "1"},
		},
		"packet2": {
			input: "[1,1,[3],1,1]",
			want:  []string{"1", "1", "[3]", "1", "1"},
		},
		"packet3": {
			input: "[1,[3,4],1]",
			want:  []string{"1", "[3,4]", "1"},
		},
		"packet4": {
			input: "[[[]]]",
			want:  []string{"[[]]"},
		},
		"packet5": {
			input: "[[]]",
			want:  []string{"[]"},
		},
		"packet6": {
			input: "[]",
			want:  []string{""},
		},
		"packet7": {
			input: "[[8,7,6]]",
			want:  []string{"[8,7,6]"},
		},
		"packet8": {
			input: "[[1],[2,3,4]]",
			want:  []string{"[1]", "[2,3,4]"},
		},
		"packet9": {
			input: "[1]",
			want:  []string{"1"},
		},
		"packet10": {
			input: "[[1],4]",
			want:  []string{"[1]", "4"},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day13.ParseList(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("parse mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestComparePackets(t *testing.T) {
	testCases := map[string]struct {
		left  string
		right string
		want  int
	}{
		"pairs0": {
			left:  "1",
			right: "1",
			want:  0,
		},
		"pairs1": {
			left:  "3",
			right: "5",
			want:  2,
		},
		"pairs2": {
			left:  "9",
			right: "8",
			want:  -1,
		},
		"pairs3": {
			left:  "[1,1,3,1,1]",
			right: "[1,1,5,1,1]",
			want:  2,
		},
		"pairs4": {
			left:  "[[1],[2,3,4]]",
			right: "[[1],4]",
			want:  2,
		},
		"pairs5": {
			left:  "[9]",
			right: "[[8,7,6]]",
			want:  -1,
		},
		"pairs6": {
			left:  "[[4,4],4,4]",
			right: "[[4,4],4,4,4]",
			want:  1,
		},
		"pairs7": {
			left:  "[7,7,7,7]",
			right: "[7,7,7]",
			want:  -1,
		},
		"pairs8": {
			left:  "[]",
			right: "[3]",
			want:  1,
		},
		"pairs9": {
			left:  "[[[]]]",
			right: "[[]]",
			want:  -1,
		},
		"pairs10": {
			left:  "[1]",
			right: "[5]",
			want:  4,
		},
		"pairs11": {
			left:  "[10]",
			right: "[5]",
			want:  -5,
		},
		"pairs12": {
			left:  "[[10],[4,3,5,[[],8,[4,3,10,9,4]],5],[[[6],[8]],9,1],[4,8,2,[[],[2,1,7]],6]]",
			right: "[[5],[10]]",
			want:  -5,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day13.ComparePackets(tc.left, tc.right)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("comparison mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetRightOrderSum(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"pairs0": {
			input: `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
			want: 13,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day13.GetRightOrderSum(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetDecoderKey(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"pairs0": {
			input: `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
			want: 140,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day13.GetDecoderKey(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("key mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
