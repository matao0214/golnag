package main
import "fmt"
func main(){
    var t int
    fmt.Scan(&t)
    if t>=30 && t<35 {
        fmt.Println("M")
    } else {
        fmt.Println(t)
    }
}
