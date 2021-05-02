// program to convert specified string
// in lowercase

package main

import "fmt"
import "strings"

func main() {
	var t,str,result string 
	fmt.Print(strings.Repeat("_",10)+" Convert string to lowercase "+ strings.Repeat("_",10)+"\n")
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	t = str
	result = strings.ToLower(str)
	fmt.Printf("String %s, in lowercase is: %s ",t,result)
}
