// switch case with optional statement and expression
package main
import "fmt"

func main() {
    var n int
    fmt.Print("Enter your choice: ")
    fmt.Scanf("%d",&n)
     switch month:=n; month{ 
       case 1: 
            fmt.Printf("Entered number is: %d, month associated is: JANUARY",n) 
       case 2: 
            fmt.Printf("Entered number is: %d, month associated is: FEBURARY",n) 
       case 3: 
            fmt.Printf("Entered number is: %d, month associated is: MARCH",n) 
       case 4: 
            fmt.Printf("Entered number is: %d, month associated is: APRIL",n) 
       case 5: 
            fmt.Printf("Entered number is: %d, month associated is: MAY",n) 
       case 6: 
            fmt.Printf("Entered number is: %d, month associated is: JUNE",n) 
       case 7: 
            fmt.Printf("Entered number is: %d, month associated is: JULY",n)
       case 8: 
            fmt.Printf("Entered number is: %d, month associated is: AUGUST",n) 
       case 9: 
            fmt.Printf("Entered number is: %d, month associated is: SEPTEMBER",n) 
       case 10: 
            fmt.Printf("Entered number is: %d, month associated is: OCTOBER",n)
       case 11: 
            fmt.Printf("Entered number is: %d, month associated is: NOVEMBER",n) 
       case 12: 
            fmt.Printf("Entered number is: %d, month associated is: DECEMBER",n) 
       default:  
            fmt.Println("Invalid Month") 
   }  
}
