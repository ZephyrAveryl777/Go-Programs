package main

import ("fmt")

func main() {
  var b,a,e int
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

    //user input need to be searched
  fmt.Println("enter the element to be searched:")
  fmt.Scanf("%d",&e)
  for i,v:=range arr {
    if e==v {
      fmt.Printf("First occurence of %d is having index value of %d in the array\n",v,i)
      break
    }
    
  }

}
