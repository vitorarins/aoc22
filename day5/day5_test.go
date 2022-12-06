package day5_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day5"
)

func TestParseStackLine(t *testing.T) {
	testCases := map[string]struct {
		stacks  map[string]day5.Stack
		numbers []string
		line    string
		want    map[string]day5.Stack
	}{
		"stack": {
			stacks:  map[string]day5.Stack{},
			numbers: []string{"1", "2", "3"},
			line:    "[Z] [M] [P]",
			want: map[string]day5.Stack{
				"1": {"Z"},
				"2": {"M"},
				"3": {"P"},
			},
		},
		"stack1": {
			stacks: map[string]day5.Stack{
				"1": {"Z"},
				"2": {"M"},
				"3": {"P"},
			},
			numbers: []string{"1", "2", "3"},
			line:    "[N] [C]    ",
			want: map[string]day5.Stack{
				"1": {"Z", "N"},
				"2": {"M", "C"},
				"3": {"P"},
			},
		},
		"stack2": {
			stacks: map[string]day5.Stack{
				"1": {"Z", "N"},
				"2": {"M", "C"},
				"3": {"P"},
			},
			numbers: []string{"1", "2", "3"},
			line:    "    [D]    ",
			want: map[string]day5.Stack{
				"1": {"Z", "N"},
				"2": {"M", "C", "D"},
				"3": {"P"},
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day5.ParseStackLine(tc.stacks, tc.numbers, tc.line)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("stacks mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestParseStacks(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  map[string]day5.Stack
	}{
		"stacks": {
			input: `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 
`,
			want: map[string]day5.Stack{
				"1": {"Z", "N"},
				"2": {"M", "C", "D"},
				"3": {"P"},
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day5.ParseStacks(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("stacks mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestParseInstructLine(t *testing.T) {
	testCases := map[string]struct {
		stacks map[string]day5.Stack
		input  string
		moveFn day5.MoveFunction
		want   map[string]day5.Stack
	}{
		"stacksMoveOne": {
			stacks: map[string]day5.Stack{
				"1": {"Z", "N"},
				"2": {"M", "C", "D"},
				"3": {"P"},
			},
			input:  "move 1 from 2 to 1",
			moveFn: day5.MoveOne,
			want: map[string]day5.Stack{
				"1": {"Z", "N", "D"},
				"2": {"M", "C"},
				"3": {"P"},
			},
		},
		"stacks1MoveOne": {
			stacks: map[string]day5.Stack{
				"1": {"Z", "N", "D"},
				"2": {"M", "C"},
				"3": {"P"},
			},
			input:  "move 3 from 1 to 3",
			moveFn: day5.MoveOne,
			want: map[string]day5.Stack{
				"1": {},
				"2": {"M", "C"},
				"3": {"P", "D", "N", "Z"},
			},
		},
		"stacks2MoveOne": {
			stacks: map[string]day5.Stack{
				"1": {},
				"2": {"M", "C"},
				"3": {"P", "D", "N", "Z"},
			},
			input:  "move 2 from 2 to 1",
			moveFn: day5.MoveOne,
			want: map[string]day5.Stack{
				"1": {"C", "M"},
				"2": {},
				"3": {"P", "D", "N", "Z"},
			},
		},
		"stacks3MoveOne": {
			stacks: map[string]day5.Stack{
				"1": {"C", "M"},
				"2": {},
				"3": {"P", "D", "N", "Z"},
			},
			input:  "move 1 from 1 to 2",
			moveFn: day5.MoveOne,
			want: map[string]day5.Stack{
				"1": {"C"},
				"2": {"M"},
				"3": {"P", "D", "N", "Z"},
			},
		},
		"stacksMoveAll": {
			stacks: map[string]day5.Stack{
				"1": {"Z", "N"},
				"2": {"M", "C", "D"},
				"3": {"P"},
			},
			input:  "move 1 from 2 to 1",
			moveFn: day5.MoveAll,
			want: map[string]day5.Stack{
				"1": {"Z", "N", "D"},
				"2": {"M", "C"},
				"3": {"P"},
			},
		},
		"stacks1MoveAll": {
			stacks: map[string]day5.Stack{
				"1": {"Z", "N", "D"},
				"2": {"M", "C"},
				"3": {"P"},
			},
			input:  "move 3 from 1 to 3",
			moveFn: day5.MoveAll,
			want: map[string]day5.Stack{
				"1": {},
				"2": {"M", "C"},
				"3": {"P", "Z", "N", "D"},
			},
		},
		"stacks2MoveAll": {
			stacks: map[string]day5.Stack{
				"1": {},
				"2": {"M", "C"},
				"3": {"P", "Z", "N", "D"},
			},
			input:  "move 2 from 2 to 1",
			moveFn: day5.MoveAll,
			want: map[string]day5.Stack{
				"1": {"M", "C"},
				"2": {},
				"3": {"P", "Z", "N", "D"},
			},
		},
		"stacks3MoveAll": {
			stacks: map[string]day5.Stack{
				"1": {"M", "C"},
				"2": {},
				"3": {"P", "Z", "N", "D"},
			},
			input:  "move 1 from 1 to 2",
			moveFn: day5.MoveAll,
			want: map[string]day5.Stack{
				"1": {"M"},
				"2": {"C"},
				"3": {"P", "Z", "N", "D"},
			},
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day5.ParseInstructLine(tc.stacks, tc.input, tc.moveFn)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("stacks mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestParseInstructs(t *testing.T) {
	testCases := map[string]struct {
		stacks map[string]day5.Stack
		input  string
		moveFn day5.MoveFunction
		want   string
	}{
		"stacksMoveOne": {
			stacks: map[string]day5.Stack{
				"1": {"Z", "N"},
				"2": {"M", "C", "D"},
				"3": {"P"},
			},
			input: `
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`,
			moveFn: day5.MoveOne,
			want:   "CMZ",
		},
		"stacksMoveAll": {
			stacks: map[string]day5.Stack{
				"1": {"Z", "N"},
				"2": {"M", "C", "D"},
				"3": {"P"},
			},
			input: `
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`,
			moveFn: day5.MoveAll,
			want:   "MCD",
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day5.ParseInstructs(tc.stacks, tc.input, tc.moveFn)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("stacks mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetStacksTopMoveOne(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  string
	}{
		"stacks": {
			input: `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`,
			want: "CMZ",
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day5.GetStacksTopMoveOne(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("stacks top mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetStacksTopMoveAll(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  string
	}{
		"stacks": {
			input: `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`,
			want: "MCD",
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day5.GetStacksTopMoveAll(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("stacks top mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
