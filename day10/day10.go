package day10

import (
	"strconv"
	"strings"
)

type Register struct {
	Value int
	Cycle int
}

func (r *Register) Addx(x int) {
	r.Value += x
}

func (r *Register) RunCycle() {
	r.Cycle++
}

func CheckAndSum(reg *Register, valBefore, sumSignal, checkSignal int) (int, int) {
	stepCheckSignal := 40

	if reg.Cycle >= checkSignal {

		signal := valBefore * checkSignal

		sumSignal += signal

		checkSignal += stepCheckSignal
	}

	return sumSignal, checkSignal
}

func GetSumSignal(input string) int {

	reg := &Register{
		Value: 1,
	}

	input = strings.TrimSpace(input)

	lines := strings.Split(input, "\n")

	checkSignal := 20

	sumSignal := 0

	for _, line := range lines {

		instruction := strings.Split(line, " ")

		cmd := instruction[0]

		value := 0

		if len(instruction) > 1 {
			var err error
			value, err = strconv.Atoi(instruction[1])
			if err != nil {
				panic(err)
			}
		}

		valBefore := reg.Value

		switch cmd {
		case "addx":
			reg.RunCycle()
			reg.RunCycle()
			reg.Addx(value)
		case "noop":
			reg.RunCycle()
		}

		sumSignal, checkSignal = CheckAndSum(reg, valBefore, sumSignal, checkSignal)
	}

	return sumSignal
}

const SizeCRT = 40

type Sprite struct {
	Pos  int
	Draw string
}

func (s *Sprite) Move(input int) {
	s.Pos = input

	result := []rune{}
	for i := 0; i < SizeCRT; i++ {
		char := '.'
		if i == s.Pos-1 || i == s.Pos || i == s.Pos+1 {
			char = '#'
		}
		result = append(result, char)
	}

	s.Draw = string(result)
}

func PaintPixel(cycle int, spr *Sprite, result string) string {
	pos := cycle % SizeCRT
	if cycle != 1 && pos == 1 {
		result += "\n"
	}

	if pos == 0 && cycle > 0 {
		pos = 39
	} else if pos > 0 {
		pos = pos - 1
	}

	char := spr.Draw[pos]

	result += string(char)

	return result
}

func PaintCRT(input string) string {
	reg := &Register{
		Value: 1,
	}

	spr := &Sprite{
		Pos:  0,
		Draw: "###.....................................",
	}

	input = strings.TrimSpace(input)

	lines := strings.Split(input, "\n")

	result := ""

	for _, line := range lines {

		instruction := strings.Split(line, " ")

		cmd := instruction[0]

		value := 0

		if len(instruction) > 1 {
			var err error
			value, err = strconv.Atoi(instruction[1])
			if err != nil {
				panic(err)
			}
		}

		switch cmd {
		case "addx":
			reg.RunCycle()

			result = PaintPixel(reg.Cycle, spr, result)
			reg.RunCycle()

			result = PaintPixel(reg.Cycle, spr, result)
			reg.Addx(value)
			spr.Move(reg.Value)
		case "noop":
			reg.RunCycle()
			result = PaintPixel(reg.Cycle, spr, result)
		}
	}

	return result
}
