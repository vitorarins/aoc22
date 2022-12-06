package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.

		return element, true
	}
}

func reverse(input []string) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func cleanup(input []string) []string {
	result := make([]string, 0)

	for _, s := range input {
		if strings.TrimSpace(s) != "" {
			result = append(result, s)
		}
	}

	return result
}

func ParseStackLine(stacks map[string]Stack, numbers []string, line string) map[string]Stack {

	boxLen := 3
	lineCounter := 0

	for i, stackNumber := range numbers {

		stack, ok := stacks[stackNumber]
		if !ok {
			stack = Stack{}
		}

		start := (i * boxLen) + lineCounter

		box := line[start : start+boxLen]
		box = strings.TrimSpace(box)

		if len(box) == boxLen {
			stack.Push(string(box[1]))
		}

		stacks[stackNumber] = stack

		lineCounter++
	}

	return stacks
}

func ParseStacks(input string) (map[string]Stack, error) {

	lines := strings.Split(input, "\n")
	lines = cleanup(lines)
	reverse(lines)

	stackNumbers := strings.Split(lines[0], " ")
	stackNumbers = cleanup(stackNumbers)

	lsLabel := stackNumbers[len(stackNumbers)-1]

	lastStackNumber, err := strconv.Atoi(lsLabel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse last stack number: %w", err)
	}

	result := make(map[string]Stack, lastStackNumber)

	for i, line := range lines {

		if i == 0 {
			continue
		}

		result = ParseStackLine(result, stackNumbers, line)
	}

	return result, nil
}

type MoveFunction func(map[string]Stack, int, string, string) (map[string]Stack, error)

func MoveOne(stacks map[string]Stack, move int, from, to string) (map[string]Stack, error) {

	stackFrom, ok := stacks[from]
	if !ok {
		return stacks, fmt.Errorf("could not find 'from' key in stacks: %s", from)
	}

	stackTo, ok := stacks[to]
	if !ok {
		return stacks, fmt.Errorf("could not find 'to' key in stacks: %s", to)
	}

	for i := 0; i < move; i++ {

		topBox, ok := stackFrom.Pop()
		if !ok {
			return stacks, fmt.Errorf("stack at 'from' key %q is empty: %+v", from, stackFrom)
		}

		stackTo.Push(topBox)

	}

	stacks[from] = stackFrom
	stacks[to] = stackTo

	return stacks, nil
}

func MoveAll(stacks map[string]Stack, move int, from, to string) (map[string]Stack, error) {

	mem := []string{}

	stackFrom, ok := stacks[from]
	if !ok {
		return stacks, fmt.Errorf("could not find 'from' key in stacks: %s", from)
	}

	stackTo, ok := stacks[to]
	if !ok {
		return stacks, fmt.Errorf("could not find 'to' key in stacks: %s", to)
	}

	for i := 0; i < move; i++ {

		topBox, ok := stackFrom.Pop()
		if !ok {
			return stacks, fmt.Errorf("stack at 'from' key %q is empty: %+v", from, stackFrom)
		}

		stacks[from] = stackFrom

		mem = append(mem, topBox)
	}

	reverse(mem)

	for _, box := range mem {

		stackTo.Push(box)
	}

	stacks[from] = stackFrom
	stacks[to] = stackTo

	return stacks, nil
}

func ParseInstructLine(stacks map[string]Stack, input string, moveFn MoveFunction) (map[string]Stack, error) {

	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		return stacks, fmt.Errorf("wrong number of matches in line: %s", input)
	}

	amntMv := matches[1]
	from := matches[2]
	to := matches[3]

	move, err := strconv.Atoi(amntMv)
	if err != nil {
		return stacks, fmt.Errorf("failed to parse amount to move: %w", err)
	}

	return moveFn(stacks, move, from, to)
}

func ParseInstructs(stacks map[string]Stack, input string, move MoveFunction) (string, error) {
	lines := strings.Split(input, "\n")
	lines = cleanup(lines)

	for _, l := range lines {
		line := strings.TrimSpace(l)

		var err error
		stacks, err = ParseInstructLine(stacks, line, move)
		if err != nil {
			return "", fmt.Errorf("failed to parse instruct line: %w", err)
		}
	}

	result := ""

	for i := 0; i < len(stacks); i++ {
		key := strconv.Itoa(i + 1)
		stack, ok := stacks[key]
		if !ok {
			return "", fmt.Errorf("could not find key in stacks: %s", key)
		}

		topBox, ok := stack.Pop()
		if !ok {
			return "", fmt.Errorf("stack at key %q is empty: %+v", key, stack)
		}

		result += topBox
	}

	return result, nil
}

func getStacksTop(input string, move MoveFunction) (string, error) {
	inputs := strings.Split(input, "\n\n")

	if len(inputs) != 2 {
		return "", fmt.Errorf("unable to split stack from instructions:\n%s", input)
	}

	stackInput := inputs[0]
	instructInput := inputs[1]

	stacks, err := ParseStacks(stackInput)
	if err != nil {
		return "", fmt.Errorf("failed to parse stacks: %w", err)
	}

	answer, err := ParseInstructs(stacks, instructInput, move)
	if err != nil {
		return "", fmt.Errorf("failed to parse instructions: %w", err)
	}

	return answer, nil
}

func GetStacksTopMoveOne(input string) (string, error) {

	return getStacksTop(input, MoveOne)
}

func GetStacksTopMoveAll(input string) (string, error) {

	return getStacksTop(input, MoveAll)
}
