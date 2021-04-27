/*
'''
Pattern 
Half Diamond Star Pattern
Enter number of rows:5
*
**
***
****
*****
****
***
**
*
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Pyramid star pattern "+strings.Repeat("_",10))
  var r,i,j int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  for i=1;i<=r;i++{
      
      for j=1; j<=i;j++{
          fmt.Print("*")
      }
      fmt.Println()
  }
 
 for i=r-1;i>0;i--{
      
      for j=1; j<=i;j++{
          fmt.Print("*")
      }
      fmt.Println()
  }
     
  }
