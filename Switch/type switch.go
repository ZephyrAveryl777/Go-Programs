package main
  
import "fmt"
  
func main() {
    var value interface{}
    switch q:= value.(type) {
       case bool:
       fmt.Println("value is of boolean type")
       case float64:
       fmt.Println("value is of float64 type")
       case int:
       fmt.Println("value is of int type")
       default:
       fmt.Printf("value is of type: %T", q)
         
   }
}
