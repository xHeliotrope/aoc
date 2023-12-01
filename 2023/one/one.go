package xHeliotrope

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"

	Util "github.com/xHeliotrope/aoc/2023/util"
)

func Run() {
	var values []int
	for _, v := range Util.ReadFile("./one/one.txt") {
		var firstInteger int = 0
		var lastInteger int = 0
		var construction string = ""
		results := make(map[string]string)
		results["one"] = "1"
		results["two"] = "2"
		results["three"] = "3"
		results["four"] = "4"
		results["five"] = "5"
		results["six"] = "6"
		results["seven"] = "7"
		results["eight"] = "8"
		results["nine"] = "9"
		for _, c := range v {
			if unicode.IsLetter(c) {
				construction = construction + string(c)
				//fmt.Println(construction)
				for key, value := range results {
					if bytes.HasSuffix([]byte(construction), []byte(key)) {
						if firstInteger == 0 {
							firstInteger, _ = strconv.Atoi(string(value))
							lastInteger, _ = strconv.Atoi(string(value))
						} else {
							lastInteger, _ = strconv.Atoi(string(value))
						}
					}
				}
			}
			if unicode.IsDigit(c) {
				construction = ""
				if firstInteger == 0 {
					firstInteger, _ = strconv.Atoi(string(c))
					lastInteger, _ = strconv.Atoi(string(c))
				} else {
					lastInteger, _ = strconv.Atoi(string(c))
				}
			}
		}
		var together int = firstInteger*10 + lastInteger
		//fmt.Printf("%d: %d\n", index, together)
		values = append(values, together)
	}
	fmt.Println(Util.Sum(values))
}
