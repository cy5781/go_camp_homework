package main

import "fmt"

func main() {
	Slice()
}

func Slice() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := []string{"Jason", "Max", "Marvin", "Jerry", "Sally"}
	fmt.Println(RemoveSliceByIndexGeneric(s1, 0))
	fmt.Println(RemoveSliceByIndexGeneric(s2, 100))

}

// RemoveSliceByIndexGeneric 使用泛型 generic
func RemoveSliceByIndexGeneric[T any](src []T, index int) ([]T, T, string) {

	if index < 0 || index >= len(src) {
		var zero T
		var errorMessage = "index out of range !"
		return nil, zero, errorMessage
	}

	copy(src[:index], src[index+1:])
	res := src[index]
	//去掉最后一个重复元素
	src = Shrink(src)
	return src[:len(src)-1], res, "nil"
}

func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	return c, false
}

func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
