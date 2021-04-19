package main 
import "fmt"
func main(){
  var a int 
  fmt.Print("Enter the number you want to check: ")
  fmt.Scanf("%d",&a)
  if(a<0){
    fmt.Printf("Number %d, is Negative\n",a)
  }else if(a>0){
    fmt.Printf("Number %d, is Positive\n", a)
  }else{
    fmt.Printf("Number %d, is neither Positive nor Negative\n",a)
  }
}
