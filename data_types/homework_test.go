package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	ptr := unsafe.Pointer(&number)

	var res uint32

	for i := 0; i < int(unsafe.Sizeof(number)); i++ {
		// number = 0xFF00AA11
		// ptr = 0xAA
		// res = 0x11

		// Берем 1-й байт значения указателя
		// tmp = 0xAA
		tmp := *(*uint8)(ptr)
		// Сдвигаем res влево на 1 байт
		// res = 0x1100
		res = res << 8
		// Добавляем 1-й байт значения указателя к res
		// res = 0x11AA
		res += uint32(tmp)
		// Сдвигаем указатель на следующий байт
		//ptr = 0x00
		ptr = unsafe.Add(ptr, 1)
	}
	return res
}

type Bitwise interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func ToLittleEndianGeneric[T Bitwise](value T) T {
	ptr := unsafe.Pointer(&value)

	var res T

	for i := 0; i < int(unsafe.Sizeof(value)); i++ {
		tmp := *(*uint8)(ptr)
		res = res << 8
		res += T(tmp)
		ptr = unsafe.Add(ptr, 1)
	}
	return res
}

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndianGeneric(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
