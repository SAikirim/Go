//
// test 코딩
//
// https://github.com/SAikirim/Go
// by SAiki
//

package main

import (
	"fmt"
)

func main() {
	var test = []byte("\x41\x61\x42\x62")
	fmt.Printf("char:%c\n str:%s\n hex:%x\n type:%T ", test, &test, test, test)
}
