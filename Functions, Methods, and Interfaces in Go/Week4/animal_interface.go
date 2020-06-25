/*
@author: Tushar Nitave
Date: June 25, 2020
Title: Interfaces in Go
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animals struct{
	eat, speak, move string
}

type animalInterface interface {
	Eat()
	Speak()
	Move()
}

func (a Animals) Eat(){
	fmt.Println(a.eat)
}

func (a Animals) Speak(){
	fmt.Println(a.speak)
}

func (a Animals) Move(){
	fmt.Println(a.move)
}

func main(){

	cow := Animals{"grass", "moo", "walk"}
	bird := Animals{"worms", "peep", "fly"}
	snake := Animals{"mice", "hsss", "slither"}

	var animalInfo animalInterface

	animalSel := map[string]Animals{
		"cow":cow,
		"bird":bird,
		"snake":snake,
	}

	for{
		fmt.Println("Enter your request (newanimal/query/exit)")
		fmt.Print("> ")

		userInput := bufio.NewReader(os.Stdin)
		input, _, _ := userInput.ReadLine()
		request := strings.Split(string(input), " ")

		if request[0] == "query"{
			animalInfo = animalSel[request[1]]
			if request[2] == "eat" {
				animalInfo.Eat()
			}else if request[2] == "speak"{
				animalInfo.Speak()
			}else if request[2] == "move"{
				animalInfo.Move()
			}
		}else if request[0] == "newanimal"{
			animalSel[request[1]] = animalSel[request[2]]
			fmt.Println("Created it!")
		}else if request[0] == "exit" {
			fmt.Println("Program exited successfully!")
			os.Exit(3)
		}
	}

}



