package day4_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day4"
)

func TestGetSections(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  map[int]struct{}
	}{
		"assign0": {
			input: "2-4",
			want: map[int]struct{}{
				2: {},
				3: {},
				4: {},
			},
		},
		"assign1": {
			input: "6-8",
			want: map[int]struct{}{
				6: {},
				7: {},
				8: {},
			},
		},
		"assign2": {
			input: "6-6",
			want: map[int]struct{}{
				6: {},
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day4.GetSections(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("bad assign mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsContained(t *testing.T) {
	testCases := map[string]struct {
		firstSections  map[int]struct{}
		secondSections map[int]struct{}
		want           bool
	}{
		"assign0": {
			firstSections: map[int]struct{}{
				2: {},
				3: {},
				4: {},
			},
			secondSections: map[int]struct{}{
				6: {},
				7: {},
				8: {},
			},
			want: false,
		},
		"assign1": {
			firstSections: map[int]struct{}{
				2: {},
				3: {},
			},
			secondSections: map[int]struct{}{
				4: {},
				5: {},
			},
			want: false,
		},
		"assign2": {
			firstSections: map[int]struct{}{
				5: {},
				6: {},
				7: {},
			},
			secondSections: map[int]struct{}{
				7: {},
				8: {},
				9: {},
			},
			want: false,
		},
		"assign3": {
			firstSections: map[int]struct{}{
				2: {},
				3: {},
				4: {},
				5: {},
				6: {},
				7: {},
				8: {},
			},
			secondSections: map[int]struct{}{
				3: {},
				4: {},
				5: {},
				6: {},
				7: {},
			},
			want: true,
		},
		"assign4": {
			firstSections: map[int]struct{}{
				4: {},
				5: {},
				6: {},
			},
			secondSections: map[int]struct{}{
				6: {},
			},
			want: true,
		},
		"assign5": {
			firstSections: map[int]struct{}{
				6: {},
			},
			secondSections: map[int]struct{}{
				4: {},
				5: {},
				6: {},
			},
			want: true,
		},
		"assign6": {
			firstSections: map[int]struct{}{
				2: {},
				3: {},
				4: {},
				5: {},
				6: {},
			},
			secondSections: map[int]struct{}{
				4: {},
				5: {},
				6: {},
				7: {},
				8: {},
			},
			want: false,
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day4.IsContained(tc.firstSections, tc.secondSections)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("bad assign mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsOverlap(t *testing.T) {
	testCases := map[string]struct {
		firstSections  map[int]struct{}
		secondSections map[int]struct{}
		want           bool
	}{
		"assign0": {
			firstSections: map[int]struct{}{
				2: {},
				3: {},
				4: {},
			},
			secondSections: map[int]struct{}{
				6: {},
				7: {},
				8: {},
			},
			want: false,
		},
		"assign1": {
			firstSections: map[int]struct{}{
				2: {},
				3: {},
			},
			secondSections: map[int]struct{}{
				4: {},
				5: {},
			},
			want: false,
		},
		"assign2": {
			firstSections: map[int]struct{}{
				5: {},
				6: {},
				7: {},
			},
			secondSections: map[int]struct{}{
				7: {},
				8: {},
				9: {},
			},
			want: true,
		},
		"assign3": {
			firstSections: map[int]struct{}{
				2: {},
				3: {},
				4: {},
				5: {},
				6: {},
				7: {},
				8: {},
			},
			secondSections: map[int]struct{}{
				3: {},
				4: {},
				5: {},
				6: {},
				7: {},
			},
			want: true,
		},
		"assign4": {
			firstSections: map[int]struct{}{
				4: {},
				5: {},
				6: {},
			},
			secondSections: map[int]struct{}{
				6: {},
			},
			want: true,
		},
		"assign5": {
			firstSections: map[int]struct{}{
				6: {},
			},
			secondSections: map[int]struct{}{
				4: {},
				5: {},
				6: {},
			},
			want: true,
		},
		"assign6": {
			firstSections: map[int]struct{}{
				2: {},
				3: {},
				4: {},
				5: {},
				6: {},
			},
			secondSections: map[int]struct{}{
				4: {},
				5: {},
				6: {},
				7: {},
				8: {},
			},
			want: true,
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day4.IsOverlap(tc.firstSections, tc.secondSections)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("bad assign mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsBadAssign(t *testing.T) {
	testCases := map[string]struct {
		input string
		ckFn  day4.CheckBadFn
		want  bool
	}{
		"assignContained0": {
			input: "2-4,6-8",
			ckFn:  day4.IsContained,
			want:  false,
		},
		"assignContained1": {
			input: "2-3,4-5",
			ckFn:  day4.IsContained,
			want:  false,
		},
		"assignContained2": {
			input: "5-7,7-9",
			ckFn:  day4.IsContained,
			want:  false,
		},
		"assignContained3": {
			input: "2-8,3-7",
			ckFn:  day4.IsContained,
			want:  true,
		},
		"assignContained4": {
			input: "6-6,4-6",
			ckFn:  day4.IsContained,
			want:  true,
		},
		"assignContained5": {
			input: "2-6,4-8",
			ckFn:  day4.IsContained,
			want:  false,
		},
		"assignOverlap0": {
			input: "2-4,6-8",
			ckFn:  day4.IsOverlap,
			want:  false,
		},
		"assignOverlap1": {
			input: "2-3,4-5",
			ckFn:  day4.IsOverlap,
			want:  false,
		},
		"assignOverlap2": {
			input: "5-7,7-9",
			ckFn:  day4.IsOverlap,
			want:  true,
		},
		"assignOverlap3": {
			input: "2-8,3-7",
			ckFn:  day4.IsOverlap,
			want:  true,
		},
		"assignOverlap4": {
			input: "6-6,4-6",
			ckFn:  day4.IsOverlap,
			want:  true,
		},
		"assignOverlap5": {
			input: "2-6,4-8",
			ckFn:  day4.IsOverlap,
			want:  true,
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day4.IsBadAssign(tc.input, tc.ckFn)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("bad assign mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetBadAssignsAmnt(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"assigns": {
			input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`,
			want: 2,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day4.GetBadAssignsAmnt(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("qty of bad assigns mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetOverlap(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"assigns": {
			input: `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`,
			want: 4,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day4.GetOverlap(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("qty of bad assigns mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
