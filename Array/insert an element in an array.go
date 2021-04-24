package main

import ("fmt")

func insert(arr []int,pos,val int){
  //incrementing size of the array by 1
  arr=arr[0:len(arr)+1]
  //copying all the elements in position incremented by 1 value 
  copy(arr[pos+1:],arr[pos:])
  //inserting value in the given index position
  arr[pos]=val
  fmt.Println(arr)
}
func main() {
  var b,a,pos,val int
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
    //getting user input to get position
  fmt.Println("enter the position of element you want to insert:")
  fmt.Scanf("%d",&pos)
  //getting value to be inserted
  fmt.Println("enter the value to be inserted")
  fmt.Scanf("%d",&val)
  insert(arr,pos,val)

}
