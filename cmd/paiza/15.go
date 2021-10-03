package main
import "fmt"
func main(){
    var ans int
    for i:=0; i<4; i++ {
        var a int
        fmt.Scan(&a)
        ans += a
    }
    fmt.Println(15-ans)
}
