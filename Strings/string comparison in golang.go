// program to demonstrate the
// strings.Compare() function.

package main

import "fmt"
import "strings"

func main() {
    var str1,str2 string 
    fmt.Print("Enter the first string: ")
    fmt.Scanf("%s",&str1)
    fmt.Print("Enter the second string: ")
    fmt.Scanf("%s",&str2)
	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.Compare(str2, str1))
	fmt.Println(strings.Compare(str1, str1))
}
