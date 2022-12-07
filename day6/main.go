package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println(text)
	}
	day1(text)
	day2(text)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day1(text string) {
	var runeText []rune = []rune(text)
	var t1, t2, t3, t4 rune
	for i, r := range runeText {
		t1 = t2
		t2 = t3
		t3 = t4
		t4 = r
		if t1 == 0 {
			//The buffer isnt full so just continue
			continue
		}
		if t1 == t2 || t1 == t3 || t1 == t3 || t1 == t4 {
			//Theres a match, continue
			continue
		}
		if t2 == t3 || t2 == t4 {
			//Theres a match, continue
			continue
		}
		if t3 == t4 {
			//Theres a match, continue
			continue
		}
		//i is the answer
		fmt.Printf("The answer: %v%v%v%v %d", string(t1), string(t2), string(t3), string(t4), i+1)
		break
	}
}

func hasMatch(text []rune) bool {
	for i := 0; i < len(text); i++ {
		for j := i + 1; j < len(text); j++ {
			if text[j] == text[i] {
				//Theres a match, continue
				return true
			}
		}
	}
	return false
}

func day2(text string) {
	var runeText []rune = []rune(text)
	for i, _ := range runeText {
		if i+14 > len(runeText) {
			return
		}
		if hasMatch(runeText[i : i+14]) {
			continue
		}
		fmt.Printf("Found answer: %d\n", i+14)
		fmt.Printf("%v\n", string(runeText[i:i+14]))
		return
	}
}
