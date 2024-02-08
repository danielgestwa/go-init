package main

type Number struct {
	value int
}

func (num *Number) AddTo(value int) {
	num.value += value
}

func Add(first, second int) int {
	return first + second
}
