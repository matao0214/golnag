package main

import "fmt"

// 引数を渡したら計算された結果が返ってくるsplit関数
func sprit(sum int) (x, y int) {
	x = sum * 10
	y = sum + 10
	return
}

func main() {
	fmt.Println(sprit(5))
}
