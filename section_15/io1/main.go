package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type team []string

var FileName = "names.txt"
var newFileName = "new.txt"

func main() {

	// Uncomment one at a time.
	// WriteToFile()
	// ReadFromFile()
	// RenameFile()
	RemoveFile()

}

// ===============================================================
func WriteToFile() {

	players := team{"Messi", "Pele", "Maradona", "Zidane", "Platini"}

	err := ioutil.WriteFile(FileName, []byte(players.toString()), 0666)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Wrote the following to file: %v \n", players.toString())
	}

}

func (t team) toString() string {

	return strings.Join([]string(t), ",")

}

// ===============================================================
func ReadFromFile() {

	bs, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	s := strings.Split(string(bs), ",")
	fmt.Println(s)

}

// ===============================================================
func RenameFile() {

	err := os.Rename(FileName, newFileName)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("File renamed.")
	}

}

// ===============================================================
func RemoveFile() {

	err := os.Remove(newFileName)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("File removed.")
	}

}
