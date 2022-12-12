package day9

import (
	"strconv"
	"strings"
)

// Point represents a position on the grid.
type Point struct {
	X int
	Y int
}

func MoveHead(direction string, head Point) Point {
	switch direction {
	case "R":
		head.X += 1
	case "L":
		head.X -= 1
	case "U":
		head.Y -= 1
	case "D":
		head.Y += 1
	}

	return head
}

func CheckTail(head, tail Point) Point {

	switch {
	case head.X == tail.X+2 && head.Y == tail.Y:
		tail.X++

	case head.X == tail.X-2 && head.Y == tail.Y:
		tail.X--

	case head.Y == tail.Y+2 && head.X == tail.X:
		tail.Y++

	case head.Y == tail.Y-2 && head.X == tail.X:
		tail.Y--

	case head.X == tail.X+2 && head.Y == tail.Y+2:
		tail.X++
		tail.Y++

	case head.X == tail.X-2 && head.Y == tail.Y-2:
		tail.X--
		tail.Y--

	case head.X == tail.X-2 && head.Y == tail.Y+2:
		tail.X--
		tail.Y++

	case head.X == tail.X+2 && head.Y == tail.Y-2:
		tail.X++
		tail.Y--

	case head.X == tail.X+1 && head.Y == tail.Y-2:
		tail.X++
		tail.Y--

	case head.X == tail.X+2 && head.Y == tail.Y-1:
		tail.X++
		tail.Y--

	case head.X == tail.X-2 && head.Y == tail.Y+1:
		tail.X--
		tail.Y++

	case head.X == tail.X+2 && head.Y == tail.Y+1:
		tail.X++
		tail.Y++

	case head.X == tail.X-2 && head.Y == tail.Y-1:
		tail.X--
		tail.Y--

	case head.X == tail.X-1 && head.Y == tail.Y-2:
		tail.X--
		tail.Y--

	case head.X == tail.X-1 && head.Y == tail.Y+2:
		tail.X--
		tail.Y++

	case head.X == tail.X+1 && head.Y == tail.Y+2:
		tail.X++
		tail.Y++

	default:

	}

	return tail
}

func UpdatePos(pos map[Point]struct{}, direction string, steps int, head, tail Point) (map[Point]struct{}, Point, Point) {
	for i := 0; i < steps; i++ {

		head = MoveHead(direction, head)
		tail = CheckTail(head, tail)

		pos[Point{tail.X, tail.Y}] = struct{}{}
	}

	return pos, head, tail
}

func GetPos(input string) int {

	head, tail := Point{0, 0}, Point{0, 0}

	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	positions := make(map[Point]struct{})

	for _, line := range lines {

		if line == "" {
			break
		}

		instruction := strings.Split(line, " ")
		direction := instruction[0]
		steps, _ := strconv.Atoi(instruction[1])

		positions, head, tail = UpdatePos(positions, direction, steps, head, tail)
	}

	return len(positions)
}

func UpdatePosManyHeads(pos map[Point]struct{}, direction string, steps int, heads []Point) (map[Point]struct{}, []Point) {

	var lastTail Point

	for i := 0; i < steps; i++ {

		heads[0] = MoveHead(direction, heads[0])
		head := heads[0]

		for j := 1; j < len(heads); j++ {

			heads[j] = CheckTail(head, heads[j])
			head = heads[j]
		}

		lastTail = heads[len(heads)-1]

		pos[Point{lastTail.X, lastTail.Y}] = struct{}{}
	}

	return pos, heads
}

func GetPosManyHeads(input string) int {

	qtyHeads := 10
	heads := []Point{}

	for i := 0; i < qtyHeads; i++ {
		heads = append(heads, Point{0, 0})
	}

	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	positions := make(map[Point]struct{})

	for _, line := range lines {

		if line == "" {
			break
		}

		instruction := strings.Split(line, " ")
		direction := instruction[0]
		steps, _ := strconv.Atoi(instruction[1])

		positions, heads = UpdatePosManyHeads(positions, direction, steps, heads)
	}

	return len(positions)
}
