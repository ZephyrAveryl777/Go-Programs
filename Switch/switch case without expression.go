package main
import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println(t)
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
}
