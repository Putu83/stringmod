package main

import (
	"fmt"
	str "github.com/Putu83/stringmod/strings"
	qt "github.com/Putu83/stringmod/quotes"
)

func main() {
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e) // 3 2
	fmt.Println(qt.GetEmoji("turtle"))
}
