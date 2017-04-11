package funcs

import "fmt"

//（1）普通函数
func add(x, y int) (z int) {
	z = x + y
	return
}

//（1）多值返回
func test() (int, int) {
	return 1, 2
}

func lamda() {
	//3）匿名函数
	fn := func() {
		println("Hello, World!")
	}
	fn()

	fn1 := add1(1, 2)
	fn1()

}

//闭包
func add1(x, y int) func() int {
	return func() int {
		return x + y
	}
}

//defer
func Defer() {
	defer func() {
		fmt.Println("Defer 1")
	}()

	defer func() {
		fmt.Println("Defer 2")
	}()

	fmt.Println("Defer index")
}
