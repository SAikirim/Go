package main

import (
	"fmt"
)

func main() {
	var test = []byte("0x410x61\x42\x62")
	fmt.Printf("char:%c\n str:%s\n hex:%x\n type:%T ", test, &test, test, test)
}
