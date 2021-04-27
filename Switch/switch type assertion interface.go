package main  
  
import "fmt"  
  
func main() {  
 var interfaceDemo interface{} = 45  
  
 switch interfaceDemo.(type) {  
 case string:  
  fmt.Println("string type matched")  
 case int:  
  fmt.Println("int type matched")  
 default:  
  fmt.Println("default type matched")  
  
 }  
}  

/*
Output:

int type matched
*/
