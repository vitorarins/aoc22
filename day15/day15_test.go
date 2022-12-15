package day15_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day15"
)

func TestGetNoBeaconsPos(t *testing.T) {
	testCases := map[string]struct {
		input string
		pos   int
		want  int
	}{
		"sensors0": {
			input: `
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`,
			pos:  10,
			want: 26,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day15.GetNoBeaconsPos(tc.input, tc.pos)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("no beacons pos mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetTuningFreqMax(t *testing.T) {
	testCases := map[string]struct {
		input string
		max   int
		want  int
	}{
		"sensors0": {
			input: `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`,
			max:  20,
			want: 56000011,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day15.GetTuningFreqMax(tc.input, tc.max)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("tuning frequency mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
