package main
import "fmt"
func main() {  
 var numbArray = [3]int{1, 2, 5}  
MyLoopName:  
 for _, numb := range numbArray {  
  switch numb {  
  case 1:  
   fmt.Println("1 matched")  
   break MyLoopName  
  case 2:  
   fmt.Println("2 matched")  
   break  
  case 0:  
   fmt.Println("0 matched")  
  }  
 }  
}  

/* 
Output
1 matched
*/
