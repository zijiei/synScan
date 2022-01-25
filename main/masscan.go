package main

import (
	"Scanner/core"
	"fmt"
	"time"
)

func main() {
	StratTimestap := int32(time.Now().Unix())
	core.Execute()
	time.Sleep(1 * time.Second)
	fmt.Printf("用时：%d 秒\n", int32(time.Now().Unix())-StratTimestap)
}
