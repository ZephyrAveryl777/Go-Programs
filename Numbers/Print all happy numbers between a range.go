package main
import (
    "fmt"
    "strings"
    )
func isHappy(num int)int {
    var rem,sum int = 0,0
    for num>0{
        rem = num%10
        sum = sum+(rem*rem)
        num = num/10
    }
    return sum
} 
func main(){
    fmt.Print(strings.Repeat("_",10)+" Happy Numbers between a range "+strings.Repeat("_",10))
    var lower,upper int 
    fmt.Printf("\nEnter lower limit: ")
    fmt.Scanf("%d",&lower)
    fmt.Printf("Enter upper limit: ")
    fmt.Scanf("%d",&upper)
    fmt.Println("\n"+strings.Repeat("_",9)+" Happy Numbers "+strings.Repeat("_",9)+"\n")
    for i:=lower;i<=upper;i++{
        var result int = i
        for result!=1 && result!=4{
            result = isHappy(result)
        }
        if(result == 1){
            fmt.Printf("%d\t",i)
        }
    }
}
