package main

import "fmt"
import "strings"

func main() {
	var str string
	fmt.Print(strings.Repeat("_",10)+" Splitting of words in string "+strings.Repeat("_",10)+"\n")
	fmt.Print("Enter a string with hyphens(-): ")
	fmt.Scanf("%s",&str)
	var strArr []string = strings.Split(str, "-")

	fmt.Println("Splited string:")
	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}

}
