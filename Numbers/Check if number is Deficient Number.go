/* 
The deficient number can be defined as the number for which the sum of the proper divisors is lesser than the number itself.

For example, the number 21 with its proper divisors (1, 3 and 7) has sum (11) lesser than itself.
*/
// if possible try making the code short and better 
package main
import (
    "math"
    "fmt"
    )

func divsum(n int)int{
    var sum int = 0 
    var i int = 1 
    for i<=int(math.Sqrt(float64(n))){
        if(n%i ==0){
            if(n/i == i){
                sum = sum+i
            }else{
                sum = sum+(n/i)
            }
        }
        i = i+1
    }
    return sum
}

func isDef(n int)bool{
    return (divsum(n) < (2*n))
}
func main(){
    var n int 
    fmt.Print("Enter number: ")
    fmt.Scanf("%d",&n)
    if(isDef(n)){
        fmt.Print("Number is Deficient")
    }else{
        fmt.Print("Number is not Deficient")
    }
}
