// print the Hex value of numbers
// from 1 to 100 using goto statement.

package main

import "fmt"

func main() {
	var val int = 1

XYZ:
	fmt.Printf("%v => %X\n",val, val)
	val = val + 1

	if val <= 100 {
		goto XYZ
	}
}
