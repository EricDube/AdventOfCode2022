package main

import (
	"adventcode/Linked"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//var grid [][]rune

type Grid struct {
	row Linked.LinkedList
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var gridreader []string
	var moves []string
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "[") {
			//Its a grid, enter into grid
			gridreader = append(gridreader, text)
		} else if strings.Contains(text, "move") {
			moves = append(moves, text)
		}
	}
	var grid [10]Grid
	var counter int
	var lastrune rune
	for _, g := range gridreader {
		counter = 1
		for j, r := range g {
			if lastrune == []rune("[")[0] {
				//Next spot is a character
				//grid[len(gridreader)-i-1][(j+2)/4] = r
				grid[((j+3)/4)-1].row.InsertAtHead(string(r))
			}

			lastrune = r
			counter++
		}
	}
	for _, g := range grid {
		g.row.Display()
	}

	//	day1(grid, moves)
	day2(grid, moves)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day1(grid [10]Grid, moves []string) {
	for _, move := range moves {
		cleaned := strings.SplitAfter(move, " ")
		//fmt.Println(cleaned)
		amountToMove, _ := strconv.Atoi(strings.TrimSpace(cleaned[1]))
		from, _ := strconv.Atoi(strings.TrimSpace(cleaned[3]))
		to, _ := strconv.Atoi(strings.TrimSpace(cleaned[5]))
		for i := 0; i < amountToMove; i++ {
			grid[to-1].row.InsertAtTail(grid[from-1].row.GetTail().GetData())
			grid[from-1].row.DeleteAtTail()
			grid[from-1].row.Display()
			grid[to-1].row.Display()
			fmt.Println()
		}
	}

	for _, g := range grid {
		g.row.Display()
	}
}

func day2(grid [10]Grid, moves []string) {
	for _, move := range moves {
		cleaned := strings.SplitAfter(move, " ")
		//fmt.Println(cleaned)
		amountToMove, _ := strconv.Atoi(strings.TrimSpace(cleaned[1]))
		from, _ := strconv.Atoi(strings.TrimSpace(cleaned[3]))
		to, _ := strconv.Atoi(strings.TrimSpace(cleaned[5]))
		var temp Linked.LinkedList
		for i := 0; i < amountToMove; i++ {
			temp.InsertAtHead(grid[from-1].row.GetTail().GetData())
			grid[from-1].row.DeleteAtTail()
		}
		grid[to-1].row.Append(&temp)
		grid[to-1].row.Display()
	}
	fmt.Println()
	for _, g := range grid {
		g.row.Display()
	}
}
