/*
'''
Pattern 
Pyramid star pattern
Enter number of rows: 5
    *
   ***
  *****
 *******
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Pyramid star pattern "+strings.Repeat("_",10))
  var r,i,j,t,temp int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  temp=r;
  for j=1;j<=r;j++{
      
    for i=1;i<=temp-1;i++{
     
          fmt.Print(" ")
      }
      for t=1;t<=2*j-1;t++{
          fmt.Print("*")
          
      }
      fmt.Println()
      temp--
      
    }
    
    
  }
