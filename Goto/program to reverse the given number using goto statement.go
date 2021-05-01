// reverse given number 
// using the goto statement

package main

import "fmt"

func main() {
	var num,rev,rem,t int 

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)
	t = num

XYZ:

	rem = num % 10
	rev = rev*10 + rem
	num = num / 10

	if num > 0 {
		goto XYZ
	}
	
	fmt.Printf("Reverse of number %d is: %d",t,rev)
}

