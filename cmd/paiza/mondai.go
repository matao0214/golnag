package main
import "fmt"
import "strings"

func main(){
    var s string
    var n int
    fmt.Scan(&s)
    fmt.Scan(&n)
    ans := strings.Split(s, "")
    var c int
    for i:=0; i<n; i++ {
        if ans[i] == "R" {
            c++
        }
    }
    if c >= n {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
