package main
import "fmt"
import "strings"
func main(){
    var s string
    fmt.Scan(&s)
    fmt.Println(strings.Count(s, "A"))
}
