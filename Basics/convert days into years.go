package main 
import "fmt"
func main(){
  var days,years int 
  fmt.Print("Enter days: ")
  fmt.Scanf("%d",&days)
  years = days/365
  fmt.Printf("Number of years is: %d\n",years)
 }
