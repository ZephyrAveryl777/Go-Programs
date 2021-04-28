package main
import (
    "fmt"
    "strings"
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
    var n,lower,upper int 
    fmt.Print(strings.Repeat("_",10)+" Abundant Numbers "+strings.Repeat("_",10))
    fmt.Print("\nEnter lower limit: ")
    fmt.Scanf("%d",&lower)
    fmt.Print("Enter upper limit: ")
    fmt.Scanf("%d",&upper)
    fmt.Print("\n"+strings.Repeat("_",5)+" Abundant Number "+strings.Repeat("_",5)+"\n")
    for i:=lower;i<=upper;i++{
        n=i
        if(More(n)){
            fmt.Printf("%d\t",n)
        }
    }
}
