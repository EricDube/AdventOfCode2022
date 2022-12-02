package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var maxelf1, maxelf2, maxelf3 int

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var newelf int
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			fmt.Printf("%d %d %d %d\n", newelf, maxelf1, maxelf2, maxelf3)
			//empty new elf
			var changed int
			if newelf > maxelf1 {
				changed = maxelf1
				maxelf1 = newelf
				if changed > maxelf2 {
					changed2 := changed
					maxelf2 = changed
					if changed2 > maxelf3 {
						maxelf3 = changed
					}
				} else if changed > maxelf3 {
					maxelf3 = changed
				}
			} else if newelf > maxelf2 {
				changed = maxelf2
				maxelf2 = newelf
				if changed > maxelf3 {
					maxelf3 = changed
				}
			} else if newelf > maxelf3 {
				changed = maxelf3
				maxelf3 = newelf
			}

			newelf = 0
		} else {
			add, _ := strconv.Atoi(text)
			newelf = newelf + add
		}
	}

	fmt.Println(maxelf1 + maxelf2 + maxelf3)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
