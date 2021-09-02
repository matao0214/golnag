package main
import "fmt"
import "strings"

func main(){
    var h,w int
    fmt.Scan(&h,&w)

    var s []string
    for i:=0; i<h; i++ {
        var score string
        fmt.Scan(&score)
        s = append(s, score)
    }
    
    
    var ans int
    for i:=0; i<h; i++ {
        f := strings.Split(s[i], "")
        for q:=0; q<w; q++ {
            if f[q] == "o" {
                ans += i*w + 1 + q
            }
        }
    }
    fmt.Println(ans)
}
