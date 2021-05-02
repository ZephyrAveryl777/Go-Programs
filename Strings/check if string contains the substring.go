// program to demonstrate the
// strings.ContainsAny() function

package main

import "fmt"
import "strings"

func main() {
    var str,substr string
    for true{
    fmt.Print("Enter a string: ")
    fmt.Scanf("%s",&str)
    fmt.Print("Enter a sub string: ")
    fmt.Scanf("%s",&substr)
	fmt.Println(strings.ContainsAny(str, substr))
    }
}
