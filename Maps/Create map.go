// Map is aunordered collection of key-value pair in golang, key and value can be any type in golang .Key is unique in map,each value can be accessed using key

//Declaration
package main

import "fmt"

func main() {
  var m map[string]int
  fmt.Printf("%T\n",m)
  fmt.Printf("%v",m)

}


//Initialization
package main

import "fmt"

func main() {
  m:=make( map[string]int)
  fmt.Printf("%T\n",m)
  fmt.Printf("%v",m)
}

//To create and initialize literally
package main

import "fmt"

func main() {
  m:=map[string]int{
    "a" : 123,
    "b" : 23,
    "c" : 56,
  }
  fmt.Printf("%v",m)
}

