package day10_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day10"
)

func TestCheckAndSum(t *testing.T) {
	testCases := map[string]struct {
		reg         *day10.Register
		valBefore   int
		sumSignal   int
		checkSignal int
		want        int
		wantCheck   int
	}{
		"register0": {
			reg: &day10.Register{
				Value: 20,
				Cycle: 21,
			},
			valBefore:   21,
			sumSignal:   0,
			checkSignal: 20,
			want:        420,
			wantCheck:   60,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, gotCheck := day10.CheckAndSum(tc.reg, tc.valBefore, tc.sumSignal, tc.checkSignal)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum signal mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantCheck, gotCheck); diff != "" {
				t.Errorf("check signal mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSumSignal(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"register1": {
			input: `
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`,
			want: 13140,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day10.GetSumSignal(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum signal mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSpriteMove(t *testing.T) {
	testCases := map[string]struct {
		spr   *day10.Sprite
		input int
		want  string
	}{
		"sprite0": {
			spr: &day10.Sprite{
				Pos:  20,
				Draw: "###.....................................",
			},
			input: 16,
			want:  "...............###......................",
		},
		"sprite1": {
			spr: &day10.Sprite{
				Pos:  20,
				Draw: "###.....................................",
			},
			input: 0,
			want:  "##......................................",
		},

		"sprite2": {
			spr: &day10.Sprite{
				Pos:  20,
				Draw: "###.....................................",
			},
			input: 1,
			want:  "###.....................................",
		},
		"sprite3": {
			spr: &day10.Sprite{
				Pos:  20,
				Draw: "###.....................................",
			},
			input: 2,
			want:  ".###....................................",
		},
		"sprite4": {
			spr: &day10.Sprite{
				Pos:  20,
				Draw: "###.....................................",
			},
			input: 38,
			want:  ".....................................###",
		},
		"sprite5": {
			spr: &day10.Sprite{
				Pos:  20,
				Draw: "###.....................................",
			},
			input: 39,
			want:  "......................................##",
		},
		"sprite6": {
			spr: &day10.Sprite{
				Pos:  20,
				Draw: "###.....................................",
			},
			input: 40,
			want:  ".......................................#",
		},
		"sprite7": {
			spr: &day10.Sprite{
				Pos:  20,
				Draw: "###.....................................",
			},
			input: 41,
			want:  "........................................",
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			tc.spr.Move(tc.input)

			got := tc.spr.Draw

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sprite draw mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPaintPixel(t *testing.T) {
	testCases := map[string]struct {
		cycle int
		spr   *day10.Sprite
		crt   string
		want  string
	}{
		"pixel0": {
			cycle: 1,
			spr: &day10.Sprite{
				Pos:  1,
				Draw: "###.....................................",
			},
			crt:  "",
			want: "#",
		},
		"pixel1": {
			cycle: 2,
			spr: &day10.Sprite{
				Pos:  1,
				Draw: "###.....................................",
			},
			crt:  "#",
			want: "##",
		},
		"pixel2": {
			cycle: 3,
			spr: &day10.Sprite{
				Pos:  16,
				Draw: "................###.....................",
			},
			crt:  "##",
			want: "##.",
		},
		"pixel3": {
			cycle: 41,
			spr: &day10.Sprite{
				Pos:  1,
				Draw: "###.....................................",
			},
			crt:  "##..##..##..##..##..##..##..##..##..##..",
			want: "##..##..##..##..##..##..##..##..##..##..\n#",
		},
		"pixel4": {
			cycle: 81,
			spr: &day10.Sprite{
				Pos:  1,
				Draw: "###.....................................",
			},
			crt: `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.`,
			want: `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
#`,
		},
		"pixel5": {
			cycle: 40,
			spr: &day10.Sprite{
				Pos:  1,
				Draw: "###.....................................",
			},
			crt:  "##..##..##..##..##..##..##..##..##..##.",
			want: "##..##..##..##..##..##..##..##..##..##..",
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day10.PaintPixel(tc.cycle, tc.spr, tc.crt)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sprite draw mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPaintCRT(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  string
	}{
		"register1": {
			input: `
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`,
			want: `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day10.PaintCRT(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("sum signal mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
