package main

import ("fmt")

func main() {
  var b,a,pos int
  var arr []int
  //getting size of the array
  fmt.Println("Enter the length of the array")
  fmt.Scan(&a)
  //getting elements for the array using for loop
  fmt.Println("Enter the elements for the array")
  for i:=0;i<a;i++{
     fmt.Scan(&b)
    arr = append(arr,b)
  }
  fmt.Println("Before removing any element",arr)
    //getting user input to get position
  fmt.Println("enter the position of element you want to remove:")
  fmt.Scanf("%d",&pos)
   fmt.Println("Removed element is",arr[pos])
  arr=append(arr[0:pos],arr[pos+1:]...)
  //Printing After removed
 fmt.Println("After removing the element ",arr)

}
