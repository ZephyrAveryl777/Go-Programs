/*
'''
PPattern 
decremental sequence number pattern 
Enter number of rows: 5
Enter number of columns: 5
12345
22345
33345
44445
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
      
    for j=1;j<=c;j++{
      if j>i{
          fmt.Print(j)
      }else{
          fmt.Print(i)
      }
    }
    fmt.Println()
    
    
  }
}
