package main 
import "fmt"
func main(){
  var num1,num2,num3 float64
  fmt.Print("Enter Three Numbers: ")
  fmt.Scanf("%f%f%f",&num1,&num2,&num3)
  if(num1>=num2){
    if(num1>=num3){
      fmt.Printf("%2.f Is the largest number: ",num1)
    }else{
      fmt.Printf("%2.f Is the largest number: ",num3)
    }
  }else{
    if(num2>=num3){
      fmt.Printf("%2.f Is the largest number: ",num1)
    }else{
      fmt.Printf("%2.f Is the largest number: ",num3)
    }
  }
}
