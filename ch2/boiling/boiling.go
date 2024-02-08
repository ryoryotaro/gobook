// boilingは水の沸点を表示します
package main

import (
	"fmt"
)

// このconstはパッケージ全体で見える
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g℉  or %g℃\n", f, c)
}
