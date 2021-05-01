*******************Method 1***************************
import (
	"fmt"
)
var e,s,count int
var r,l []int
func length(){
  for i:=0;i<s;i=i+1 {
    count++
  }
fmt.Println("Number of elements in the array ",count)
}
func main() {
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
length()
}


*****************Method 2*******************************
package main

import (
	"fmt"
)
var e,s int
var l []int
func main() {
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
//using len() method finding the size of the array
fmt.Println("Length of the array",len(l))
}
