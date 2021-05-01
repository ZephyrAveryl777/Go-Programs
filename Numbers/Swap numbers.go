package main  
import "fmt"
func main(){
    var x,y,t int // x and y are the swapping variables and t is another variable  
    fmt.Print("Enter the values of X and Y: ")
    fmt.Scanf("%d%d",&x,&y)
    fmt.Printf("before swapping numbers are: %d and %d",x,y)
    // swapping 
    t = x 
    x = y 
    y = t
    fmt.Printf("\nAfter swapping: %d and %d",x,y)

}
