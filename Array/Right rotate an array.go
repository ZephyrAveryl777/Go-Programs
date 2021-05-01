package main

import (
	"fmt"
)
var e,s,re int
var r,l []int
func rightRotate(){
  if re>s{
    re=re%s
  }
  for i:=s-re;i<=s-1;i=i+1 {
    r=append(r,l[i])
  }
 for i:=0;i<s-re;i=i+1{
    r=append(r,l[i])
  }
  
  
fmt.Println("After rotation ",r)
}
func main() {
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
fmt.Print("Enter the number of times array need to be rotated :")
fmt.Scan(&re)
fmt.Println("Before rotation",l)
rightRotate()

}
