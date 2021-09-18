package main
import "fmt"
func main(){
    var a int
    fmt.Scan(&a)
    if a >= 1000 {
        fmt.Println(a+3)
    } else {
        fmt.Println(a)
    }
}
