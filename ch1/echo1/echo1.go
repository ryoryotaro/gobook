package main

// Echo1 は、そのコマンドライン引数を表示します
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sep string
	for i := 0; i < len(os.Args); i++ {
		//s += sep + strconv.Itoa(i) + "," + os.Args[i]
		sep = " "
		s := sep + strconv.Itoa(i) + "," + os.Args[i]
		fmt.Println(s)
	}
	//fmt.Println(s)
}
