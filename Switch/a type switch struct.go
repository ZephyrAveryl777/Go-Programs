package main  
  
import "fmt"  
  
func main() {  
 type Employee struct {  
  ID   int  
  Name string  
 }  
 emp := Employee{1, "Kiran"}  
 var empInterface interface{}  
 empInterface = emp  
 switch empInterface.(type) {  
 case Employee:  
  fmt.Println("Employee type matched")  
 case string:  
  fmt.Println("String type matched")  
 default:  
  fmt.Println("default type matched")  
 }  
}  

/*
Output
Employee type matched
*/
