/*
'''
Pattern
Right Arrow Pattern:
Enter number of rows: 5
*****
  ****
    ***
      **
        *
      **
    ***
  ****
*****
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Pattern Box "+strings.Repeat("_",10))
  var r,c,i,j,k int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  c=r
  for i=0;i<r;i++{
      for k=0;k<=2*i;k++{
          fmt.Print(" ")
      }
      
      for j=1; j<=c;j++{
          
             fmt.Print("*")
      }
      c--
      fmt.Println()
  }
 
 c=2
  for i=r-2;i>=0;i--{
      for k=0;k<=2*i;k++{
          fmt.Print(" ")
      }
      
      for j=1; j<=c;j++{
          
             fmt.Print("*")
      }
      c++
      
      fmt.Println()
  }
 
     
  }
