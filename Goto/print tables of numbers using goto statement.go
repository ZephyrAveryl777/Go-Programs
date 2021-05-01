package main 
import "fmt"
func main(){
    var val,num int 
    fmt.Print("Enter Number: ")
    fmt.Scanf("%d",&num)
XYZ: 
    val = val+1
    fmt.Printf("%d*%d=%d\n",num,val,num*val)
    if val < 10{
        goto XYZ
    }
}
