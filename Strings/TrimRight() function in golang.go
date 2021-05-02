// program to trim specified characters from
// the right side of string using TrimRight() function

package main

import "fmt"
import "strings"

func main() {
	str := "!%Hello World!%"

	fmt.Printf("%s", strings.TrimRight(str, "!%"))
}
