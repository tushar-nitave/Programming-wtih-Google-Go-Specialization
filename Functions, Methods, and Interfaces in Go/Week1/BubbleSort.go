package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(seq []int, j int){
	var temp int
	temp = seq[j]
	seq[j] = seq[j+1]
	seq[j+1] = temp
}

func BubbleSort(seq []int){
	for i:=0 ; i < len(seq) ; i++{
		for j := 0; j<len(seq)-1-i; j++{
			if seq[j+1] < seq[j]{
				Swap(seq, j)
			}
		}
	}
}

func main(){
	// getting sequence of numbers to be sorted from user
	fmt.Print("Enter a sequence of unsorted integers with spaces: ")
	userInput := bufio.NewReader(os.Stdin)
	data, _, _ := userInput.ReadLine()
	numbers := strings.Split(string(data), " ")

	var seq []int

	for _, str := range(numbers){
		n, _ := strconv.Atoi(str)
		seq = append(seq, n)
	}

	BubbleSort(seq)
	
	// Prints the list of sequence in sorted order
	fmt.Println(seq)
}
