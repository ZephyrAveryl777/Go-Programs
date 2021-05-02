// program to demonstrate the
// strings.Join() function.

package main

import "fmt"
import "strings"

func main() {
	str := []string{"hello","I","am","fine","thank","you"}
	var result string

	result = strings.Join(str, "_")
	fmt.Println("Result: ", result)
}
