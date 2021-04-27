package main
import ("fmt")
func occurance(f []int) {
  var str []int
  str=append(str,f[0])
  for i:=0;i<len(f);i++ {
    for j:=0;j<len(str);j++ {
        if f[i]==str[j] {
          break
        }
         if j==len(str)-1{
          str=append(str,f[i])
          }
    }
  }
fmt.Println("str",str)

for i:=0;i<len(str);i++{

  count:=0
  for j:=0;j<len(f);j++ {
  if str[i]==f[j] {
    count++
  }
}
fmt.Printf("the value %d is having %d occurance\n",str[i],count)
}
}
func main() {
  var b,a int
  var arr []int
  
  //getting size of the array
  fmt.Println("Enter the length of the array")
  fmt.Scan(&a)
  //getting elements for the array using for loop
  fmt.Println("Enter the elements for the array")
  for i:=0;i<a;i++{
     fmt.Scan(&b)
    arr = append(arr,b)
  }
occurance(arr)
}
