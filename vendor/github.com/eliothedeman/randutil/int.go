package randutil

import "math/rand"

// IntRange returns a random integer in the range from min to max.
func IntRange(min, size int) int {
	return (rand.Int() % size) + min
}

func Int8() int8 {
	return int8(IntRange(-128, 1<<8))
}

func Uint8() uint8 {
	return uint8(IntRange(0, 1<<8))
}

func Int16() int16 {
	return int16(IntRange(-32768, 1<<16))

}

func Uint16() uint16 {
	return uint16(IntRange(0, 1<<16))
}

func Int32() int32 {
	return int32(IntRange(-2147483648, 1<<32))
}

func Uint32() uint32 {
	return uint32(IntRange(0, 1<<32))
}

func Int64() int64 {
	return int64(Int())
}

func Uint64() uint64 {
	return uint64(Int())
}

func Int() int {
	return rand.Int()
}

func Uint() uint {
	return uint(Int())
}
