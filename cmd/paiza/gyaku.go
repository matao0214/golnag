package main
import "fmt"
import "strings"

func main(){
    var n int
    fmt.Scan(&n)
    var p string
    fmt.Scan(&p)

    f := strings.Split(p, "")

    var ans string
    for i:=0; i<n; i++ {
        ans += f[n-i-1]
    }
    fmt.Println(ans)
}
