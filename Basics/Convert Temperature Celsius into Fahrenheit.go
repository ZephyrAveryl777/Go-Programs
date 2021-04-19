/*
Fahrenheit = ((9/5)*Celsius)+32
(9/5) is same as 1.8
*/
package main

import "fmt"

func main() {
        var Celsius, Fahrenheit float64
        fmt.Print("Enter Temperature in Celsius: ")
        fmt.Scanf("%f", &Celsius)
        Fahrenheit = (1.8 * Celsius) + 32
        fmt.Printf("Temperature in Fahrenheit: %.2f", Fahrenheit)
}
