package main

import "fmt"

func main() {
  var b,a int
  var s,l []int
  //getting size of the array
  fmt.Println("Enter the length of the array")
  fmt.Scan(&a)
  //getting elements for the array using for loop
  fmt.Println("Enter the elements for the array")
  for i:=0;i<a;i++{
     fmt.Scan(&b)
    s = append(s,b)
  }
  //copying the elements of 's' array to 'l'array
    fmt.Println("Array s ",s)
   for i:=0;i<len(s);i++{
     l=append(l,s[i])
    }

  fmt.Println("Array l",l)
}
