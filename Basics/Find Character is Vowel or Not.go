package main
import "fmt"
func main(){
  var a byte 
  fmt.Print("Enter a character: ")
  fmt.Scanf("%c", &a)
  if(a=='a'||a=='e'||a=='i'||a=='o'||a=='u'||a=='A'||a=='E'||a=='I'||a=='O'||a=='U'){
    fmt.Printf("Character %c, is Vowel\n",a)
  }else{
    fmt.Printf("Character %c, is not Vowel\n", a)
  }
}
