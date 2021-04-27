/*
'''
Pattern 
box pattern with number 1 or 0:
Enter number of rows: 5
Enter number of columns: 5
11111
1   1
1   1
1   1
11111
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Pattern Box "+strings.Repeat("_",10))
  var r,c,i,j int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  fmt.Print("Enter number of columns: ")
  fmt.Scanf("%d",&c)
  for i=1;i<=r;i++{
      
      for j=1; j<=c;j++{
          if(j==1||j==c||i==1||i==r){
             fmt.Print("1")
          }else{
              fmt.Print(" ")
          }
      }
      fmt.Println()
  }
 
     
  }
