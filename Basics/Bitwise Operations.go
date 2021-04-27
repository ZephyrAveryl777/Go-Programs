// LEFT SHIFT 
package main
import "fmt"
func main(){
    var n int = 0 
    var res int = 0 
    fmt.Printf("Enter number: ")
    fmt.Scanf("%d",&n)
    var l int 
    fmt.Printf("Enter the value to be shifted by: ")
    fmt.Scanf("%d",&l)
    res = n<<l
    fmt.Printf("Original Number is %d\nBinary of Original Number %d is %b\nResult of Left shift is %d\nBinary of result is %b\n",n,n,n,res,res)
}
// RIGHT SHIFT
package main
import "fmt"
func main(){
    var l,num,res int  
    fmt.Printf("Enter number: ")
    fmt.Scanf("%d",&num)
    fmt.Printf("Enter value by which to be shifted: ")
    fmt.Scanf("%d",&l)
    res = num>>l
    fmt.Printf("Original Number is %d\nBinary of Original Number %d is %b\nResult of Right shift is %d\nBinary of result is %b\n",num,num,num,res,res)
}

// BINARY AND 
package main
import "fmt"
func main(){
    var num1,num2,res int  
    fmt.Printf("Enter number1: ")
    fmt.Scanf("%d",&num1)
    fmt.Printf("Enter number2: ")
    fmt.Scanf("%d",&num2)
    res = num1 & num2 
    fmt.Printf("First number is: %d\nBinay of First Number %d is: %b\nSecond Number is: %d\nBinary of Second Number is: %b\nResult of bitwise and between %d and %d is: %d\nBinary of the result is:%b\n",num1,num1,num1,num2,num2,num1,num2,res,res)
}
