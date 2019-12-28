package object

import (
	"fmt"
	"time"
)

type IntConv func(op int) int //自定义类型

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
