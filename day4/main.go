package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var counter int
	for scanner.Scan() {
		text := scanner.Text()
		values := strings.Split(text, ",")
		one := strings.Split(values[0], "-")
		two := strings.Split(values[1], "-")
		fmt.Printf("%s %s\n", one, two)
		leftone, _ := strconv.Atoi(one[0])
		rightone, _ := strconv.Atoi(one[1])
		lefttwo, _ := strconv.Atoi(two[0])
		righttwo, _ := strconv.Atoi(two[1])

		if leftone <= lefttwo && rightone >= lefttwo {
			counter++
		} else if leftone <= righttwo && rightone >= righttwo {
			counter++
		} else if lefttwo <= leftone && righttwo >= leftone {
			counter++
		} else if lefttwo <= rightone && righttwo >= rightone {
			counter++
		}

	}
	fmt.Println(counter)
	day1()
	day2()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day1() {

}

func day2() {

}
