// program to trim specified characters from
// the left side of string using TrimLeft() function

package main

import "fmt"
import "strings"

func main() {
	str := "!%Hello World!%"

	fmt.Printf("%s", strings.TrimLeft(str, "!%"))
}
