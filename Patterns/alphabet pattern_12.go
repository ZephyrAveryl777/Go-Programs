/*
'''
Alphabet Pattern:
Enter number of rows: 5
A A A A A
B B B B B
C C C C C
D D D D D
E E E E E
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Pattern  "+strings.Repeat("_",10))
  var r,i,j int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  
   c:='A'
   
  for i=0;i<r;i++{
      
      for j=0;j<r;j++{
          fmt.Printf("%c",c)
          
      }
      fmt.Println()
      
      c+=1
  }
  

  }
