package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Hand struct {
	Opponent string
	Choice   string
	Strategy string
	Score    int
}

func main() {
	part1()
	part2()
}

func part2() {
	var Game []Hand
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var total int
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")
		h := Hand{s[0], "", s[1], 0}
		switch h.Strategy {
		case "X": //Lose
			switch h.Opponent {
			case "A":
				h.Choice = "C"
				h.Score = 3
			case "B":
				h.Choice = "A"
				h.Score = 1
			case "C":
				h.Choice = "B"
				h.Score = 2
			}
			h.Score = h.Score + 0 //Loss
		case "Y": //Draw
			switch h.Opponent {
			case "A":
				h.Choice = "A"
				h.Score = 1
			case "B":
				h.Choice = "B"
				h.Score = 2
			case "C":
				h.Choice = "C"
				h.Score = 3
			}
			h.Score = h.Score + 3 //Draw
		case "Z": //Win
			switch h.Opponent {
			case "A":
				h.Choice = "B"
				h.Score = 2
			case "B":
				h.Choice = "C"
				h.Score = 3
			case "C":
				h.Choice = "A"
				h.Score = 1
			}
			h.Score = h.Score + 6 //Win
		}
		//fmt.Printf("%d %d\n", total, h.Score)
		total = total + h.Score
		Game = append(Game, h)
	}
	var t int
	for _, g := range Game {
		t = t + g.Score
	}
	fmt.Println(total)
	fmt.Println(t)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	var Game []Hand
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var total int
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")
		h := Hand{s[0], "", s[1], 0}
		switch h.Strategy {
		case "X":
			h.Choice = "A"
			h.Score = 1
		case "Y":
			h.Choice = "B"
			h.Score = 2
		case "Z":
			h.Choice = "C"
			h.Score = 3
		}
		switch {
		case h.Opponent == "A" && h.Choice == "C":
			h.Score = h.Score + 0
		case h.Opponent == "C" && h.Choice == "A":
			h.Score = h.Score + 6
		case h.Opponent > h.Choice:
			h.Score = h.Score + 0
		case h.Opponent < h.Choice:
			h.Score = h.Score + 6
		case h.Opponent == h.Choice:
			h.Score = h.Score + 3
		}
		//fmt.Printf("%d %d\n", total, h.Score)
		total = total + h.Score
		Game = append(Game, h)
	}
	var t int
	for _, g := range Game {
		t = t + g.Score
	}
	fmt.Println(total)
	fmt.Println(t)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
