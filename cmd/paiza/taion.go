package main
import "fmt"
func main(){
    var a float64
    fmt.Scan(&a)
    if a > 37.0 {
        fmt.Println("NG")
    } else {
        fmt.Println("OK")
    }
}
