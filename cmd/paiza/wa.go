package main
import "fmt"
func main(){
    var a int
    fmt.Scan(&a)
    if a%2 == 0 {
        fmt.Println((a/2)*(a+1))
    } else {
        fmt.Println((a/2)*(a+1)+(a/2+1))
    }
}
