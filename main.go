package main

import (
	"log"
	"os"

	"github.com/vitorarins/aoc22/day1"
	"github.com/vitorarins/aoc22/day10"
	"github.com/vitorarins/aoc22/day11"
	"github.com/vitorarins/aoc22/day2"
	"github.com/vitorarins/aoc22/day3"
	"github.com/vitorarins/aoc22/day4"
	"github.com/vitorarins/aoc22/day5"
	"github.com/vitorarins/aoc22/day6"
	"github.com/vitorarins/aoc22/day7"
	"github.com/vitorarins/aoc22/day8"
	"github.com/vitorarins/aoc22/day9"
	"github.com/vitorarins/aoc22/util"
)

// day1 #part1
func main() {

	fsys := os.DirFS("./day1")
	inputContent, err := util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	fattiest, err := day1.GetFattiesElf(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 1, part 1: %d", fattiest)

	// day1 #part2

	fsys = os.DirFS("./day1")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	threeFattiest, err := day1.GetThreeFattiest(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 1, part 2: %d", threeFattiest)

	// day2 #part1

	fsys = os.DirFS("./day2")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	scoreSpec, err := day2.GetScoreSpec(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 2, part 1: %d", scoreSpec)

	// day2 #part2

	fsys = os.DirFS("./day2")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	score, err := day2.GetScore(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 2, part 2: %d", score)

	// day3 #part1

	fsys = os.DirFS("./day3")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	sumPrioItems, err := day3.GetSumPrioItems(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 3, part 1: %d", sumPrioItems)

	// day3 #part2

	fsys = os.DirFS("./day3")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	sumPrioBadges, err := day3.GetSumPrioBadges(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 3, part 2: %d", sumPrioBadges)

	// day4 #part1

	fsys = os.DirFS("./day4")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	amntBadAssigns, err := day4.GetBadAssignsAmnt(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 4, part 1: %d", amntBadAssigns)

	// day4 #part2

	fsys = os.DirFS("./day4")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	overlaps, err := day4.GetOverlap(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 4, part 2: %d", overlaps)

	// day5 #part1

	fsys = os.DirFS("./day5")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	stacksTopMoveOne, err := day5.GetStacksTopMoveOne(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 5, part 1: %s", stacksTopMoveOne)

	// day5 #part2

	fsys = os.DirFS("./day5")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	stacksTopMoveAll, err := day5.GetStacksTopMoveAll(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 5, part 2: %s", stacksTopMoveAll)

	// day6 #part1

	fsys = os.DirFS("./day6")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	firstMark, err := day6.GetFirstMark(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 6, part 1: %d", firstMark)

	// day6 #part2

	fsys = os.DirFS("./day6")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	startOfMsg, err := day6.GetStartOfMsg(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 6, part 2: %d", startOfMsg)

	// day7 #part1

	fsys = os.DirFS("./day7")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	sumUpTo, err := day7.GetFileSizes(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 7, part 1: %d", sumUpTo)

	// day7 #part2

	fsys = os.DirFS("./day7")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	sumMinDir, err := day7.GetMinDir(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Day 7, part 2: %d", sumMinDir)

	// day8 #part1

	fsys = os.DirFS("./day8")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	trees := day8.GetVisibleTrees(inputContent)

	log.Printf("Day 8, part 1: %d", trees)

	// day8 #part2

	fsys = os.DirFS("./day8")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	scenic := day8.GetScenic(inputContent)

	log.Printf("Day 8, part 2: %d", scenic)

	// day9 #part1

	fsys = os.DirFS("./day9")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	lenPos := day9.GetPos(inputContent)

	log.Printf("Day 9, part 1: %d", lenPos)

	// day9 #part2

	fsys = os.DirFS("./day9")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	lenPosMany := day9.GetPosManyHeads(inputContent)

	log.Printf("Day 9, part 2: %d", lenPosMany)

	// day10 #part1

	fsys = os.DirFS("./day10")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	sumSignal := day10.GetSumSignal(inputContent)

	log.Printf("Day 10, part 1: %d", sumSignal)

	// day10 #part2

	fsys = os.DirFS("./day10")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	crt := day10.PaintCRT(inputContent)

	log.Printf("Day 10, part 2: \n%s", crt)

	// day11 #part1

	fsys = os.DirFS("./day11")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	mBus := day11.GetMonkeyBusiness20(inputContent)

	log.Printf("Day 11, part 1: %d", mBus)

	// day11 #part2

	fsys = os.DirFS("./day11")
	inputContent, err = util.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	mBus10000 := day11.GetMonkeyBusiness10000(inputContent)

	log.Printf("Day 11, part 2: %d", mBus10000)

}
