/*
'''
Pattern:
Enter number of rows: 5
         A
       B A B
     C B A B C
   D C B A B C D
 E D C B A B C D E

'''
*/

package main

import (
    "fmt"
    "strings"
    )

func main() {
  fmt.Println(strings.Repeat("_",10)+" Pattern  "+strings.Repeat("_",10))
  var r,i,j,k int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  
  space:=r
  for i=0;i<r;i++{
      for k=1;k<=space;k++{
          fmt.Print(" ")
      }
      c:='A'+i
      for j=0;j<=2*i;j++{
          if c=='A'|| j>(i){
               fmt.Printf("%c",c)
              c++
              
          }else{
               fmt.Printf("%c",c)
          c--
                   }
          
         
          
      }
      fmt.Println()
      space--
      
  }
  }
