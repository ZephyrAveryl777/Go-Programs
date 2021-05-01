package main

import (
	"fmt"
)
var e,s,max int
var r,l []int
func largest(){
  //through the array storing the maximum value and comparing it with upcoiming elements and finally printing the largest value strored in max
  for i:=0;i<s;i=i+1 {
    if max<l[i]{
      max=l[i]
    }
  }
fmt.Println("Largest Element is ",max)
}
func main() {
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
largest()
}
