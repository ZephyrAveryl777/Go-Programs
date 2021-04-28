/*
Program to print all prime numbers between 1 and 100

Out:
2 3 4 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97 

*/




package main

import (
    "fmt"
    )

func prime(n int) bool{
    
    if n==1{
        return false
    }
   
  
    
    for i:=2;i<n/2;i++{
        if(n%i==0){
            return false
        }
        
    }
    return true
    
}

func main() {
    
    for i:=1;i<100;i++{
        
        if (prime(i)){
            fmt.Print(i, " ")
        }
    }
}
