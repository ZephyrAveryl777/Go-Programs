// program to calculate the Highest Common Factor

package main
import "fmt"

func main() {
    var num1,num2,temp int 
    
    fmt.Printf("Enter number1: ")
    fmt.Scanf("%d",&num1)
    
    fmt.Printf("Enter number2: ")
    fmt.Scanf("%d",&num2)
    
    for (num2!=0){
        temp = num1 % num2
        num1 = num2
        num2 = temp
    }

    fmt.Printf("Highest Common Factor is:%d", num1)
}
