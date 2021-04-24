******************************Method 1***************************


package main

import "fmt"

func main() {
  var b int
  var a=4
  var s []int
  fmt.Println("Enter the length of the array")
  fmt.Scan(&a)
  fmt.Println("Enter the elements for the array")
   
  for i:=0;i<a;i++{
     fmt.Scan(&b)
    s = append(s, b)
  
  }
    fmt.Println("original array is ",s)
   for i:=0;i<len(s)/2;i++{
      j:=len(s)-i-1
      s[i],s[j]=s[j],s[i]
    }

  fmt.Println("Reversed array is",s)
}


************************************Method 2**********************************************
package main

import "fmt"

func main() {
  var b int
  var a=4
  var s,l []int
  fmt.Println("Enter the length of the array")
  fmt.Scan(&a)
  fmt.Println("Enter the elements for the array")
  for i:=0;i<a;i++{
     fmt.Scan(&b)
    s = append(s,b)
  }
    fmt.Println("original array is ",s)
   for i:=len(s)-1;i>=0;i--{
     l=append(l,s[i])
    }

  fmt.Println("Reversed array is",l)
}
