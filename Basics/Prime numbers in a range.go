package main 
import "fmt"
func main(){
    var r int 
    fmt.Print("Enter upper range: ")
    fmt.Scanf("%d",&r)
    for i:=2;i<r+1;i++{
        var k int = 0
        for j:=2;j<(i/2)+1;j++{
            if(i%j==0){
                k=k+1
            }
        }
        if(k<=0){
            fmt.Printf("%d\t",i)
        }
    }
}
