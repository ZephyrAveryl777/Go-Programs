package main
import "fmt"

func main() {
    var year int=0
    
    fmt.Printf("Enter number: ")
    fmt.Scanf("%d",&year)
    
    if( (year%4==0 && year%100!=0)||(year%4==0 && year%100==0 && year%400==0) ){
        fmt.Printf("Entered year is leap year")
    }else{
        fmt.Printf("Entered year is not leap year")
    }
}
