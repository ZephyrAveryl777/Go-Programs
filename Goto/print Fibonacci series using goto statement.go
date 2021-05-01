package main
import (
    "fmt"
    "strings"
)
func main(){
    var num1,num3,r int
    var num2 int = 1 
    var num4 int = 3
    fmt.Print("Enter the upper limit of range: ")
	fmt.Scanf("%d",&r)
	fmt.Print("\n"+strings.Repeat("_",10)+" Fibonacci Series "+strings.Repeat("_",10)+"\n")
    fmt.Printf("%d", num1)
	fmt.Printf(" %d", num2)
	

Repeat:
	num3 = num1 + num2
	fmt.Printf(" %d", num3)
	num1 = num2
	num2 = num3
	num4 = num4 + 1
	
	if num4 <= r {
		goto Repeat
	}
}
