package main

import (
        "fmt"
        "unsafe"
)

func main() {
        var a int
        fmt.Printf("Int is %2d bytes", unsafe.Sizeof(a))
        var b int16
        fmt.Printf("\nInt16 is %2d bytes", unsafe.Sizeof(b))
        var e int32
        fmt.Printf("\nInt32 is %2d bytes", unsafe.Sizeof(e))
        var f int64
        fmt.Printf("\nInt64 is %2d bytes", unsafe.Sizeof(f))
        var c float32
        fmt.Printf("\nFloat32 is %2d bytes", unsafe.Sizeof(c))
        var d float64
        fmt.Printf("\nFloat64 is %2d bytes", unsafe.Sizeof(d))
        var g string
        fmt.Printf("\nString is %2d bytes",unsafe.Sizeof(g))
        var h uint
        fmt.Printf("\nUnit is %2d bytes",unsafe.Sizeof(h))
        var i uint16
        fmt.Printf("\nUnit16 is %2d bytes",unsafe.Sizeof(i))
        var j uint32
        fmt.Printf("\nUnit32 is %2d bytes",unsafe.Sizeof(j))
        var k uint64
        fmt.Printf("\nUnit64 is %2d bytes",unsafe.Sizeof(k))
        var l bool
        fmt.Printf("\nBool is %2d bytes",unsafe.Sizeof(l))
        var m complex64
        fmt.Printf("\nComplex64 is %2d bytes",unsafe.Sizeof(m))
        var n complex128
        fmt.Printf("\nComplex128 is %2d bytes",unsafe.Sizeof(n))
        var o byte
        fmt.Printf("\nByte is %2d bytes",unsafe.Sizeof(o))
        
}
