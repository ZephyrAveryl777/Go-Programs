package main

import "fmt"

func main() {
    fmt.Println("Enter a string:")
	var str string 
	fmt.Scanf("%s",&str)
	fmt.Println("\nASCII value of string:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c => %d\n", str[i], str[i])
	}
}

// Method 2 
 
package main
import "fmt"
func main(){
	lowercase := "abcdefghijklmnopqrstunwxyz"
	for _,c := range lowercase{
		fmt.Println(c)
		}
	uppercase := "ABCDEFGHIJKLMNOPQRSTUNWXYZ"
	for _,c := range uppercase{
		fmt.Println(c)
		}
	numbers := "0123456789"
	for _,n := range numbers {
		fmt.Println(n)
	}
}
