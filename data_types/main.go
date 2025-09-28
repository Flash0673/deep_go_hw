package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x uint32 = 0xFF00FF00
	ptr := unsafe.Pointer(&x)

	var res uint32 = 0

	for i := 0; i < 4; i++ {
		tmp := *(*uint8)(ptr)
		fmt.Printf("%#X\n", tmp)
		fmt.Printf("OFFSETED: %#X\n", uint32(tmp)<<i*8)
		res = res << 8
		res += uint32(tmp)
		fmt.Printf("RES: %#X\n", res)
		ptr = unsafe.Add(ptr, 1)
	}

	//res = 0xFF
	fmt.Printf("RES: %#X\n", res)
}
