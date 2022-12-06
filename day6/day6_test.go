package day6_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/vitorarins/aoc22/day6"
)

func TestIsValidMark(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  bool
	}{
		"signal0": {
			input: "bvwb",
			want:  false,
		},
		"signal1": {
			input: "vwbj",
			want:  true,
		},
		"signal2": {
			input: "nppd",
			want:  false,
		},
		"signal3": {
			input: "ppdv",
			want:  false,
		},
		"signal4": {
			input: "pdvj",
			want:  true,
		},
		"signal5": {
			input: "nznr",
			want:  false,
		},
		"signal6": {
			input: "znrn",
			want:  false,
		},
		"signal7": {
			input: "nrnf",
			want:  false,
		},
		"signal8": {
			input: "rnfr",
			want:  false,
		},
		"signal9": {
			input: "nfrf",
			want:  false,
		},
		"signal10": {
			input: "frfn",
			want:  false,
		},
		"signal11": {
			input: "rfnt",
			want:  true,
		},
		"signal12": {
			input: "zcfz",
			want:  false,
		},
		"signal13": {
			input: "cfzf",
			want:  false,
		},
		"signal14": {
			input: "fzfw",
			want:  false,
		},
		"signal15": {
			input: "zfwz",
			want:  false,
		},
		"signal16": {
			input: "fwzz",
			want:  false,
		},
		"signal17": {
			input: "wzzq",
			want:  false,
		},
		"signal18": {
			input: "zzqf",
			want:  false,
		},
		"signal19": {
			input: "zqfr",
			want:  true,
		},
		"signal20": {
			input: "mjqj",
			want:  false,
		},
		"signal21": {
			input: "jqjp",
			want:  false,
		},
		"signal22": {
			input: "qjpq",
			want:  false,
		},
		"signal23": {
			input: "jpqm",
			want:  true,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := day6.IsValidMark(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("signal valid mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetFirstMark(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"signal0": {
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
		},
		"signal1": {
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  6,
		},
		"signal2": {
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  10,
		},
		"signal3": {
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  11,
		},
		"signal4": {
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  7,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day6.GetFirstMark(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("first mark mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGetStartOfMsg(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  int
	}{
		"signal0": {
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  23,
		},
		"signal1": {
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  23,
		},
		"signal2": {
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  29,
		},
		"signal3": {
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  26,
		},
		"signal4": {
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  19,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got, err := day6.GetStartOfMsg(tc.input)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("first mark mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
