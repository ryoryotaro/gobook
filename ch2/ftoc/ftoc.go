// ftocは華氏から摂氏への変換を2つ表示します
package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g℉ = %g℃\n", freezingF, fToC(freezingF))
	fmt.Printf("%g℉ = %g℃\n", boilingF, fToC(boilingF))

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}