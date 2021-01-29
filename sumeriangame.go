package main

import (
	"bufio"
	"fmt"
	"os"
)

var name string
var population uint
var land uint
var inventory uint
var season string
var year uint

func main() {

	printIntroduction()
	fmt.Println("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')
	fmt.Println(name)

	printIstructions()

	setStartVars()

}

func printIntroduction() {
	fmt.Println("\n**********************************")
	fmt.Println("The Sumerian Game")
	fmt.Println("A long time ago... bla bla bla")
	fmt.Println("bla bla bla... ")
	fmt.Println("**********************************")
}

func printIstructions() {
	fmt.Println("To play use the following commands:")
	fmt.Println("bla bla bla")
	fmt.Println("bla bla bla... ")
}

func setStartVars() {
	population = 500
	land = 600
	inventory = 900
	season = "SPRING"
	year = 1
}
