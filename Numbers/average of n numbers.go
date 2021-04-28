
/*

Program to print the average of n numbers

OUTPUT
Enter the number of elements : 3
Enter the elements : 2
3
4
Average =  3


*?





package main

import (
    "fmt"
    )



func main() {
    
    var n ,val,sum int
    fmt.Print("Enter the number of elements : ")
    fmt.Scanf("%d",&n)
 
    fmt.Print("Enter the elements : ")
    sum=0
    for i:=0;i<n;i++{
        fmt.Scanf("%d",&val)
        
        sum=sum+val 
        
        
    }
    sum=sum/n
    fmt.Println("Average = ",sum)
}
