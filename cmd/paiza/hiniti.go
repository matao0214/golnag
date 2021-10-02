package main
import "fmt"
func main(){
    var n int
    fmt.Scan(&n)
    if n%2 == 0 {
        fmt.Println(n/24)
    } else {
        fmt.Println(n/24 + 1)
    }
}
