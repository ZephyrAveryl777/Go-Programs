package main

import ("fmt"
"os"
"strings")

func main() {
  var b,a int
  var op string
  var s []int
  //getting size of the array
  fmt.Println("Enter the length of the array")
  fmt.Scan(&a)
  //getting elements for the array using for loop
  fmt.Println("Enter the elements for the array")
  for i:=0;i<a;i++{
     fmt.Scan(&b)
    s = append(s,b)
  }
  for {
  fmt.Println("Enter Y to find value ,N to exit")
  fmt.Scan(&op)
  op=strings.ToUpper(op)
  if(op=="Y"){
    //getting user input to get position
  fmt.Println("enter the position of element you want to access:")
  fmt.Scanf("%d",&a)
  if(a<len(s)){
    //printing the elementof that location
    fmt.Printf("element in position %d of s array is %d\n",a,s[a])

  }else{
    fmt.Println("Invalid choice,Please enter the number within the length of the array")
  }
  }
   if (op=="N") {
      os.Exit(0)
  }
  }

}
