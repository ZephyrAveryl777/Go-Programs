/*
'''
Alphabet Pattern:
Enter number of rows: 5
     A
    CCC
   EEEEE
  GGGGGGG
 IIIIIIIII
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
  
  space:=r
   c:='A'
   
  for i=0;i<r;i++{
      for k=1;k<=space;k++{
          fmt.Print(" ")
      }
     
      for j=0;j<=2*i;j++{
          fmt.Printf("%c",c)
          
      }
      fmt.Println()
      space--
      c+=2
  }
  

  }
