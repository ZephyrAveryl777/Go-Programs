// check the given number is an
// Armstrong number or not using the goto statement

package main

import "fmt"

func main() {
	var num,rem,res,temp int

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)

	temp = num
A:

	rem = num % 10
	res = res + (rem * rem * rem)
	num = num / 10

	if num > 0 {
		goto A
	}

	if temp == res {
		fmt.Printf("Number %d, is an armstrong number",temp)
	} else {
		fmt.Printf("Number %d, is not an armstrong number",temp)
	}
}

