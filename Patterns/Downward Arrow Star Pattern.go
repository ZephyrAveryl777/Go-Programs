/*
'''
Pattern
Downward Arrow Star Pattern:
Enter number of rows: 5
*                 * 
* *             * * 
* * *         * * * 
* * * *     * * * * 
* * * * * * * * * * 
'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Decremental sequence number pattern rowise to left "+strings.Repeat("_",10))
  var r,i,j,t int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  
  for i=0;i<r;i++{
      
    for j=0;j<2*r;j++{
        t=2*r -i-1
      if j<=i || j==t || j>t{
          fmt.Print("*")
      }else{
          fmt.Print(" ")
      }
    }
    fmt.Println()
    
    
  }
}
