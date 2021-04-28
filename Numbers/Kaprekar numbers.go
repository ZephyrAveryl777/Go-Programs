/*

Program to print all Kaprekar numbers between 1 to 100


Examples:

45 =  (45)2 = 2025 =20 + 25 -45
1 = 12 = 01 = 0 + 1 = 1


OUTPUT

1
9
45
55
99

*/


package main

import (
    "fmt"
    "math")

func kaprekar(n int) bool{
    
    var sq,le, temp int
    if n==1{
        return true
    }
    if(n%10==0){
        return false
    }
    sq=n*n
    temp=sq
    le=0
    for temp!=0{
        le++
        temp/=10
    }
    
    for i:=1;i<le;i++{
        part:=int(math.Pow(float64(10),float64(i)))
        
        sum:=sq/part+sq%part
        
        if sum==n{
            return true
        }
    }
    return false
    
}

func main() {
    
    for i:=1;i<100;i++{
        
        if (kaprekar(i)){
            fmt.Println(i)
        }
    }
}
