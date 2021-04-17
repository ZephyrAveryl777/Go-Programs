package main

import "fmt"

func main() {
    fmt.Println("Enter a string:")
	var str string 
	fmt.Scanf("%s",&str)
	fmt.Println("\nASCII value of string:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c => %d\n", str[i], str[i])
	}
}
