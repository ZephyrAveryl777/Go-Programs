/*
Pattern 
Decremental sequence number pattern rowise  to left
Enter number of rows: 5
Enter number of columns : 5
5 5 5 5 5 
4 4 4 4 5 
3 3 3 4 5 
2 2 3 4 5 
1 2 3 4 5 
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
      
    for j=1;j<=r;j++{
      if j<i{
          fmt.Print(i)
      }else{
          fmt.Print(j)
      }
    }
    fmt.Println()
    
    
  }
}
