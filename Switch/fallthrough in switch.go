package main  
  
import "fmt"  
  
func main() {  
 number := 1  
 switch number {  
 case 1:  
  fmt.Println("1 matched")  
  fallthrough  
 case 2:  
  fmt.Println("2 matched")  
  fallthrough  
 case 3:  
  fmt.Println("3 matched")  
 case 4:  
  fmt.Println("4 matched")  
 default:  
  fmt.Println("default matched")  
 }  
}  
/*
Output 
1 matched
2 matched
3 matched
*/
