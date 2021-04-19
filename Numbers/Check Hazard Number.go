/*
Harshad Number:  if number is divisible by the sum of its digit.
Example:
156 => 1+5+6 = 12, 156 is divisible by 12 
*/
package main 
import "fmt"
func main(){
    var num,n int 
    fmt.Print("Enter a number: ")
    fmt.Scanf("%d",&num)
    var rem,sum int = 0,0
    n = num
    for num>0{
        rem = num%10
        sum = sum+rem
        num = num/10
    }
    if(n%sum == 0){
        fmt.Printf("%d is a harshard number",n)
    }else{
        fmt.Printf("%d is not a harshard number",n)
    }
}
