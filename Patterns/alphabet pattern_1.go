/*
'''
Pattern
Enter number of rows: 5
A
BB
CCC
DDDD
EEEEE
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Pattern  "+strings.Repeat("_",10))
  var r,i,j int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  c:='A'
  for i=1;i<=r;i++{
      for j=1;j<=i;j++{
          fmt.Printf("%c",c)
      }
      fmt.Println()
      c++
  }
  

  }
