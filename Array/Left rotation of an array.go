package main

import (
	"fmt"
)
func main() {
var r,l []int
var n,s,e int
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
fmt.Print("Enter the number of times rotation needs to be done: ")
fmt.Scan(&n)
if n>s{
n=n%s
}
	for i:=n;i<len(l);i++ {
	r=append(r,l[i])
	}
	for x:=0;x<n;x++{
	r=append(r,l[x])
	}
	fmt.Println("Before rotation",l)
	fmt.Println("After rotation",r)
}
