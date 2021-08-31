package main
import "fmt"
import "strings"
import "strconv"

func main(){
    var Bob, Alice string
    var b, a int
    fmt.Scan(&Bob, &Alice)
    keisan(Bob, &b)
    keisan(Alice, &a)
    if a > b {
        fmt.Println("Alice")
    } else if a < b {
        fmt.Println("Bob")
    } else {
        fmt.Println("Draw")
    }
}

func keisan(score string, value *int) *int{
    s := strings.Split(score, "")
    var x_0 int
    var x_1 int
    x_0, _ = strconv.Atoi(s[0])
    x_1, _ = strconv.Atoi(s[1])
    *value = x_0 + x_1
    if *value >= 10 {
        *value -= 10
    }
    return value
}
