package day11

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type WorryOp func(old *big.Int) *big.Int

type Item struct {
	Worry *big.Int
}

type Monkey struct {
	Id          int
	Items       []*Item
	Operation   WorryOp
	TestDiv     *big.Int
	ThrowTrue   int
	ThrowFalse  int
	Inspections int
	Worry       bool
}

func (m *Monkey) Add(it *Item) {
	m.Items = append(m.Items, it)
}

func (m *Monkey) Inspect(it *Item) *Item {

	it.Worry = m.Operation(it.Worry)
	if !m.Worry {
		three := big.NewInt(3)
		it.Worry = it.Worry.Div(it.Worry, three)
	}

	m.Inspections++

	return it
}

func (m *Monkey) Dance(monkeys []*Monkey, id int, lcm int64) []*Monkey {
	zero := big.NewInt(0)

	for i := 0; i < len(m.Items); i++ {
		item := m.Inspect(m.Items[i])

		// making item worry level smaller
		item.Worry = item.Worry.Mod(item.Worry, big.NewInt(lcm))

		modVal := new(big.Int)
		modVal = modVal.Mod(item.Worry, m.TestDiv)

		if modVal.Cmp(zero) == 0 {
			monkeys[m.ThrowTrue].Add(item)
		} else {
			monkeys[m.ThrowFalse].Add(item)
		}
	}

	m.Items = []*Item{}

	monkeys[id] = m

	return monkeys
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey:{ Id: %d, TestDiv: %v, Inspections: %d}", m.Id, m.TestDiv, m.Inspections)
}

func ParseItems(input string) []*Item {
	input = strings.TrimSpace(input)

	p := strings.Split(input, ":")

	it := p[1]
	it = strings.TrimSpace(it)

	itemList := strings.Split(it, ",")

	items := []*Item{}

	for _, iStr := range itemList {
		iStr = strings.TrimSpace(iStr)
		itemWorry, err := strconv.ParseInt(iStr, 10, 64)
		if err != nil {
			panic(err)
		}

		bigWorry := big.NewInt(itemWorry)

		item := &Item{
			Worry: bigWorry,
		}

		items = append(items, item)
	}

	return items
}

func ParseWorryOp(input string) WorryOp {
	input = strings.TrimSpace(input)

	p := strings.Split(input, ":")

	it := p[1]
	it = strings.TrimSpace(it)

	sides := strings.Split(it, "=")
	opSide := sides[1]
	opSide = strings.TrimSpace(opSide)

	elements := strings.Fields(opSide)

	if len(elements) != 3 {
		panic(fmt.Errorf("failed to parse worry op: %s", input))
	}

	oper := elements[1]
	s := elements[2]

	switch oper {
	case "*":
		if s == "old" {
			return func(old *big.Int) *big.Int {
				return old.Mul(old, old)
			}
		}

		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}

		nBig := big.NewInt(n)

		return func(old *big.Int) *big.Int {
			return old.Mul(old, nBig)
		}
	case "+":
		if s == "old" {
			return func(old *big.Int) *big.Int {
				return old.Add(old, old)
			}
		}

		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}

		nBig := big.NewInt(n)

		return func(old *big.Int) *big.Int {
			return old.Add(old, nBig)
		}
	}

	return nil
}

func ParseTestDiv(input string) *big.Int {
	input = strings.TrimSpace(input)

	p := strings.Split(input, ":")

	it := p[1]
	it = strings.TrimSpace(it)

	div := strings.Replace(it, "divisible by ", "", 1)
	testDiv, err := strconv.ParseInt(div, 10, 64)
	if err != nil {
		panic(err)
	}

	return big.NewInt(testDiv)
}

func ParseThrow(input string) int {
	input = strings.TrimSpace(input)

	p := strings.Split(input, ":")

	it := p[1]
	it = strings.TrimSpace(it)

	monkey := strings.Replace(it, "throw to monkey ", "", 1)
	throw, err := strconv.Atoi(monkey)
	if err != nil {
		panic(err)
	}

	return throw
}

func ParseMonkeyNote(id int, input string, worry bool) *Monkey {
	input = strings.TrimSpace(input)

	lines := strings.Split(input, "\n")

	if len(lines) != 6 {
		panic(fmt.Errorf("failed to parse monkey note: %s", input))
	}

	items := ParseItems(lines[1])
	worryOp := ParseWorryOp(lines[2])
	testDiv := ParseTestDiv(lines[3])
	throwTrue := ParseThrow(lines[4])
	throwFalse := ParseThrow(lines[5])

	return &Monkey{
		Id:         id,
		Items:      items,
		Operation:  worryOp,
		TestDiv:    testDiv,
		ThrowTrue:  throwTrue,
		ThrowFalse: throwFalse,
		Worry:      worry,
	}
}

func RunRound(round int, monkeys []*Monkey, lcm int64) []*Monkey {

	for i := 0; i < len(monkeys); i++ {
		m := monkeys[i]
		monkeys = m.Dance(monkeys, i, lcm)
	}

	return monkeys
}

func GetMonkeyBusiness(input string, rounds int, worry bool) int {
	input = strings.TrimSpace(input)

	monkeyNotes := strings.Split(input, "\n\n")

	monkeys := []*Monkey{}

	lcm := int64(1)

	for i, mn := range monkeyNotes {

		m := ParseMonkeyNote(i, mn, worry)

		lcm = m.TestDiv.Int64() * lcm

		monkeys = append(monkeys, m)
	}

	for i := 0; i < rounds; i++ {
		monkeys = RunRound(i+1, monkeys, lcm)
	}

	sort.Slice(monkeys, func(i, j int) bool {

		return monkeys[i].Inspections > monkeys[j].Inspections
	})

	monkey0 := monkeys[0]
	monkey1 := monkeys[1]

	return monkey0.Inspections * monkey1.Inspections
}

func GetMonkeyBusiness20(input string) int {
	return GetMonkeyBusiness(input, 20, false)
}
func GetMonkeyBusiness10000(input string) int {
	return GetMonkeyBusiness(input, 10000, true)
}
