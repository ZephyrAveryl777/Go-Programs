package main
import (
    "fmt"
    "strings"
    )
func main(){
    fmt.Print(strings.Repeat("_",10)+" Calculator "+strings.Repeat("_",10))
    var operator byte 
    var first,second float64
    fmt.Printf("\nEnter two operands: ")
    fmt.Scanf("%f %f",&first,&second)
    fmt.Printf("\nEnter an operator(+,-,*,/,%): ")
    fmt.Scanf("%c",&operator)
    switch operator{
        case '+':
        fmt.Printf("\n%.1f + %.1f = %.1f ",first,second,first+second)
        case '-':
        fmt.Printf("\n%.1f - %.1f = %.1f",first,second,first-second)
        case '*':
        fmt.Printf("\n%.1f * %.1f = %.1f",first,second,first*second)
        case '/':
        fmt.Printf("\n%.1f / %.1f = %.1f",first,second,first/second)
        case '%':
        fmt.Printf("\n%.1f % %.1f = %d",first,second,(int(first)%int(second)))
        // if no operator matches 
        default:
        fmt.Print("\nInvalid Operator")
    }
    
}
