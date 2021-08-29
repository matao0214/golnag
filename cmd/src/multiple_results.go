package main

import "fmt"

// swap関数を作成
// (x, y)の順番で受け取った引数を(y, x)の引数の形に変換する
// 第一引数で受け取る引数の形とデータ型を定義
// 第二引数で戻り値のデータ型を定義
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a,b)
}
