package day2

import (
	"strings"
)

type Match struct {
	points map[string]int
}

func GetPlayScore(handMatch map[string]Match, play string) int {

	hands := strings.Split(play, " ")

	oponent := hands[0]
	strategy := hands[1]

	return handMatch[strategy].points[oponent]
}

func GetScoreSpec(content string) (int, error) {

	handMatch := map[string]Match{
		"X": {
			points: map[string]int{
				"A": 4,
				"B": 1,
				"C": 7,
			},
		},
		"Y": {
			points: map[string]int{
				"A": 8,
				"B": 5,
				"C": 2,
			},
		},
		"Z": {
			points: map[string]int{
				"A": 3,
				"B": 9,
				"C": 6,
			},
		},
	}

	content = strings.TrimSpace(content)

	plays := strings.Split(content, "\n")

	score := 0

	for _, play := range plays {
		play = strings.TrimSpace(play)

		score += GetPlayScore(handMatch, play)
	}

	return score, nil
}

func GetScore(content string) (int, error) {

	handMatch := map[string]Match{
		"X": {
			points: map[string]int{
				"A": 3,
				"B": 1,
				"C": 2,
			},
		},
		"Y": {
			points: map[string]int{
				"A": 4,
				"B": 5,
				"C": 6,
			},
		},
		"Z": {
			points: map[string]int{
				"A": 8,
				"B": 9,
				"C": 7,
			},
		},
	}

	content = strings.TrimSpace(content)

	plays := strings.Split(content, "\n")

	score := 0

	for _, play := range plays {
		play = strings.TrimSpace(play)

		score += GetPlayScore(handMatch, play)
	}

	return score, nil
}
