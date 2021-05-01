package main 
import "fmt"
func main(){
    var val int = 1
XYZ: 
    fmt.Printf("%v => %o\n",val,val)
    val = val+1
    if val <= 100{
        goto XYZ
    }
}
