package main

import ("fmt")
var arr []int
func update(st,val int) {
  //replacing the value
  arr[st]=val
  fmt.Println("Array after updation",arr)
}
func main() {
  var b,a,st,val int
  
  //getting size of the array
  fmt.Println("Enter the length of the array")
  fmt.Scan(&a)
  //getting elements for the array using for loop
  fmt.Println("Enter the elements for the array")
  for i:=0;i<a;i++{
     fmt.Scan(&b)
    arr = append(arr,b)
  }
  fmt.Println("Array before updation",arr)

  //getting index value and number for updation
  fmt.Println("enter the position to be updated & value to be replaced")
  fmt.Scanf("%d %d ",&st,&val)
  update(st,val)
}
