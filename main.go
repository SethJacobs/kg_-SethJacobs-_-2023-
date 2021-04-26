package main

import (
	"fmt"
	"os"
)

func convert(n string) {
	for i := 0; i <len(n); i++ {
		if string(n[i]) == "0" {
			fmt.Print("Zero")
		}
		if string(n[i]) == "1"{
			fmt.Print("One")
		}
		if string(n[i]) == "2" {
			fmt.Print("Two")
		}
		if string(n[i]) == "3" {
			fmt.Print("Three")
		}
		if string(n[i]) == "4" {
			fmt.Print("Four")
		}
		if string(n[i]) == "5" {
			fmt.Print("Five")
		}
		if string(n[i]) == "6" {
			fmt.Print("Six")
		}
		if string(n[i]) == "7" {
			fmt.Print("Seven")
		}
		if string(n[i]) == "8" {
			fmt.Print("Eight")
		}
		if string(n[i]) == "9" {
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
