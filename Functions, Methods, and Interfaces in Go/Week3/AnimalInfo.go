/*
@author: Tushar Nitave
Date: June 24, 2020
Title: Using struct in Go
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (info Animal) Eat() string{
	return info.food
}

func (info Animal) Move() string{
	return info.locomotion
}

func (info Animal) Speak() string{
	return info.noise
}


func main(){

	// Data structure to store animal information
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	// Infinite loop to take input from user
	for{
		fmt.Print("> ")
		userInput := bufio.NewReader(os.Stdin)
		input, _, _ := userInput.ReadLine()
		data := strings.Split(string(input), " ")

		// maps user input to respective animal
		selector1:= map[string]Animal{
			"cow":cow,
			"bird": bird,
			"snake": snake,
		}

		// calls desired function
		if data[1] == "eat"{
			fmt.Println(selector1[data[0]].Eat())
		}else if data[1] == "speak"{
			fmt.Println(selector1[data[0]].Speak())
		} else{
			fmt.Println(selector1[data[0]].Move())
		}

	}
}
