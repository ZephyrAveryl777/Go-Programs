//  calculate the factorial of a
// given number using the goto statement

package main

import "fmt"

func main() {
	var num int 
	var fact int = 1
	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)
    var temp int = num
	if num < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		if num == 0 {
			fact = 1
		} else {
		XYZ:
			fact = fact * num

			num = num - 1

			if num > 1 {
				goto XYZ
			}
		}
		fmt.Printf("Factorial of %d is: %d",temp,fact)
	}
}

