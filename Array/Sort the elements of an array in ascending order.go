***********************************Method 1***********************************************
//By using sort package
package main

import (
	"fmt"
  "sort"
)
func as(l []int){
  //If the given array is of type string ,can use sort.Strings() to get ascending order
sort.Ints(l)
fmt.Println(l)
}
func main() {
var l []int
var s,e int
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
as(l)
}

*************************************Method 2**************************************
//without using sort package
package main

import (
	"fmt"
)
func as(l []int){
 // var r []int
  var m int
for i:=0;i<len(l);i++{
  for j:=i+1;j<len(l);j++{
    if l[i]>l[j] {
      //swapping the values
      m=l[i]
      l[i]=l[j]
      l[j]=m
    }
  }
}
 fmt.Println("Array in ascending order",l)
 }
func main() {
var l []int
var s,e int
fmt.Print("Enter the size of the array:")
fmt.Scan(&s)
fmt.Println("Enter the elements of the array:")
for i:=0;i<s;i++{
  fmt.Scan(&e)
  l=append(l,e)
}
as(l)
}
