package main
import "fmt"

func main(){
    var n int
    fmt.Scan(&n)
    
    var a [] int
    for i:=0; i<n; i++ {
        var g int
        fmt.Scan(&g)
        a = append(a, g)
    }
    
    var t,q int
    fmt.Scan(&t, &q)
    
    for i:=0; i<q; i++ {
        var x,y int
        fmt.Scan(&x,&y)
        if t > a[x-1] * y{
            t -= a[x-1] * y
        }
    }
    fmt.Println(t)
}
