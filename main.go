package main

import (
	"log"
	"os"

	"github.com/vitorarins/aoc22/day2"
)

// day1 #part1
// func main() {

// 	fsys := os.DirFS("./day1")
// 	inputContent, err := day1.ReadInput(fsys, "input")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fattiest, err := day1.GetFattiesElf(inputContent)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print(fattiest)

// }

// day1 #part2
// func main() {

// 	fsys := os.DirFS("./day1")
// 	inputContent, err := day1.ReadInput(fsys, "input")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fattiest, err := day1.GetThreeFattiest(inputContent)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print(fattiest)

// }

// day2 #part1
// func main() {

// 	fsys := os.DirFS("./day2")
// 	inputContent, err := day2.ReadInput(fsys, "input")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	score, err := day2.GetScoreSpec(inputContent)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print(score)

// }

// day2 #part2
func main() {

	fsys := os.DirFS("./day2")
	inputContent, err := day2.ReadInput(fsys, "input")
	if err != nil {
		log.Fatal(err)
	}

	score, err := day2.GetScore(inputContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(score)

}
