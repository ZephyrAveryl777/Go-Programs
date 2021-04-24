package main

import ("fmt")

func main() {
  var b,a,st,e int
  var arr,arr1 []int
  //getting size of the array
  fmt.Println("Enter the length of the array")
  fmt.Scan(&a)
  //getting elements for the array using for loop
  fmt.Println("Enter the elements for the array")
  for i:=0;i<a;i++{
     fmt.Scan(&b)
    arr = append(arr,b)
  }
  //getting starting and ending position value
  fmt.Println("enter the starting and ending position:")
  fmt.Scanf("%d %d",&st,&e)
  //slicing
  arr1=append(arr1,arr[st:e]...)
  fmt.Println("The sliced array value is",arr1)
}
