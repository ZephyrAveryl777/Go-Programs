package main 
import (
    "fmt"
    "strings"
)
func main(){
    var val int
    var str string 
    fmt.Print("Enter a string: ")
    fmt.Scanf("%s",&str)
    fmt.Println("\n"+strings.Repeat("_",10)+" ASCII value of string "+strings.Repeat("_",10)+"\n")
XYZ: 
    fmt.Printf("%c => %d\n",str[val],str[val])
    val = val+1
    if(val <len(str)){
        goto XYZ
    }
    
}
