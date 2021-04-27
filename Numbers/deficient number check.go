package main
import (
    "fmt"
    "math"
)
func  divsum(n int)int{
    var i int
    var sum int = 0
    var c int = int(math.Sqrt(int(n))) // type errors 
    for i:=1;i<=c;i++{
        if(n%i == 0){
            if((n/i)==i){
                sum = sum+i
            }else{
                sum = sum+i
                sum = sum+(n/i)
            }
        }
    }
    return sum
}
func isDef(n int) int{
    return (divsum(n)<(2*n)) // type error 
}
func main(){
    var n int
    fmt.Print("Enter a number: ")
    fmt.Scanf("%d",&n)
    if(isDef(n)==0){
        fmt.Print("Number is  deficient.")
    }else{
        fmt.Print("Number is not deficient")
    }
}
