package main 
import "fmt"
func main(){
    var n,sum,t,remainder int 
    fmt.Print("Enter a number: ")
    fmt.Scanf("%d",&n)
    t = n 
    for (t!=0){
        remainder = t%10
        sum = sum+remainder 
        t = t/10 
    }
    fmt.Printf("sum of digits of %d is: %d\n",n,sum)
}

// **************** Method 2 *************************
package main 
import "fmt"
func main(){
    var n,sum,remainder int 
    fmt.Print("Enter a number: ")
    for  fmt.Scanf("%d",&n);n!=0;n=n/10{
        remainder = n%10
        sum = sum+remainder 
    }
    fmt.Printf("Sum of digits is: %d",sum)
}

