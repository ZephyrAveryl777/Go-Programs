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
