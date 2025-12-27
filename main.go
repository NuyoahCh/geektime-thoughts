package main

const (
	Status0 = 1 << iota
	Status1
	Status2
	Status3

	Status4 // 不管间隔多少行，都是接着 4

	Status5 = 6 // 插入一个主动赋值，就中断了 iota
	Status6
)

// 递归事例
func recursive() {
	recursive()
}

// A 方法
func A() {
	B()
}

// B 方法
func B() {
	C()
}

// C 方法
func C() {
	A()
}

func main() {

}
