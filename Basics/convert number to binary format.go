package main 
import "fmt"
func main(){
    var n int 
    fmt.Print("Enter number here: ")
    fmt.Scanf("%d",&n)
    fmt.Printf("Binary value of %d is: %b\n",n,n)
    fmt.Printf("Octal value of %d is: %o\n",n,n)
    fmt.Printf("Hexa value of %d is: %X\n",n,n)
}
