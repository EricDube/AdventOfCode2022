package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	size int
	name string
}

type Folder struct {
	size    int
	name    string
	files   []*File
	folders []*Folder
	parent  *Folder
}

var currentFolder *Folder
var rootFolder Folder

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var text string

	rootFolder = Folder{0, "/", []*File{}, []*Folder{}, nil}
	currentFolder = &rootFolder
	commands := make(chan string, 250)
	var newCommand bool
	for scanner.Scan() {
		text = scanner.Text()
		switch strings.HasPrefix(text, "$") {
		case true:
			if newCommand == false {
				//Start new command
				newCommand = true
				commands <- text
			} else {
				//Run the last command
				close(commands)
				performCommands(commands)
				//Start a new queue
				newCommand = true
				commands = make(chan string, 250)
				commands <- text
			}
		case false:
			commands <- text
		}
	}
	close(commands)
	performCommands(commands)

	minimumSizeToDelete = 30000000 - (70000000 - rootFolder.size)

	//var tabcount int
	fmt.Println("/")
	for _, f := range rootFolder.files {
		fmt.Printf(" %d %s\n", f.size, f.name)
	}

	for _, f1 := range rootFolder.folders {
		fmt.Printf(" %d %s\n", f1.size, f1.name)
		display(f1, 1)
	}
	fmt.Printf("\n\nTotal folders with max of 100000: %d", sizeCount)

	fmt.Printf("\nTotal size: %d", rootFolder.size)
	fmt.Printf("\nLooking to delete folder of size: %d", minimumSizeToDelete)
	fmt.Printf("\nSmallest folder to delete: %+v", folderToDelete)
	day1(text)
	day2(text)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var minimumSizeToDelete int
var folderToDelete *Folder
var sizeCount int

func display(folder *Folder, space int) {
	//print each file in folder
	for _, f := range folder.files {
		for i := 0; i < space; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf(" %d %s\n", f.size, f.name)
	}
	//print each folder and iterate over each subfolder
	for _, f1 := range folder.folders {
		for i := 0; i < space; i++ {
			fmt.Printf(" ")
		}
		if f1.size <= 100000 {
			sizeCount += f1.size
		}
		if f1.size > minimumSizeToDelete {
			if folderToDelete == nil {
				folderToDelete = f1
			}
			if f1.size < folderToDelete.size {
				folderToDelete = f1
			}
		}
		fmt.Printf(" %d %s\n", f1.size, f1.name)
		display(f1, space+2)
	}
}

func performCommands(commands chan string) {
	var startList bool
	for command := range commands {
		println(command)
		if strings.HasPrefix(command, "$") {
			parse := strings.Split(command, " ")
			switch parse[1] {
			case "cd":
				switch parse[2] {
				case "/": //Go back to root
					currentFolder = &rootFolder
				case "..": //Go up a folder
					if currentFolder.parent == nil {
						fmt.Println("Cant go up more. already at root")
					} else {
						currentFolder = currentFolder.parent
					}
				default:
					//Find folder in current folder
					for _, folder := range currentFolder.folders {
						if folder.name == parse[2] {
							currentFolder = folder
							break
						}
					}
				}

			case "ls":
				startList = true
			}
		} else if startList == true {
			parse := strings.Split(command, " ")
			switch parse[0] {
			case "dir":
				newFolder := Folder{0, parse[1], []*File{}, []*Folder{}, currentFolder}
				currentFolder.folders = append(currentFolder.folders, &newFolder)
			default:
				size, _ := strconv.Atoi(parse[0])
				newFile := File{size, parse[1]}
				currentFolder.files = append(currentFolder.files, &newFile)
				currentFolder.size += size
				tempFolder := currentFolder.parent
				for tempFolder != nil {
					tempFolder.size += size
					tempFolder = tempFolder.parent
				}
			}
		}
	}
	/*if strings.HasPrefix("text", "$") {
		parse := strings.Split(text, " ")
		switch parse[1] {
		case "cd":
			switch parse[2] {
			case "/": //Go back to root
				currentFolder = &rootFolder
			case "..": //Go up a folder
				if currentFolder.parent == nil {
					fmt.Println("Cant go up more. already at root")
				} else {
					currentFolder = currentFolder.parent
				}
			default:
				//Find folder in current folder
				for _, folder := range currentFolder.folders {
					if folder.name == parse[2] {
						currentFolder = folder
						break
					}
				}
			}

		case "ls":
		}
	}*/
}

func day1(text string) {

}

func day2(text string) {

}
