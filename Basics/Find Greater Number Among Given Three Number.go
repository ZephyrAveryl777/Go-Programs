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

// Method 2 

package main
import "fmt"
func main(){
  var num1,num2,num3 float64
  fmt.Printf("Enter three numbers:\t")
  fmt.Scanf("%f%f%f",&num1,&num2,&num3)
  if(num1>num2 && num1>num3){
    fmt.Printf("\nGreatest Number is:%.2f\t",num1)
  }else if(num2>num3 && num2>num1){
    fmt.Printf("\nGreatest Number is:%.2f\t",num2)
  } else if(num1==num2 && num2==num3){
    fmt.Printf("\nAll are equal")
  }else{
    fmt.Printf("\nGreatest Number is:%.2f\t",num3)
  }
}
