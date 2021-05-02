// program to check the specified substring
// is the prefix of the given string.

package main

import "fmt"
import "strings"

func main() {
	var str,substr string 
	var prefix,suffix bool = false,false
	fmt.Print(strings.Repeat("_",10)+" Check for Prefix "+strings.Repeat("_",10)+"\n")
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	fmt.Print("Enter a substring: ")
	fmt.Scanf("%s",&substr)

	prefix = strings.HasPrefix(str, substr)
	suffix = strings.HasSuffix(str,substr)

	if prefix == true || suffix == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
