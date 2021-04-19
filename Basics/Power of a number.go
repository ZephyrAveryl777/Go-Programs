// Method 1

package main
import "fmt"
func main(){
  var num int 
  fmt.Print("Enter a number: ")
  fmt.Scanf("%d",&num)
  fmt.Printf("First three powers (N^1,N^2,N^3) of number %d, are :  %d %d %d",num,num,num*num,num*num*num)
}

// Method 2 

package main 
import ( 
  "fmt"
  "math"
)
func main(){
  var num,a,b,c float64 
  fmt.Print("Enter a Number: ")
  fmt.Scanf("%f",&num)
  a = float64(math.Pow(num,1))
  b = float64(math.Pow(num,2))
  c = float64(math.Pow(num,3))
  fmt.Printf("First three powers (N^1,N^2,N^3) of number %.0f, are: %.0f %.0f %.0f",num,a,b,c)
 }
