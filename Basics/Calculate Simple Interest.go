package main
import "fmt"
func main() {
  var amount,rate,time,simple_interest float64
  fmt.Printf("Enter Principal Amount: ")
  fmt.Scanf("%f",&amount)
  fmt.Printf("Enter Rate of Interest: ")
  fmt.Scanf("%f",&rate)
  fmt.Printf("Enter Period of Time: ")
  fmt.Scanf("%f",&time)
  simple_interest = (amount*rate*time)/100
  fmt.Printf("\nSimple Interest: %f",simple_interest)
}
