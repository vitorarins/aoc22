package day9_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day9"
)

func TestUpdatePos(t *testing.T) {
	testCases := map[string]struct {
		pos       map[day9.Point]struct{}
		direction string
		steps     int
		head      day9.Point
		tail      day9.Point
		want      map[day9.Point]struct{}
		wantHead  day9.Point
		wantTail  day9.Point
	}{
		"pos0": {
			// R 4
			// ......
			// ......
			// ......
			// ......
			// s..TH.
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "R",
			steps:     4,
			head:      day9.Point{0, 0},
			tail:      day9.Point{0, 0},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
				{1, 0}: struct{}{},
				{2, 0}: struct{}{},
				{3, 0}: struct{}{},
			},
			wantHead: day9.Point{4, 0},
			wantTail: day9.Point{3, 0},
		},
		"pos1": {
			// U 4
			// ....H.
			// ....T.
			// ......
			// ......
			// s.....
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
				{1, 0}: struct{}{},
				{2, 0}: struct{}{},
				{3, 0}: struct{}{},
			},
			direction: "U",
			steps:     4,
			head:      day9.Point{4, 0},
			tail:      day9.Point{3, 0},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
			},
			wantHead: day9.Point{4, -4},
			wantTail: day9.Point{4, -3},
		},
		"pos2": {
			// L 3
			// .HT...
			// ......
			// ......
			// ......
			// s.....
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
			},
			direction: "L",
			steps:     3,
			head:      day9.Point{4, -4},
			tail:      day9.Point{4, -3},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
			},
			wantHead: day9.Point{1, -4},
			wantTail: day9.Point{2, -4},
		},
		"pos3": {
			// D 1
			// ..T...
			// .H....
			// ......
			// ......
			// s.....
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
			},
			direction: "D",
			steps:     1,
			head:      day9.Point{1, -4},
			tail:      day9.Point{2, -4},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
			},
			wantHead: day9.Point{1, -3},
			wantTail: day9.Point{2, -4},
		},
		"pos4": {
			// R 4
			// ......
			// ....TH
			// ......
			// ......
			// s.....
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
			},
			direction: "R",
			steps:     4,
			head:      day9.Point{1, -3},
			tail:      day9.Point{2, -4},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
				{3, -3}: struct{}{},
			},
			wantHead: day9.Point{5, -3},
			wantTail: day9.Point{4, -3},
		},
		"pos5": {
			// D 1
			// ......
			// ....T.
			// .....H
			// ......
			// s.....
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
				{3, -3}: struct{}{},
			},
			direction: "D",
			steps:     1,
			head:      day9.Point{5, -3},
			tail:      day9.Point{4, -3},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
				{3, -3}: struct{}{},
			},
			wantHead: day9.Point{5, -2},
			wantTail: day9.Point{4, -3},
		},
		"pos6": {
			// L 5
			// ......
			// ......
			// HT....
			// ......
			// s.....
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
				{3, -3}: struct{}{},
			},
			direction: "L",
			steps:     5,
			head:      day9.Point{5, -2},
			tail:      day9.Point{4, -3},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
				{3, -3}: struct{}{},
				{3, -2}: struct{}{},
				{2, -2}: struct{}{},
				{1, -2}: struct{}{},
			},
			wantHead: day9.Point{0, -2},
			wantTail: day9.Point{1, -2},
		},
		"pos7": {
			// R 2
			// ......
			// ......
			// .TH...
			// ......
			// s.....
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
				{3, -3}: struct{}{},
				{3, -2}: struct{}{},
				{2, -2}: struct{}{},
				{1, -2}: struct{}{},
			},
			direction: "R",
			steps:     2,
			head:      day9.Point{0, -2},
			tail:      day9.Point{1, -2},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, 0}:  struct{}{},
				{2, 0}:  struct{}{},
				{3, 0}:  struct{}{},
				{4, -1}: struct{}{},
				{4, -2}: struct{}{},
				{4, -3}: struct{}{},
				{3, -4}: struct{}{},
				{2, -4}: struct{}{},
				{3, -3}: struct{}{},
				{3, -2}: struct{}{},
				{2, -2}: struct{}{},
				{1, -2}: struct{}{},
			},
			wantHead: day9.Point{2, -2},
			wantTail: day9.Point{1, -2},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, gotHead, gotTail := day9.UpdatePos(tc.pos, tc.direction, tc.steps, tc.head, tc.tail)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("positions mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantHead, gotHead); diff != "" {
				t.Errorf("head mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantTail, gotTail); diff != "" {
				t.Errorf("tail mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetPos(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"bridge": {
			input: `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`,
			want: 13,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day9.GetPos(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("positions mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// == R 4 ==

// ......
// ......
// ......
// ......
// 1H....  (1 covers 2, 3, 4, 5, 6, 7, 8, 9, s)

// ......
// ......
// ......
// ......
// 21H...  (2 covers 3, 4, 5, 6, 7, 8, 9, s)

// ......
// ......
// ......
// ......
// 321H..  (3 covers 4, 5, 6, 7, 8, 9, s)

// ......
// ......
// ......
// ......
// 4321H.

func TestUpdatePosManyHeads(t *testing.T) {
	testCases := map[string]struct {
		pos       map[day9.Point]struct{}
		direction string
		steps     int
		heads     []day9.Point
		want      map[day9.Point]struct{}
		wantHeads []day9.Point
	}{
		"pos0": {
			// R 4
			// ......
			// ......
			// ......
			// ......
			// 4321H.
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "R",
			steps:     4,
			heads:     []day9.Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		"pos1": {
			// U 4
			// ....H.
			// ....1.
			// ..432.
			// .5....
			// 6.....
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "U",
			steps:     4,
			heads:     []day9.Point{{4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{4, -4}, {4, -3}, {4, -2}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		"pos2": {

			// L 3
			// .H1...
			// ...2..
			// ..43..
			// .5....
			// 6.....

			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "L",
			steps:     3,
			heads:     []day9.Point{{4, -4}, {4, -3}, {4, -2}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{1, -4}, {2, -4}, {3, -3}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		"pos3": {
			// D 1
			// ..1...
			// .H.2..
			// ..43..
			// .5....
			// 6.....
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "D",
			steps:     1,
			heads:     []day9.Point{{1, -4}, {2, -4}, {3, -3}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{1, -3}, {2, -4}, {3, -3}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		"pos4": {
			// R 4
			// ......
			// ...21H
			// ..43..
			// .5....
			// 6.....
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "R",
			steps:     4,
			heads:     []day9.Point{{1, -3}, {2, -4}, {3, -3}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{5, -3}, {4, -3}, {3, -3}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		"pos5": {
			// D 1
			// ......
			// ...21.
			// ..43.H
			// .5....
			// 6.....
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "D",
			steps:     1,
			heads:     []day9.Point{{5, -3}, {4, -3}, {3, -3}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{5, -2}, {4, -3}, {3, -3}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		"pos6": {
			// L 5
			// ......
			// ......
			// H123..  (2 covers 4)
			// .5....
			// 6.....
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "L",
			steps:     5,
			heads:     []day9.Point{{5, -2}, {4, -3}, {3, -3}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{0, -2}, {1, -2}, {2, -2}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		"pos7": {
			// R 2
			// ......
			// ......
			// .1H3..  (H covers 2, 4)
			// .5....
			// 6.....
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "R",
			steps:     2,
			heads:     []day9.Point{{0, -2}, {1, -2}, {2, -2}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{2, -2}, {1, -2}, {2, -2}, {3, -2}, {2, -2}, {1, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		// R 5
		// U 8
		// L 8
		// D 3
		// R 17
		// D 10
		// L 25
		// U 20
		"posLarger0": {
			// R 5

			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "R",
			steps:     5,
			heads:     []day9.Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{5, 0}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		"posLarger1": {
			// U 8
			// .....H
			// .....1
			// .....2
			// .....3
			// ....54
			// ...6..
			// ..7...
			// .8....
			// 9.....
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "U",
			steps:     8,
			heads:     []day9.Point{{5, 0}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{5, -8}, {5, -7}, {5, -6}, {5, -5}, {5, -4}, {4, -4}, {3, -3}, {2, -2}, {1, -1}, {0, 0}},
		},
		"posLarger2": {
			// L 8
			//
			// H1234
			// ....5
			// ....6
			// ....7
			// ....8
			// ....9
			// .....
			// .....
			// ...s.
			pos: map[day9.Point]struct{}{
				{0, 0}: struct{}{},
			},
			direction: "L",
			steps:     8,
			heads:     []day9.Point{{5, -8}, {5, -7}, {5, -6}, {5, -5}, {5, -4}, {4, -4}, {3, -3}, {2, -2}, {1, -1}, {0, 0}},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, -1}: struct{}{},
				{2, -2}: struct{}{},
				{1, -3}: struct{}{},
			},
			wantHeads: []day9.Point{{-3, -8}, {-2, -8}, {-1, -8}, {0, -8}, {1, -8}, {1, -7}, {1, -6}, {1, -5}, {1, -4}, {1, -3}},
		},
		"posLarger3": {
			// D 3
			//
			// .2345
			// 1...6
			// H...7
			// ....8
			// ....9
			// .....
			// .....
			// ...s.
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, -1}: struct{}{},
				{2, -2}: struct{}{},
				{1, -3}: struct{}{},
			},
			direction: "D",
			steps:     3,
			heads:     []day9.Point{{-3, -8}, {-2, -8}, {-1, -8}, {0, -8}, {1, -8}, {1, -7}, {1, -6}, {1, -5}, {1, -4}, {1, -3}},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, -1}: struct{}{},
				{2, -2}: struct{}{},
				{1, -3}: struct{}{},
			},
			wantHeads: []day9.Point{{-3, -5}, {-3, -6}, {-2, -7}, {-1, -7}, {0, -7}, {1, -7}, {1, -6}, {1, -5}, {1, -4}, {1, -3}},
		},
		"posLarger4": {
			// R 17
			//
			// .....987654321H
			// ...............
			// ...............
			// ...............
			// ...............
			// s..............
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, -1}: struct{}{},
				{2, -2}: struct{}{},
				{1, -3}: struct{}{},
			},
			direction: "R",
			steps:     17,
			heads:     []day9.Point{{-3, -5}, {-3, -6}, {-2, -7}, {-1, -7}, {0, -7}, {1, -7}, {1, -6}, {1, -5}, {1, -4}, {1, -3}},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, -1}: struct{}{},
				{2, -2}: struct{}{},
				{1, -3}: struct{}{},
				{2, -4}: struct{}{},
				{3, -5}: struct{}{},
				{4, -5}: struct{}{},
				{5, -5}: struct{}{},
			},
			wantHeads: []day9.Point{{14, -5}, {13, -5}, {12, -5}, {11, -5}, {10, -5}, {9, -5}, {8, -5}, {7, -5}, {6, -5}, {5, -5}},
		},
		"posLarger5": {
			// D 10
			//
			// s.........98765
			// ..............4
			// ..............3
			// ..............2
			// ..............1
			// ..............H
			pos: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, -1}: struct{}{},
				{2, -2}: struct{}{},
				{1, -3}: struct{}{},
				{2, -4}: struct{}{},
				{3, -5}: struct{}{},
				{4, -5}: struct{}{},
				{5, -5}: struct{}{},
			},
			direction: "D",
			steps:     10,
			heads:     []day9.Point{{14, -5}, {13, -5}, {12, -5}, {11, -5}, {10, -5}, {9, -5}, {8, -5}, {7, -5}, {6, -5}, {5, -5}},
			want: map[day9.Point]struct{}{
				{0, 0}:  struct{}{},
				{1, -1}: struct{}{},
				{2, -2}: struct{}{},
				{1, -3}: struct{}{},
				{2, -4}: struct{}{},
				{3, -5}: struct{}{},
				{4, -5}: struct{}{},
				{5, -5}: struct{}{},
				{6, -4}: struct{}{},
				{7, -3}: struct{}{},
				{8, -2}: struct{}{},
				{9, -1}: struct{}{},
				{10, 0}: struct{}{},
			},
			wantHeads: []day9.Point{{14, 5}, {14, 4}, {14, 3}, {14, 2}, {14, 1}, {14, 0}, {13, 0}, {12, 0}, {11, 0}, {10, 0}},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, gotHeads := day9.UpdatePosManyHeads(tc.pos, tc.direction, tc.steps, tc.heads)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("positions mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantHeads, gotHeads); diff != "" {
				t.Errorf("heads mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCheckTail(t *testing.T) {
	testCases := map[string]struct {
		head day9.Point
		tail day9.Point
		want day9.Point
	}{
		// Step: 0
		"move0": {
			head: day9.Point{4, -1},
			tail: day9.Point{3, 0},
			want: day9.Point{3, 0},
		},
		"move1": {
			head: day9.Point{3, 0},
			tail: day9.Point{2, 0},
			want: day9.Point{2, 0},
		},
		"move2": {
			head: day9.Point{2, 0},
			tail: day9.Point{1, 0},
			want: day9.Point{1, 0},
		},
		"move3": {
			head: day9.Point{1, 0},
			tail: day9.Point{0, 0},
			want: day9.Point{0, 0},
		},
		"move4": {
			head: day9.Point{0, 0},
			tail: day9.Point{0, 0},
			want: day9.Point{0, 0},
		},
		// Step: 1
		"move5": {
			head: day9.Point{4, -2},
			tail: day9.Point{3, 0},
			want: day9.Point{4, -1},
		},
		"move6": {
			head: day9.Point{4, -1},
			tail: day9.Point{2, 0},
			want: day9.Point{3, -1},
		},
		"move7": {
			head: day9.Point{3, -1},
			tail: day9.Point{1, 0},
			want: day9.Point{2, -1},
		},
		"move8": {
			head: day9.Point{2, -1},
			tail: day9.Point{0, 0},
			want: day9.Point{1, -1},
		},
		"move9": {
			head: day9.Point{1, -1},
			tail: day9.Point{0, 0},
			want: day9.Point{0, 0},
		},
		// Step: 2
		"move10": {
			head: day9.Point{4, -3},
			tail: day9.Point{4, -1},
			want: day9.Point{4, -2},
		},
		"move11": {
			head: day9.Point{4, -2},
			tail: day9.Point{3, -1},
			want: day9.Point{3, -1},
		},
		"move12": {
			head: day9.Point{3, -1},
			tail: day9.Point{2, -1},
			want: day9.Point{2, -1},
		},
		"move13": {
			head: day9.Point{2, -1},
			tail: day9.Point{1, -1},
			want: day9.Point{1, -1},
		},
		// Step: 3
		"move14": {
			head: day9.Point{4, -4},
			tail: day9.Point{4, -2},
			want: day9.Point{4, -3},
		},
		"move15": {
			head: day9.Point{4, -3},
			tail: day9.Point{3, -1},
			want: day9.Point{4, -2},
		},
		"move16": {
			head: day9.Point{4, -2},
			tail: day9.Point{2, -1},
			want: day9.Point{3, -2},
		},
		"move17": {
			head: day9.Point{3, -2},
			tail: day9.Point{1, -1},
			want: day9.Point{2, -2},
		},
		"move18": {
			head: day9.Point{2, -2},
			tail: day9.Point{0, 0},
			want: day9.Point{1, -1},
		},
		"move19": {
			head: day9.Point{2, -2},
			tail: day9.Point{4, -3},
			want: day9.Point{3, -2},
		},
		"move20": {
			head: day9.Point{3, -5},
			tail: day9.Point{5, -6},
			want: day9.Point{4, -5},
		},
		"move21": {
			head: day9.Point{4, -3},
			tail: day9.Point{2, -4},
			want: day9.Point{3, -3},
		},
		"move22": {
			head: day9.Point{3, -4},
			tail: day9.Point{4, -2},
			want: day9.Point{3, -3},
		},
		"move23": {
			head: day9.Point{3, -4},
			tail: day9.Point{4, -3},
			want: day9.Point{4, -3},
		},
		"move24": {
			head: day9.Point{2, -4},
			tail: day9.Point{4, -3},
			want: day9.Point{3, -4},
		},
		"move25": {
			head: day9.Point{-3, -6},
			tail: day9.Point{-2, -8},
			want: day9.Point{-3, -7},
		},
		"move26": {
			head: day9.Point{14, -3},
			tail: day9.Point{13, -5},
			want: day9.Point{14, -4},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day9.CheckTail(tc.head, tc.tail)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("tail mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetPosManyHeads(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"bridge0": {
			input: `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`,
			want: 1,
		},
		"bridge1": {
			input: `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`,
			want: 36,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day9.GetPosManyHeads(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("positions mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
