/*Pattern
Column increment number pattern:
Enter number of rows: 5
Enter number of columns: 5
12345
12345
12345
12345
12345
*/
package main

import "fmt"

func main() {
  fmt.Println("Column increment number pattern")
  var r,c,i,j int
  fmt.Print("Enter number of rows: ")
  fmt.Scanf("%d",&r)
  fmt.Print("Enter number of columns: ")
  fmt.Scanf("%d",&c)
  for i=0;i<r;i++{
    for j=1;j<=c;j++{
      fmt.Print(j)
    }
    fmt.Println()
  }
}
