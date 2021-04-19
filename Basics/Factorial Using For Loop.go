package main
import "fmt"
func main(){
  var fact int = 1 
  var num int 
  fmt.Print("Enter a number to find Factorial: ")
  fmt.Scanf("%d",&num)
  for i:=1;i<=num;i++{
    fact = fact*i
  }
  fmt.Printf("Factorial %d is: %d\n",num,fact)
}
