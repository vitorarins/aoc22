package day1_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/vitorarins/aoc22/day1"
)

func TestSum(t *testing.T) {
	testCases := map[string]struct {
		calories []string
		want     int
	}{
		"ok": {
			calories: []string{"0"},
			want:     0,
		},

		"elves": {
			calories: []string{"7183", "3394"},
			want:     10577,
		},
		"elves2": {
			calories: []string{"25380", "25324", "16859"},
			want:     67563,
		},

		"elves3": {
			calories: []string{"7300", "3697"},
			want:     10997,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day1.Sum(tc.calories)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestInsertToRank(t *testing.T) {
	testCases := map[string]struct {
		rank    []int
		sumCals int
		want    []int
	}{
		"elves": {
			rank:    []int{40789, 23098, 19873},
			sumCals: 22342,
			want:    []int{22342, 23098, 40789},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day1.InsertToRank(tc.rank, tc.sumCals)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("rank mismatch (-want +got):\n%s", diff)
			}
		})
	}

}

func TestSumRank(t *testing.T) {
	testCases := map[string]struct {
		rank []int
		want int
	}{
		"elves": {
			rank: []int{40789, 23098, 19873},
			want: 83760,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day1.SumRank(tc.rank)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum rank mismatch (-want +got):\n%s", diff)
			}
		})
	}

}

func TestGetFattiesElf(t *testing.T) {
	testCases := map[string]struct {
		content string
		want    int
	}{
		"elves": {
			content: `7183
3394

25380
25324
16859

7300
3697
`,
			want: 67563,
		},

		"elves2": {
			content: `
8023
3736

40312

8527
4368
1291
1159
2869
10302
8489

3268
4846`,
			want: 40312,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day1.GetFattiesElf(tc.content)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetThreeFatties(t *testing.T) {
	testCases := map[string]struct {
		content string
		want    int
	}{
		"elves": {
			content: `7183
3394

25380
25324
16859

7300
3697
`,
			want: 89137,
		},

		"elves2": {
			content: `
8023
3736

40312

8527
4368
1291
1159
2869
10302
8489

3268
4846`,
			want: 89076,
		},

		// 40312
		// 21660
		// 15345
		// 11759
		// 8114
		"elves3": {
			content: `
8023
3736

40312

8527
4368
1291
1159

2869
10302
8489

3268
4846`,
			want: 77317,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day1.GetThreeFattiest(tc.content)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
