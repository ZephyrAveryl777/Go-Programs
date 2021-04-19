/*
Happy Number: a number which will yield 1 when it is replaced by the sum of the square of its digits repeatedly. 
Example: 
32 => 
3^2 + 2^2 = 13
1^2 + 3^2 = 10
1^2 + 0^2 = 1
other examples of happy numbers are 7,28,82,100,320.. etc
*/
package main
import "fmt"
func Happy(num int)int{
    var rem,sum int = 0,0
    for num>0{
        rem = num%10
        sum = sum+(rem*rem)
        num = num/10
    }
    return sum
}

func main(){
    var num int 
    fmt.Print("Enter a number: ")
    fmt.Scanf("%d",&num)
    var result int 
    result = num
    for result != 1 && result != 4{
        result = Happy(result)
    }
    if(result==1){
        fmt.Printf("%d is a happy number",num)
    }else if(result==4){
        fmt.Printf("%d is not a happy number",num)
    }
}
