package variable

//1
const (
	Monday  = iota
	a       = "abc"
	Tuesday = iota
	Wednesday
)

//2
func Va() int {
	a := 1
	return a
}

//3
func Vsi() (string, int) {
	var s, i = "string", 123
	return s, i
}

//4
func Vdatai() ([3]int, int) {
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100
	return data, i
}

//5
type Colour int

const (
	Black Colour = iota
	Red
	Blue
)