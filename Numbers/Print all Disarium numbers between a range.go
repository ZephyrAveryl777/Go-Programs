package main 
import (
    "fmt"
    "math"
    "strings"
    )
    
func calculateLength(n int) int{
    var l int = 0 
      for n!= 0{ // while loop
        l= l+1
        n = n/10
     }
    return l
 }

func sumofDigits(num int)int{
    var sum,rem int = 0,0
    var l int = calculateLength(num) // function reference 
    for num>0{
        rem = num%10
        sum = sum+int(math.Pow(float64(rem),float64(l))) // Power integers
        num = num/10
        l = l-1
    }
    return sum
}
func main(){
    var result int = 0
    fmt.Print(strings.Repeat("_",10)+" Disarium Numbers between a range "+strings.Repeat("_",10))
    var lower,upper int 
    fmt.Printf("\nEnter lower limit: ")
    fmt.Scanf("%d",&lower)
    fmt.Printf("Enter upper limit: ")
    fmt.Scanf("%d",&upper)
    fmt.Println("\n"+strings.Repeat("_",9)+" Disarium Numbers "+strings.Repeat("_",9)+"\n")
    for i:=lower;i<=upper;i++{
        result = sumofDigits(i)
        if(result == i){
            fmt.Printf("%d\t",i)
        }
    }
}
