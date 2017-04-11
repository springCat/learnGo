package ifFor

import "fmt"

func If() {
	x := 0
	if n := "abc"; x > 0 {
		println(n[2])
	} else if x < 0 {
		println(n[1])
	} else {
		println(n[0])
	}

}

func For() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Printf("index:%d,value:%d", i, v)
		fmt.Println()
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("index:%d,value:%d", i, s[i])
		fmt.Println()
	}
}

func Switch(x int) {
	switch x {
	case 10:
		println("Switch：a")
		fallthrough
	case 0:
		println("Switch：b")

	}
}
