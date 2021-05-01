// read the name and print it 5 times

package main

import "fmt"

func main() {

	var val int = 0
	var name string

	fmt.Print("Enter Name: ")
	fmt.Scanf("%s", &name)

XYZ:
	val = val + 1

	fmt.Printf("%d. My name is: %s\n",val,name)
	if val < 5 {
		goto XYZ
	}
}
