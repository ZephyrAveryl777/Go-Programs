// program to check the specified substring
// is the prefix of the given string.

package main

import "fmt"
import "strings"

func main() {
	var str,substr string 
	var result int 
	fmt.Print(strings.Repeat("_",10)+" Count Occurance of Character "+strings.Repeat("_",10)+"\n")
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s",&str)
	fmt.Print("Enter a  substring: ")
	fmt.Scanf("%s",&substr)
    result = strings.Count(str,substr)
	fmt.Printf("Total occurance of %s in string %s is: %d",substr,str,result)
}
