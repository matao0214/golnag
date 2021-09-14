package main
import "fmt"
func main(){
    var a string
    var b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    
    if a == "chocolate" {
        fmt.Println(b*2)
    } else {
        fmt.Println(b*5)
    }
}
