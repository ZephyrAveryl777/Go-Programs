*********************Method 1 *********************
package main

import (
	"fmt"
  "sort"
)
func des(l []int){
  //using sort package
sort.Sort(sort.Reverse(sort.IntSlice(l)))

 fmt.Println("Array in descending  order",l)
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
des(l)
}

********************Method 2 ************************************
package main

import (
	"fmt"
)
func des(l []int){
 
  var m int
for i:=0;i<len(l);i++{
  for j:=i+1;j<len(l);j++{
    if l[i]<l[j] {
      m=l[i]
      l[i]=l[j]
      l[j]=m
    }
  }

}
 fmt.Println("Array in descending order",l)
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
des(l)
}
