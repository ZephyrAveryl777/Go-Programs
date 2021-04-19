package main
import (
    "fmt"
    "strings"
)
func main(){
  var n int 
  fmt.Print("Enter Integer: ")
  fmt.Scanf("%d",&n)
  fmt.Println(strings.Repeat("*",25))  
  fmt.Printf("Multiplication table of %d \n",n)
  for i:=1;i<=20;i++{
    fmt.Printf("%d X %d = %d\n",n,i,n*i)
  }
}
