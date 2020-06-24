/*
@author: Tushar Nitave
Date: June 24, 2020
Title: Calculate displacement using "returning a function concept"
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(acc, init_v, init_dis float64) func(float64) float64{
	compute := func(time float64) float64{
		final_disp := (acc*math.Pow(time, 2)*0.5) + (init_v*time) + init_dis
		return final_disp
	}
	return compute
}

func main(){
	fmt.Println("\nEnter values of acceleration, velocity and displacement: ")
	userInput := bufio.NewReader(os.Stdin)
	values, _, _ := userInput.ReadLine()
	inputs := strings.Split(string(values), " ")

	acceleration, _ := strconv.ParseFloat(inputs[0], 64)
	initial_velocity, _ := strconv.ParseFloat(inputs[1], 64)
	initial_disp, _ := strconv.ParseFloat(inputs[2], 64)

	fmt.Println("\nEnter the time: ")
	userInput = bufio.NewReader(os.Stdin)
	value, _, _ := userInput.ReadLine()
	input := strings.Split(string(value), " ")

	time, _ := strconv.ParseFloat(input[0], 64)

	fn := GenDisplaceFn(acceleration, initial_velocity, initial_disp)
	fmt.Println("Final displacement = ", fn(time))

}
