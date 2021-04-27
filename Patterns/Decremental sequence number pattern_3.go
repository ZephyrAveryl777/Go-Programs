/*
'''
Pattern
Decremental sequence number pattern_3
Enter number of rows: 5
Enter number of columns: 5
54321
54322
54333
54444
55555
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
  for i=1;i<=r;i++{
      
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
