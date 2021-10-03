package main
import "fmt"
func main(){
    var a,b int
    fmt.Scan(&a, &b)
    if (a >= 25 && b <= 40){
        fmt.Println("No")
    } else if a >= 25 || b <= 40 {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
