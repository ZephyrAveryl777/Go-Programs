// calculate the power of a 
// given number using the goto statement

package main

import "fmt"

func main() {
	var num,power int
	var result int = 1

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)
	var t int = num

	fmt.Print("Enter Power: ")
	fmt.Scanf("%d", &power)
	var t2 int = power
XYZ:
	result = result * num

	power = power - 1

	if power >= 1 {
		goto XYZ
	}
	
	fmt.Printf("Result of %d to the power of %d (%d^%d) is: %d",t,t2,t,t2,result)
}

