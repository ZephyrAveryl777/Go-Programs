/*
Disarium Number: 175 =  1^1 + 7^2 + 5^3 = 1 + 49 + 125 = 175
other examples are: 89,135,518 etc
*/
package main 
import (
    "fmt"
    "math"
    )
    
func calculateLength(n int) int{
    var l int = 0 
      for n!= 0{ // while loop
        l= l+1
        n = n/10
     }
    return l
 }

func main(){
    var num,n int
    fmt.Print("Enter number: ")
    fmt.Scanf("%d",&num)
    var sum,rem int = 0,0
    var l int = calculateLength(num) // function reference 
    n = num // imporant to have an original copy
    for num>0{
        rem = num%10
        sum = sum+int(math.Pow(float64(rem),float64(l))) // Power integers
        num = num/10
        l = l-1
    }
    if(sum == n){
        fmt.Printf("%d is a disarium number",n)
    }else{
        fmt.Printf("%d is not a disarium number",n)
    }
}
