package day14

import (
	"strconv"
	"strings"
)

const infinity = 1000

type Point struct {
	X int
	Y int
}

func (p Point) Compare(b Point) (string, int) {
	switch {
	case p.X > b.X:
		return "X", p.X - b.X
	case p.X < b.X:
		return "-X", b.X - p.X
	case p.Y > b.Y:
		return "Y", p.Y - b.Y
	case p.Y < b.Y:
		return "-Y", b.Y - p.Y
	default:
		return "", 0
	}
}

type Cave struct {
	Rocks  map[Point]string
	StartX int
	LimitX int
	LimitY int
}

func (c *Cave) String() string {
	result := ""

	for j := 0; j <= c.LimitY; j++ {
		for i := c.StartX; i <= c.LimitX; i++ {
			if p, ok := c.Rocks[Point{i, j}]; ok {
				result += p
			} else {
				result += "."
			}
		}
		result += "\n"
	}

	return result
}

func (c *Cave) UpdateStartAndLimit(rock Point) {
	if rock.X < c.StartX {
		c.StartX = rock.X
	}

	if rock.X > c.LimitX {
		c.LimitX = rock.X
	}

	if rock.Y > c.LimitY {
		c.LimitY = rock.Y
	}
}

func (c *Cave) DrawLine(direction string, iterations int, curRock Point) {

	c.UpdateStartAndLimit(curRock)

	curPos := curRock

	c.Rocks[curPos] = "#"

	for i := 0; i < iterations; i++ {
		switch direction {
		case "X":
			curPos.X++
		case "Y":
			curPos.Y++
		case "-X":
			curPos.X--
		case "-Y":
			curPos.Y--
		}

		c.Rocks[curPos] = "#"
	}

	c.UpdateStartAndLimit(curPos)
}

func (c *Cave) DrawBottom() {

	c.StartX -= infinity
	c.LimitX += infinity

	c.LimitY = c.LimitY + 2

	for i := c.StartX; i <= c.LimitX; i++ {
		c.Rocks[Point{i, c.LimitY}] = "#"
	}
}

func ParseRock(rock string) Point {
	rockPos := strings.Split(rock, ",")

	rockX, _ := strconv.Atoi(rockPos[0])
	rockY, _ := strconv.Atoi(rockPos[1])

	return Point{
		X: rockX,
		Y: rockY,
	}
}

func ParseCave(input string, withBottom bool) Cave {
	cave := &Cave{
		Rocks: make(map[Point]string),
	}

	rockPaths := strings.Split(input, "\n")

	for _, rockPath := range rockPaths {

		rocks := strings.Split(rockPath, " -> ")

		curRock := ParseRock(rocks[0])

		if cave.StartX == 0 {
			cave.StartX = curRock.X
		}

		for _, rock := range rocks {

			rockPos := ParseRock(rock)

			dir, iter := rockPos.Compare(curRock)
			if dir == "" {
				continue
			}

			cave.DrawLine(dir, iter, curRock)

			curRock = rockPos
		}
	}

	if withBottom {
		cave.DrawBottom()
	}

	return *cave
}

func PourGrain(cave *Cave, startSand Point) (*Cave, bool) {
	curSand := startSand
	for i := startSand.Y; i <= cave.LimitY; i++ {

		curSand = Point{
			X: curSand.X,
			Y: i,
		}

		if _, ok := cave.Rocks[curSand]; !ok {
			continue
		}

		if _, cur := curSand.Compare(startSand); cur == 0 {
			return cave, true
		}

		curSand.X--

		if _, ok := cave.Rocks[curSand]; !ok {
			continue
		}

		curSand.X += 2

		if _, ok := cave.Rocks[curSand]; !ok {
			continue
		}

		curSand.X--
		curSand.Y--

		cave.Rocks[curSand] = "o"

		return cave, false
	}

	return cave, true
}

func PourSand(cave *Cave, startSand Point) int {

	units := 0
	reachedLimit := false

	for !reachedLimit {

		cave, reachedLimit = PourGrain(cave, startSand)
		units++
	}

	return units - 1
}

func GetSandUnits(input string, withBottom bool) int {
	input = strings.TrimSpace(input)

	cave := ParseCave(input, withBottom)

	startSand := Point{
		X: 500,
		Y: 0,
	}

	sandUnits := PourSand(&cave, startSand)

	return sandUnits
}

func GetSandUnitsBottomless(input string) int {
	return GetSandUnits(input, false)
}

func GetSandUnitsWithBottom(input string) int {
	return GetSandUnits(input, true)
}
