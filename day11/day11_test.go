package day11_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day11"
)

func TestGetMonkeyBusiness(t *testing.T) {
	input := `
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`

	testCases := map[string]struct {
		rounds int
		worry  bool
		want   int
	}{
		"monkeys1Round": {
			rounds: 1,
			worry:  true,
			want:   24,
		},
		"monkeys20Rounds": {
			rounds: 20,
			worry:  true,
			want:   10197,
		},
		"monkeys1000Rounds": {
			rounds: 1000,
			worry:  true,
			want:   27019168,
		},
		"monkeys2000Rounds": {
			rounds: 2000,
			worry:  true,
			want:   108263829,
		},
		"monkeys10000Rounds": {
			rounds: 10000,
			worry:  true,
			want:   2713310158,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day11.GetMonkeyBusiness(input, tc.rounds, tc.worry)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("monkey business mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test20GetMonkeyBusiness(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"monkeys0": {
			input: `
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`,
			want: 10605,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day11.GetMonkeyBusiness20(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("monkey business mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test10000GetMonkeyBusiness(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"monkeys0": {
			input: `
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`,
			want: 2713310158,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day11.GetMonkeyBusiness10000(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("monkey business mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
