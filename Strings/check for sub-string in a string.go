// program to check a string contains a
// specified substring

package main

import "fmt"
import "strings"

func main() {
	var str,subStr string 
	fmt.Print(strings.Repeat("_",10)+" Check for Substring "+strings.Repeat("_",10)+"\n")
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	fmt.Print("Enter a sub string: ")
	fmt.Scanf("%s",&subStr)
	
	if strings.Contains(str, subStr) == true {
		fmt.Printf("String '%s' contains sub-string '%s'", str, subStr)
	} else {
		fmt.Printf("String '%s' does not contains substring '%s'", str, subStr)
	}
}
