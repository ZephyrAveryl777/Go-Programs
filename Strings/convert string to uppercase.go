// program to convert specified string
// in uppercase

package main

import "fmt"
import "strings"

func main() {
	var t,str,result string 
	fmt.Print(strings.Repeat("_",10)+" Convert string to uppercase "+ strings.Repeat("_",10)+"\n")
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	t = str
	result = strings.ToUpper(str)
	fmt.Printf("String %s, in uppercase is: %s ",t,result)
}
