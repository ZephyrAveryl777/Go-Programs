// program to trim string from both ends
// using Trim() function

package main

import "fmt"
import "strings"

func main() {
	var str,substr string
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&substr)
	str = str+substr
	fmt.Printf("%s", strings.Trim(str,substr))
}
