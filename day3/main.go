package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Backpack struct {
	Left  []rune
	Right []rune
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var elf []Backpack
	//var counter int
	for scanner.Scan() {
		text := scanner.Text()
		length := len(text)
		var newBackpack Backpack
		for i, t := range []rune(text) {
			if i >= length/2 {
				newBackpack.Right = append(newBackpack.Right, t)
			} else {
				newBackpack.Left = append(newBackpack.Left, t)
			}
		}
		elf = append(elf, newBackpack)
	}
	day1(elf)
	day2(elf)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day1(elf []Backpack) {
	var recheck int
	var found bool = false
	for _, e := range elf {
		for _, t := range e.Right {
			for _, c := range e.Left {
				if c == t {
					fmt.Printf("%c %c\n", c, t)
					if unicode.IsUpper(t) {
						amount := t - 38
						fmt.Println(int(amount))
						recheck = recheck + int(amount)
					} else {
						amount := t - 96
						fmt.Println(int(amount))
						recheck = recheck + int(amount)
					}
					found = true
					break
				}
			}
			if found {
				found = false
				fmt.Println(recheck)
				break
			}
		}
	}
	fmt.Println(recheck)
}

func day2(elf []Backpack) {
	var recheck int
	var found bool = false
	for i := 0; i < len(elf); {
		for _, t := range append(elf[i].Left, elf[i].Right...) {
			for _, u := range append(elf[i+1].Left, elf[i+1].Right...) {
				if t == u {
					//Match on 2, check 3
					for _, v := range append(elf[i+2].Left, elf[i+2].Right...) {
						if t == v {
							found = true
							//Match one 3
							if unicode.IsUpper(t) {
								amount := t - 38
								fmt.Println(int(amount))
								recheck = recheck + int(amount)
							} else {
								amount := t - 96
								fmt.Println(int(amount))
								recheck = recheck + int(amount)
							}
							break
						}
					}
				}
				if found == true {
					break
				}
			}
			if found == true {
				found = false
				break
			}
		}
		i = i + 3
	}
	fmt.Println(recheck)
}
