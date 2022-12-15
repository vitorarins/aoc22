package day15

import (
	"fmt"
	"math"
	"strings"
)

type Point struct {
	X int
	Y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func (p Point) GetDistance(b Point) int {
	return abs(p.X-b.X) + abs(p.Y-b.Y)
}

type Sensor struct {
	Pos         Point
	BeaconPos   Point
	MaxDistance int
}

func ParseSensors(input string) []Sensor {
	lines := strings.Split(input, "\n")

	sensors := make([]Sensor, 0, len(lines))

	for _, line := range lines {

		var sy, sx, by, bx int

		_, _ = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)

		sensorPos := Point{Y: sy, X: sx}
		beaconPos := Point{Y: by, X: bx}

		maxDistance := sensorPos.GetDistance(beaconPos)

		sensors = append(sensors, Sensor{Pos: sensorPos, MaxDistance: maxDistance, BeaconPos: beaconPos})
	}

	return sensors
}

func CheckSensors(sensors []Sensor, checkFreq, startX, limitX, thePos int) int {
	total := 0

outer:
	for x := startX; x <= limitX; x++ {

		thisPos := Point{Y: thePos, X: x}

		for _, thisSensor := range sensors {

			dist := thisPos.GetDistance(thisSensor.Pos)

			if thisSensor.BeaconPos != thisPos && dist <= thisSensor.MaxDistance {
				jump := thisSensor.MaxDistance - abs(thisSensor.Pos.Y-thisPos.Y)
				jump += thisSensor.Pos.X - thisPos.X

				x += jump

				total += jump + 1

				continue outer
			}
		}

		if checkFreq > 0 {
			return thisPos.X*checkFreq + thisPos.Y
		}
	}

	if checkFreq > 0 {
		return 0
	}

	// need to subtract the last found position for it might be the distress beacon
	return total - 1
}

func GetNoBeaconsPos(input string, thePos int) int {
	input = strings.TrimSpace(input)

	sensors := ParseSensors(input)

	startX := math.MaxInt
	limitX := math.MinInt

	for _, sensor := range sensors {
		startX = int(math.Min(float64(sensor.Pos.X-sensor.MaxDistance-1), float64(startX)))
		limitX = int(math.Max(float64(sensor.Pos.X+sensor.MaxDistance+1), float64(limitX)))
	}

	return CheckSensors(sensors, 0, startX, limitX, thePos)
}

func GetNoBeacons(input string) int {

	thePos := 2000000

	return GetNoBeaconsPos(input, thePos)
}

func GetTuningFreqMax(input string, limitY int) int {
	input = strings.TrimSpace(input)

	sensors := ParseSensors(input)

	freqMult := 4000000

	for y := 0; y <= limitY; y++ {

		if freq := CheckSensors(sensors, freqMult, 0, maxCoord, y); freq > 0 {
			return freq
		}
	}

	return 0
}

func GetTuningFreq(input string) int {
	limitY := 4000000

	return GetTuningFreqMax(input, limitY)
}
