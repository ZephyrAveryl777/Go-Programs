package main
import (
    "fmt"
    "strings"
)
func main(){
    var str string 
    fmt.Print(strings.Repeat("_",10)+" Get characters from string using index "+strings.Repeat("_",10)+"\n")
    fmt.Print("Enter a string: ")
    fmt.Scanf("%s",&str)
    for i:=0;i<len(str);i++{
        fmt.Printf("%d => %c\n",i,str[i])
    }
    fmt.Print("The string is: ")
    for i:=0;i<len(str);i++{
        fmt.Printf("%c",str[i])
    }
}
