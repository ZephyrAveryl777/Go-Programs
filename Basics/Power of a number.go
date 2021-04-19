package main
import "fmt"
func main(){
  var num int 
  fmt.Print("Enter a number: ")
  fmt.Scanf("%d",&num)
  fmt.Printf("First three powers (N^1,N^2,N^3) of number %d, are :  %d %d %d",num,num,num*num,num*num*num)
}
