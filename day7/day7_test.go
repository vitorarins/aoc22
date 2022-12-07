package day7_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day7"
)

func TestParseCommand(t *testing.T) {
	testCases := map[string]struct {
		rootAndCurrent []*day7.Directory
		input          string
		want           *day7.Directory
	}{
		"command0": {
			rootAndCurrent: func() []*day7.Directory {
				root := day7.NewDirectory("/", nil)

				return []*day7.Directory{root, root}
			}(),
			input: "$ cd /",
			want: &day7.Directory{
				Name:     "/",
				Size:     0,
				Parent:   nil,
				Children: map[string]*day7.Directory{},
			},
		},
		"command1": {
			rootAndCurrent: func() []*day7.Directory {
				root := day7.NewDirectory("/", nil)

				return []*day7.Directory{root, root}
			}(),
			input: "$ ls",
			want: &day7.Directory{
				Name:     "/",
				Size:     0,
				Parent:   nil,
				Children: map[string]*day7.Directory{},
			},
		},
		"command2": {
			rootAndCurrent: func() []*day7.Directory {
				root := day7.NewDirectory("/", nil)

				return []*day7.Directory{root, root}
			}(),
			input: "$ cd a",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				result := day7.NewDirectory("a", root)
				root.AddDirectory(result)

				return result
			}(),
		},
		"command3": {
			rootAndCurrent: func() []*day7.Directory {
				root := day7.NewDirectory("/", nil)

				// cd a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				return []*day7.Directory{root, a}
			}(),
			input: "$ cd e",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// cd a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// cd e
				result := day7.NewDirectory("e", a)
				a.AddDirectory(result)

				return result
			}(),
		},
		"command4": {
			rootAndCurrent: func() []*day7.Directory {
				root := day7.NewDirectory("/", nil)

				// cd a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// cd e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				return []*day7.Directory{root, e}
			}(),
			input: "$ cd ..",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// cd a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// cd e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// cd ..
				return a
			}(),
		},
		"command5": {
			rootAndCurrent: func() []*day7.Directory {
				root := day7.NewDirectory("/", nil)

				// cd a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// cd e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// cd ..
				return []*day7.Directory{root, a}
			}(),
			input: "$ cd ..",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// cd a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// cd e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// cd ..
				// cd ..
				return root
			}(),
		},
		"command6": {
			rootAndCurrent: func() []*day7.Directory {
				root := day7.NewDirectory("/", nil)

				// cd a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// cd e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// cd ..
				// cd ..
				return []*day7.Directory{root, root}
			}(),
			input: "$ cd d",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// cd a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// cd e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// cd ..
				// cd ..
				// cd d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				return d
			}(),
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day7.ParseCommand(tc.rootAndCurrent[0], tc.rootAndCurrent[1], tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("parse command mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestParseListOutput(t *testing.T) {
	testCases := map[string]struct {
		current *day7.Directory
		input   string
		want    *day7.Directory
	}{
		"lsOutput0": {
			current: day7.NewDirectory("/", nil),
			input:   "dir a",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				return root
			}(),
		},
		"lsOutput1": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				return root
			}(),
			input: "14848514 b.txt",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				root.AddFileSize(14848514)

				return root
			}(),
		},
		"lsOutput2": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				return root
			}(),
			input: "8504156 c.dat",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				return root
			}(),
		},
		"lsOutput3": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				return root
			}(),
			input: "dir d",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				return root
			}(),
		},
		"lsOutput4": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				return a
			}(),
			input: "dir e",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				return a
			}(),
		},
		"lsOutput5": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				return a
			}(),
			input: "29116 f",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				return a
			}(),
		},
		"lsOutput6": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				return a
			}(),
			input: "2557 g",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				return a
			}(),
		},
		"lsOutput7": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				return a
			}(),
			input: "62596 h.lst",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				return a
			}(),
		},
		"lsOutput8": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				return e
			}(),
			input: "584 i",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				return e
			}(),
		},
		"lsOutput9": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				return d
			}(),
			input: "4060174 j",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				// 4060174 j
				d.AddFileSize(4060174)

				return d
			}(),
		},
		"lsOutput10": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				// 4060174 j
				d.AddFileSize(4060174)

				return d
			}(),
			input: "8033020 d.log",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				// 4060174 j
				d.AddFileSize(4060174)

				// 8033020 d.log
				d.AddFileSize(8033020)

				return d
			}(),
		},
		"lsOutput11": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				// 4060174 j
				d.AddFileSize(4060174)

				// 8033020 d.log
				d.AddFileSize(8033020)

				return d
			}(),
			input: "5626152 d.ext",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				// 4060174 j
				d.AddFileSize(4060174)

				// 8033020 d.log
				d.AddFileSize(8033020)

				// 5626152 d.ext
				d.AddFileSize(5626152)

				return d
			}(),
		},
		"lsOutput12": {
			current: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				// 4060174 j
				d.AddFileSize(4060174)

				// 8033020 d.log
				d.AddFileSize(8033020)

				// 5626152 d.ext
				d.AddFileSize(5626152)

				return d
			}(),
			input: "7214296 k",
			want: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				// 4060174 j
				d.AddFileSize(4060174)

				// 8033020 d.log
				d.AddFileSize(8033020)

				// 5626152 d.ext
				d.AddFileSize(5626152)

				// 7214296 k
				d.AddFileSize(7214296)

				return d
			}(),
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day7.ParseListOutput(tc.current, tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("parse list output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetSumUpTo(t *testing.T) {
	testCases := map[string]struct {
		root  *day7.Directory
		limit int
		want  int
	}{
		"fs": {
			root: func() *day7.Directory {
				root := day7.NewDirectory("/", nil)

				// dir a
				a := day7.NewDirectory("a", root)
				root.AddDirectory(a)

				// 14848514 b.txt
				root.AddFileSize(14848514)

				// 8504156 c.dat
				root.AddFileSize(8504156)

				// dir d
				d := day7.NewDirectory("d", root)
				root.AddDirectory(d)

				// cd a
				// dir e
				e := day7.NewDirectory("e", a)
				a.AddDirectory(e)

				// 29116 f
				a.AddFileSize(29116)

				// 2557 g
				a.AddFileSize(2557)

				// 62596 h.lst
				a.AddFileSize(62596)

				// cd e
				// 584 i
				e.AddFileSize(584)

				// cd ..
				// cd ..
				// cd d
				// 4060174 j
				d.AddFileSize(4060174)

				// 8033020 d.log
				d.AddFileSize(8033020)

				// 5626152 d.ext
				d.AddFileSize(5626152)

				// 7214296 k
				d.AddFileSize(7214296)

				return root
			}(),
			limit: 100000,
			want:  95437,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day7.GetSumUpTo(tc.root, tc.limit)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("file sizes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetFileSizes(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"fs": {
			input: `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
			want: 95437,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day7.GetFileSizes(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("file sizes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetMinDir(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"fs": {
			input: `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
			want: 24933642,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day7.GetMinDir(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("file sizes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
