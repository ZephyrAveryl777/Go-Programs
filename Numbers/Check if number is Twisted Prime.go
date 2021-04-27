/*
A number is called a twisted prime number if it is a prime number and reverse of this number is also a prime number.
Examples: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79
*/
package main 
import "fmt"
func main(){
    var n,dig int 
    fmt.Print("Enter a number: ")
    fmt.Scanf("%d",&n)
    var num int = n
    var rev,flag int = 0,0
    for n>0{
        dig = n%10
        rev = rev*10+dig
        n = n/10
    }
    for j:=2;j<rev/2;j++{
        if((rev%j)==0){
            flag = 1
            break
        }
    }
    if(flag == 0){
        fmt.Printf("%d, is a Twisted Prime",num)
    }else{
        fmt.Printf("%d, is Not a Twisted Prime",num)
    }
}
