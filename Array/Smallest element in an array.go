package main

import (
	"fmt"
)
var e,s,small int
var r,l []int
func smallest(){
  //storing 1st element of the array in small comparing it with other elements,found any small element replace the value stored in small 
  small=l[0]
  for i:=1;i<s;i=i+1 {
    if small>l[i]{
      small=l[i]
    }
  }
fmt.Println("Smallest Element is ",small)
}
func main() {
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
smallest()
}
