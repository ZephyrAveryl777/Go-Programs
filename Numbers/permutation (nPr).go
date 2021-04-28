/*
Program to print the permutation (nPr) of the given number

Output:

Enter n value: 5
Enter r value: 2
Result =  20
*/




package main

import (
    "fmt"
    )
func per( n,  r int) int{
    return fact(n)/(fact(n-r))
   
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
    
    res= per(n,r)
    fmt.Println("Result = ",res)
}
