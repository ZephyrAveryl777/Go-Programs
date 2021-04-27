/* 
The deficient number can be defined as the number for which the sum of the proper divisors is lesser than the number itself.

For example, the number 21 with its proper divisors (1, 3 and 7) has sum (11) lesser than itself.
*/
// if possible try making the code short and better 
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
