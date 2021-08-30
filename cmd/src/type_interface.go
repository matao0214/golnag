package main

import "fmt"

// := でデータ型を指定しない場合、値を見て推論したデータ型が定義される
func main() {
	v := 0.1 + "" // change me!
	fmt.Printf("v is of type %T\n", v)
}
