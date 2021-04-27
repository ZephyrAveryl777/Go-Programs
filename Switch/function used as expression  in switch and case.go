package main 
import "fmt"
func myfunction(parameter string) string {  
 return parameter  
}  
func main() {  
 switch myfunction("mon") {  
 case myfunction("mon"), myfunction("tue"), myfunction("wed"), myfunction("thu"), myfunction("fri"):  
  fmt.Println("working day")  
 case myfunction("sat"), myfunction("sun"):  
  fmt.Println("Not working day: weekend")  
 }  
}  

/*
Output:
working day
*/
