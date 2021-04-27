/*
Pronic Number can be expressed as N x (N+1)
Example: 
6 = 2x3
20 = 4x5
72 = 8x9

*/
package main 
import (
    "fmt"
    "strings"
)

func isPronic(num int)bool{
    var flag bool = false
    for j:=1;j<=num;j++{
        if(j*(j+1) == num){
            flag = true
            break
        }
    }
    return flag
}
func main(){
    fmt.Print(strings.Repeat("_",15)+" Pronic Numbers between a range "+strings.Repeat("_",15))
    var lower,upper int 
    fmt.Printf("\nEnter lower limit: ")
    fmt.Scanf("%d",&lower)
    fmt.Printf("Enter upper limit: ")
    fmt.Scanf("%d",&upper)
    fmt.Println("\n"+strings.Repeat("_",9)+" Pronic Numbers "+strings.Repeat("_",9)+"\n")
    for i:=lower;i<=upper;i++{
        if(isPronic(i)){
            fmt.Printf("%d\t",i)
        }
    }
}
