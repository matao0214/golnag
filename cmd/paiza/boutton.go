package main
import "fmt"
func main(){
    var n,y,m int
    fmt.Scan(&n,&y)
    fmt.Scan(&m)
    
    var ok, ng int
    for i:=0; i<m; i++ {
        x := i
        if i+1 > n {
            x -= n
        }
        var s int
        fmt.Scan(&s)
        if s == x+1 {
            ok++
        } else {
            ng++
        }
    }
    
    if ng >= y {
        fmt.Println(-1)
    } else {
        fmt.Println(ok*1000)
    }
    
}
