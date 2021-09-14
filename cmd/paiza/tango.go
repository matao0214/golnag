package main
import "fmt"
func main(){
    var n int
    var s []string
    fmt.Scan(&n)
    
    for i:=0; i<n; i++ {
        var r string
        fmt.Scan(&r)
        s = append(s, r)
    }
    fmt.Println(s[n-1])
}
