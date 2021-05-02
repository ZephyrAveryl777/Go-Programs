package main 
import (
    "fmt"
    "strings"
)
func main(){
    var count int 
    var str string 
    fmt.Print("Enter a string: ")
    fmt.Scanf("%s",&str)
    for i:=0;i<len(str);i++{
        if(strings.Contains(str,",") || strings.Contains(str,".") || strings.Contains(str,"'") || strings.Contains(str,":") ||strings.Contains(str,";")){
            count = count+1
        }
    }
    fmt.Printf("Total number of punctuations in the string "%s" is: %d",str,count)
}
