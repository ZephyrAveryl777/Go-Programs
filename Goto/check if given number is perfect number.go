// program to check the given number is a perfect number
// or not using the goto statement

package main

import "fmt"

func main() {
	var num,sum,rem,t  int
	var count int = 1

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)
	t = num

Repeat:
	rem = num % count
	if rem == 0 {
		sum = sum + count
	}
	count = count + 1
	if count < num {
		goto Repeat
	}

	if num == sum {
		fmt.Printf("Given number %d is perfect number",t)
	} else {
		fmt.Printf("Given number %d is not perfect number",t)
	}
}

