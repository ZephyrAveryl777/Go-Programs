// program to check given number is a PRIME number or not
// using the goto statement

package main

import (
    "fmt"
    "strings"
)

func main() {
	var num,flag,t int 
	var count int = 1
	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)
	t = num

Repeat:
	count = count + 1

	if num%count == 0 {
		flag = 1
		goto OUT
	}

	if count < num/2 {
		goto Repeat
	}
OUT:
    fmt.Print("\n"+strings.Repeat("_",10)+" Checking for Prime Number "+strings.Repeat("_",10)+"\n")
	if flag == 1 {
		fmt.Printf("Number %d, is not a PRIME number",t)
	} else {
		fmt.Printf("Number %d, is a PRIME number",t)
	}
}

