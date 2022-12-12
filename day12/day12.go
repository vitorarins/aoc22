package day12

import (
	"fmt"
	"sort"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Elev rune
}

type Path struct {
	Pos      Point
	Steps    int
	Children []*Path
}

func (p *Path) String() string {
	return fmt.Sprintf("Path: %+v, Steps: %d", p.Pos, p.Steps)
}

type HeightMap map[Point]Square

func (hm HeightMap) GetPathsFrom(path Path) []*Path {
	p := path.Pos

	var neighbors []*Path

	curSqr, ok := hm[p]
	if !ok {
		return neighbors
	}

	for _, offset := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		neighbor := Point{X: p.X + offset.X, Y: p.Y + offset.Y}
		if sqr, ok := hm[neighbor]; ok {
			if sqr.Elev <= curSqr.Elev+1 {
				newPath := &Path{
					Pos:   neighbor,
					Steps: path.Steps + 1,
				}

				neighbors = append(neighbors, newPath)
			}
		}
	}

	return neighbors
}

func ParseHeightMap(input string) (HeightMap, Point, Point, []Point) {
	input = strings.TrimSpace(input)

	lines := strings.Split(input, "\n")

	hm := make(HeightMap)
	start, end := Point{}, Point{}
	possibleStarts := []Point{}

	for y, line := range lines {
		for x, c := range line {
			pos := Point{
				X: x,
				Y: y,
			}

			elev := c
			switch elev {
			case 'a':
				possibleStarts = append(possibleStarts, pos)
			case 'S':
				elev = 'a'
				start = pos
			case 'E':
				elev = 'z'
				end = pos
			}

			hm[pos] = Square{
				Elev: elev,
			}
		}
	}

	return hm, start, end, possibleStarts
}

func WalkMap(hm HeightMap, start, end Point) *Path {

	startPath := Path{
		Pos:   start,
		Steps: 0,
	}
	startChdrn := hm.GetPathsFrom(startPath)
	startPath.Children = startChdrn

	visitedPaths := map[Point]*Path{
		start: &startPath,
	}

	queueChdrn := make([]*Path, len(startChdrn))
	copy(queueChdrn, startChdrn)

	endPaths := []*Path{}

	for len(queueChdrn) > 0 {
		chd := queueChdrn[0]

		queueChdrn = queueChdrn[1:]

		if chd.Pos.X == end.X && chd.Pos.Y == end.Y {
			endPaths = append(endPaths, chd)
			continue
		}

		if _, ok := visitedPaths[chd.Pos]; ok {
			continue
		}

		chdrn := hm.GetPathsFrom(*chd)
		queueChdrn = append(queueChdrn, chdrn...)
		visitedPaths[chd.Pos] = chd
	}

	if len(endPaths) == 0 {
		return nil
	}

	sort.Slice(endPaths, func(i, j int) bool {
		return endPaths[i].Steps < endPaths[j].Steps
	})

	return endPaths[0]
}

func GetPathSteps(input string) int {

	hm, start, end, _ := ParseHeightMap(input)

	endPath := WalkMap(hm, start, end)

	if endPath != nil {
		return endPath.Steps
	}

	return 0
}

func GetPathStepsManyStarts(input string) int {

	quicker := GetPathSteps(input)

	hm, _, end, possibleStarts := ParseHeightMap(input)

	for _, start := range possibleStarts {

		endPath := WalkMap(hm, start, end)

		if endPath != nil {
			if endPath.Steps < quicker {
				quicker = endPath.Steps
			}
		}
	}

	return quicker
}
