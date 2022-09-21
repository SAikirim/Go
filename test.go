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

	s := make([]int, 5, 6)
	s = []int{1, 2, 3, 4}
	fmt.Println("s\n")
}
