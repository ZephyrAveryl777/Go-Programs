package main
import "fmt"
func main(){
    var b,h,area float64
    fmt.Printf("Enter Height and Base of Triangle: ")
    fmt.Scanf("%f %f",&h,&b)
    area = (h*b)/2
    fmt.Printf("Area of Triangle is: %f\n",area)
}
