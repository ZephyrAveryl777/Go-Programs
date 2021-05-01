package main

import (
	"fmt"
)
var s,e int
var r,l []int
func evenpos(){
  fmt.Println("Elements at odd position are ")
  //starting from 1st position(not with 0th position) and increment i value by 2 
  for i:=1;i<s;i=i+2 {
    fmt.Print(l[i]," ")
  }
}
func main() {
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
evenpos()
}
