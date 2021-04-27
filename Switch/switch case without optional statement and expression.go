// switch case without optional statement and expression 
package main
import "fmt"

func main() {
    var month int
    fmt.Print("Enter a number: ")
    fmt.Scanf("%d",&month)
     switch { 
       case month==1: 
            fmt.Println("Your selected month is: JANUARY") 
       case month==2: 
            fmt.Println("Your selected month is: FEBUARY") 
       case month==3: 
            fmt.Println("Your selected month is: MARCH") 
       case month==4: 
            fmt.Println("Your selected month is: APRIL") 
       case month==5: 
            fmt.Println("Your selected month is: MAY") 
       case month==6: 
            fmt.Println("Your selected month is: JUNE") 
       case month==7: 
            fmt.Println("Your selected month is: JULY")
       case month==8: 
            fmt.Println("Your selected month is: AUGUST") 
       case month==9: 
            fmt.Println("Your selected month is: SEPTEMBER") 
       case month==10: 
            fmt.Println("Your selected month is: OCTOBER")
       case month==11: 
            fmt.Println("Your selected month is: NOVEMBER") 
       case month==12: 
            fmt.Println("Your selected month is: DECEMBER") 
       default:  
            fmt.Println("Invalid Month") 
   }  
}
