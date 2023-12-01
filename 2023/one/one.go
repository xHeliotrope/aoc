package xHeliotrope

import (
	"fmt"

	Util "github.com/xHeliotrope/2023/util"
)

func Run() {
	output := Util.ReadFile("./one/one.txt")
	fmt.Println(output)
}
