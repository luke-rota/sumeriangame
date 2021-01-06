package main

import (
	"bufio"
	"fmt"
	"os"
)

var name string

func main() {

	printIntroduction()
	fmt.Println("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println(name)

}

func printIntroduction() {
	fmt.Println("\n**********************************")
	fmt.Println("The Sumerian Game")
	fmt.Println("A long time ago... bla bla bla")
	fmt.Println("bla bla bla... ")
	fmt.Println("**********************************")
}
