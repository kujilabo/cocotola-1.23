package main

import (
	"errors"
	"fmt"
)

// 型パラメータがanyとなるジェネリクス
// anyはinterface{}と同義なので型引数としてどの型でも代入できます
func Display[T any](nums []T) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

type data struct {
	Name string
}

func Tx1[T any](a func() (T, error)) (T, error) {
	return a()
}

type Manager[T any] interface {
	Tx1(a func() (T, error)) (T, error)
}

func main() {
	ints := []int{0, 1, 2}
	floats := []float64{10.0, 11.0, 12.0}
	strs := []string{"a", "b", "c"}

	// それぞれの型引数を代入することで関数を実行します
	Display[int](ints)
	Display[float64](floats)
	Display[string](strs)

	fn1 := func() (int, error) {
		return 0, errors.New("xxx")
	}
	fmt.Println(Tx1(fn1))

	fn2 := func() (*data, error) {
		return nil, errors.New("xxx")
	}
	fmt.Println(Tx1(fn2))
}
