package main

import "fmt"
import "strings"

func main() {
	var str string
	var char byte // to get a character 
	var index int 
	fmt.Print(strings.Repeat("_",10)+" Check for index of character "+strings.Repeat("_",10)+"\n")
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	fmt.Print("Enter a character: ")
	fmt.Scanf("%c",&char) // inputing character
	index = strings.Index(str,string(char)) // type casted 
	fmt.Printf("\nIndex of character %c, in string %s, is: %d",char,str,index)
}
