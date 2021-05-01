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
