//
//  Binary 코딩
//  go-bindata 를 사용하면 바이너리 코드로 바꿔줌
// https://pronist.dev/101
//  by SAiki
//

package main

import (
	"fmt"
)

func main() {
	var test = []byte("\x41\x61\x42\x62")
	fmt.Printf("char:%c\n str:%s\n hex:%x\n type:%T ", test, &test, test, test)
}
