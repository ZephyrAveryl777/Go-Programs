package main 
import "fmt"
func main(){
    var radius float64
    const PI float64 = 3.14 
    var area,circumference float64
    fmt.Printf("\nEnter radius of circle: ")
    fmt.Scanf("%f",&radius)
    area = PI*radius*radius
    fmt.Printf("\nArea of circle: %f",area)
    circumference = 2*PI*radius
    fmt.Printf("\nCircumference of circle: %f",circumference)
    
}
