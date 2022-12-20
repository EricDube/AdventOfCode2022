package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var text string
	var counter int
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println(text)
		grid = append(grid, []int{})
		for _, c := range text {
			n, _ := strconv.Atoi(string(c))
			grid[counter] = append(grid[counter], n)
		}
		counter++
	}
	max = len(grid[counter-1])
	//day1()
	day2()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var grid [][]int
var max int
var countTotalVisible int

func checkLeft(x, y int) bool {
	n := grid[x][y]
	for left := y - 1; left >= 0; left-- {
		if grid[x][left] >= n {
			return false
		}
	}
	return true
}
func checkRight(x, y int) bool {
	n := grid[x][y]
	for right := y + 1; right < max; right++ {
		if grid[x][right] >= n {
			return false
		}
	}
	return true
}
func checkUp(x, y int) bool {
	n := grid[x][y]
	for up := x - 1; up >= 0; up-- {
		if grid[up][y] >= n {
			return false
		}
	}
	return true
}
func checkDown(x, y int) bool {
	n := grid[x][y]
	for down := x + 1; down < max; down++ {
		if grid[down][y] >= n {
			return false
		}
	}
	return true
}

func day1() {
	countTotalVisible = len(grid)*2 + len(grid[0])*2 - 4
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if checkLeft(i, j) {
				countTotalVisible++
				continue
			}
			if checkRight(i, j) {
				countTotalVisible++
				continue
			}
			if checkUp(i, j) {
				countTotalVisible++
				continue
			}
			if checkDown(i, j) {
				countTotalVisible++
				continue
			}

		}
	}
	fmt.Printf("\nFinal visible count: %d", countTotalVisible)
}

func countLeft(x, y int) int {
	n := grid[x][y]
	var count int
	for left := y - 1; left >= 0; left-- {
		count++
		if grid[x][left] >= n {
			return count
		}
	}
	return count
}
func countRight(x, y int) int {
	n := grid[x][y]
	var count int
	for right := y + 1; right < max; right++ {
		count++
		if grid[x][right] >= n {
			return count
		}
	}
	return count
}
func countUp(x, y int) int {
	n := grid[x][y]
	var count int
	for up := x - 1; up >= 0; up-- {
		count++
		if grid[up][y] >= n {
			return count
		}
	}
	return count
}
func countDown(x, y int) int {
	n := grid[x][y]
	var count int
	for down := x + 1; down < max; down++ {
		count++
		if grid[down][y] >= n {
			return count
		}
	}
	return count
}

var maxScenic int

func day2() {
	countTotalVisible = len(grid)*2 + len(grid[0])*2 - 4
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			scenic := countRight(i, j) * countLeft(i, j) * countUp(i, j) * countDown(i, j)
			if scenic > maxScenic {
				maxScenic = scenic
			}
		}
	}
	fmt.Printf("\nFinal visible count: %d", maxScenic)
}
