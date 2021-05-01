package main

import (
	"fmt"
)
func main() {
var r,l []int
var s,e int
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
	for i:=0;i<s;i++ {
    for x:=i+1;x<s;x++{
      if l[i]==l[x]{
        r=append(r,l[i])
      }
	}
  }
	fmt.Println("Duplicate elements present in the given array ",r)
}
