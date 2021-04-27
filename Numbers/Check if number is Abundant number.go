/*
Abundant number defined as the number for which the sum of its proper divisors is greater than the number itself.
Example :
12 => (1,2,3,4,6) sum of (1,2,3,4,6) is 16; 16>12 hence 12 is abundant number 
other examples: 12,18,20,24,30,36
*/
package main
import (
    "fmt"
    "math"
)
func sumfinder(n int)int{
    var sum int = 0
    for i:=1;i<=int(math.Sqrt(float64(n)));i++{
        if(n%i==0){
            if(n/i==i){
                sum=sum+i
            }else{
                sum=sum+i+(n/i)
            }
        }
    }
    sum=sum-n
    return sum
}

func More(n int)bool{
    return (sumfinder(n) > n)
} 

func main(){
    var n int 
    fmt.Print("Enter a number: ")
    fmt.Scanf("%d",&n)
    if(More(n)){
        fmt.Print("Number is an Abundant number")
    }else{
        fmt.Print("Number is not an Abundant number")
    }
}
