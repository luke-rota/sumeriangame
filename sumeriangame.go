package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type state struct {
	gamerName         string
	population        uint
	land              uint
	inventory         uint
	season            string
	year              uint
	populationBushels uint
	sowingBushels     uint
}

var game state

func main() {

	userinputchan := make(chan string) //Create return channel for user input
	go getInput(userinputchan)         //Constantly wait for user input

	printIntroduction()
	fmt.Println("Enter your name: ")

	game.gamerName = <-userinputchan

	printIstructions()

	setStartVars(&game)

	outputSituation(&game)

	var confirmBushels bool = false

	for !confirmBushels {

		askBushels(&game, userinputchan)

		if (game.populationBushels + game.sowingBushels) > game.inventory {
			confirmBushels = false
		} else {
			fmt.Println("Confirm the bushels destination?")
			answer := <-userinputchan
			if answer == "Y" {
				confirmBushels = true
			} else {
				confirmBushels = false
			}
		}
	}

}

func getInput(returnchan chan<- string) {
	for {
		reader := bufio.NewReader(os.Stdin)    //Create reader
		input, _ := reader.ReadString('\n')    //Read from
		returnchan <- strings.TrimSpace(input) //Trim the input and send it
	}
}

func printIntroduction() {
	fmt.Println()
	fmt.Println("**********************************")
	fmt.Println("The Sumerian Game")
	fmt.Println("A long time ago... bla bla bla")
	fmt.Println("bla bla bla... ")
	fmt.Println("**********************************")
	fmt.Println()
}

func printIstructions() {
	fmt.Println()
	fmt.Println("To play use the following commands:")
	fmt.Println("bla bla bla")
	fmt.Println("bla bla bla... ")
	fmt.Println()
}

func setStartVars(game *state) {
	game.population = 500
	game.land = 600
	game.inventory = 900
	game.season = "SPRING"
	game.year = 1
	game.populationBushels = 0
	game.sowingBushels = 0
}

func outputSituation(game *state) {
	fmt.Println()
	fmt.Println("Population =", game.population)
	fmt.Println("Land =", game.land)
	fmt.Println("Inventory =", game.inventory)
	fmt.Println("Season =", game.season)
	fmt.Println("Year =", game.year)
	fmt.Println()
}

func askBushels(game *state, input chan string) {
	var strInput string
	fmt.Println()
	fmt.Println("How many bushels for the population?")
	strInput = <-input
	bushels, err := strconv.Atoi(strInput)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	game.populationBushels = uint(bushels)
	fmt.Println("How many bushels for the sowing?")
	strInput = <-input
	sowing, err := strconv.Atoi(strInput)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	game.sowingBushels = uint(sowing)
}
