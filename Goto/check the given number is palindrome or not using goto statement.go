package main 
import "fmt"
func main(){
    var num,rev,rem,temp int 
    fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)

	temp = num
A:

	rem = num % 10
	rev = rev*10 + rem
	num = num / 10

	if num > 0 {
		goto A
	}

	if temp == rev {
		fmt.Printf("Number %d, is a palindrome",temp)
	} else {
		fmt.Printf("Number %d, is not a palindrome",temp)
	}
}
