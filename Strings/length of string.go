package main

import "fmt"
import "strings"

func main() {
	var str string
	fmt.Print(strings.Repeat("_",10)+" Check for index of character "+strings.Repeat("_",10)+"\n")
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	length := len(str)
	fmt.Printf("Length of string %s, is: %d",str,length)
}
