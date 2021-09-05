package main
import "fmt"
func main(){
    var a,b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    ans := a + b
    if ans == 21 {
        fmt.Println("Win")
    } else {
        fmt.Println(ans)
    }
}
