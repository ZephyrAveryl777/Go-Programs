/*
'''
Pattern 
Decremental sequence number pattern rowise
Enter number of rows: 5
Enter number of columns: 5
55555
54444
54333
54322
54321
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Decremental sequence number pattern rowise to left "+strings.Repeat("_",10))
  var r,c,i,j int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  fmt.Print("Enter number of columns: ")
  fmt.Scanf("%d",&c)
  for i=r;i>=1;i--{
      
    for j=c;j>=1;j--{
      if j>i{
          fmt.Print(j)
      }else{
          fmt.Print(i)
      }
    }
    fmt.Println()
    
    
  }
}
