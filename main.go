package main

import (
	"fmt"
	"os"
)

func convert(n string) {
	for i := 0; i <len(n); i++ {
		if string(n[i]) == "0" {
			fmt.Print("Zero")
		} else if string(n[i]) == "1"{
			fmt.Print("One")
		} else if string(n[i]) == "2" {
			fmt.Print("Two")
		} else if string(n[i]) == "3" {
			fmt.Print("Three")
		} else if string(n[i]) == "4" {
			fmt.Print("Four")
		} else if string(n[i]) == "5" {
			fmt.Print("Five")
		} else if string(n[i]) == "6" {
			fmt.Print("Six")
		} else if string(n[i]) == "7" {
			fmt.Print("Seven")
		} else if string(n[i]) == "8" {
			fmt.Print("Eight")
		} else if string(n[i]) == "9" {
			fmt.Print("Nine")
		}
	}

}



func main(){
	for i := 1; i < len(os.Args); i++ {
		convert(os.Args[i])
		fmt.Print(",")
	}

}
