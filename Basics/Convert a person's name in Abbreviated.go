package main

import "fmt"

func main() {
  var fn,mn,ln string 
  fmt.Print("Enter fullname in (firstName MiddleName )")
  fmt.Scanf("%s%s%s",&fn,&mn,&ln)
  fmt.Println("Abbrivated name:")
  fmt.Printf("%c.%c.%c\n",fn[0],mn[0],ln)
}
