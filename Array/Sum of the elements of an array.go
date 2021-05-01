package main
import (
	"fmt"
)
var e,s,total int
var r,l []int
func sum(){
  //summing all the values of the array to the variable 'total'
  for i:=0;i<s;i=i+1 {
    total+=l[i]
  }
fmt.Println("Sum of the Elements is ",total)
}
func main() {
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
sum()
}
