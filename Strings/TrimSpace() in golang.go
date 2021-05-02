// program to trim string from both ends
// using TrimSpace() function

package main

import "fmt"
import "strings"

func main() {
	var str string
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	fmt.Printf("#%s#", strings.TrimSpace(str))
}
