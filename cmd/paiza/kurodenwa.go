package main
import "fmt"
import "strings"
import "strconv"

func main(){
    var s string
    fmt.Scan(&s)
    
    var ans int
    b := strings.Split(s, "")
    for i:=0; i<len(s); i++ {
        if b[i] == "-" {
            ans += 0
        } else if b[i] == "0" {
            ans += 24
        } else {
            var num int
            num, _ = strconv.Atoi(b[i])
            ans += (num + 2) * 2
        }
    }
    fmt.Println(ans)
}
