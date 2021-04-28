
/*
Program to print the combination (nCr) of the given number


OUTPUT

Enter n value: 6
Enter r value: 4
Result =  15
*/

package main

import (
    "fmt"
    )
func combin( n,  r int) int{
    return fact(n)/(fact(n-r)*fact(r))
   
}
func fact( n int)int{
    if(n==1){
        return 1
    }else{
        return n*fact(n-1)
    }
}

func main() {
    
    var n , r, res int
    fmt.Print("Enter n value: ")
    fmt.Scanf("%d",&n)
    fmt.Print("Enter r value: ")
    fmt.Scanf("%d",&r)
    
    res= combin(n,r)
    fmt.Println("Result = ",res)
}
