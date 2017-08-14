package slices

import "fmt"

func SliceDemo() {
	ints := make([]int, 0)
	ints = append(ints, 0, 1, 2, 3)
	fmt.Println(ints)

	newInts1 := ints[1:3]
	newInts2 := ints[2:]
	newInts3 := ints[:3]

	fmt.Println(newInts1)
	fmt.Println(newInts2)
	fmt.Println(newInts3)

	copy(ints[2:3], ints[3:4])
	fmt.Println(ints)

}
