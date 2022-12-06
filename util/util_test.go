package util_test

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/util"
)

func TestReadInput(t *testing.T) {
	testCases := map[string]struct {
		fsys     fs.FS
		filename string
		want     string
	}{
		"ok": {
			fsys: fstest.MapFS{
				"input": {
					Data: []byte("hello, world"),
				},
			},
			filename: "input",
			want:     "hello, world",
		},
		"strategy": {
			fsys: fstest.MapFS{
				"input": {
					Data: []byte(`C Y
C Y
C Y
B Z
C Y
`),
				},
			},
			filename: "input",
			want: `C Y
C Y
C Y
B Z
C Y
`,
		},
		"elves": {
			fsys: fstest.MapFS{
				"input": {
					Data: []byte(`
7183
3394

25380
25324
16859

7300
3697
`),
				},
			},
			filename: "input",
			want: `
7183
3394

25380
25324
16859

7300
3697
`,
		},
		"stacks": {
			fsys: fstest.MapFS{
				"input": {
					Data: []byte(`
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`),
				},
			},
			filename: "input",
			want: `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := util.ReadInput(tc.fsys, tc.filename)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("content mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
