package funcs

import (
		"testing"
		"fmt"
)

func Test_Defer(t *testing.T) {

		fmt.Println(add(1,2))
		fmt.Println("-----add------")
		fmt.Println(swap(1,2))
		fmt.Println("-----swap------")
		lamda()
		fmt.Println("-----lamda------")
		Defer()
		fmt.Println("-----Defer------")
}
