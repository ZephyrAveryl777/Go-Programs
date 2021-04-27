/*
'''
Alphabet Pattern:
Enter number of rows: 5
 ABCDE
  ABCD
   ABC
    AB
     A
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Pattern  "+strings.Repeat("_",10))
  var r,i,j,k int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  space:=0
  for i=r;i>=1;i--{
      for k=1;k<=space;k++{
          fmt.Print(" ")
      }
      c:='A'
      for j=1;j<=i;j++{
          fmt.Printf("%c",c)
          c++
      }
      fmt.Println()
      space++
  }
  

  }
