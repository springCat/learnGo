package variable

import "testing"

func Test_const(t *testing.T) {
	t.Logf("Test_const :%s", a)
}

func Test_iota(t *testing.T) {
	t.Logf("Test_iota :%d,%d,%d", Monday, Tuesday, Wednesday)
}

func Test_Va(t *testing.T) {
	a := Va()
	t.Logf("Test_Va :%d", a)
}

func Test_Vsi(t *testing.T) {
	s, i := Vsi()
	t.Logf("Test_Vsi s:%s,i:%d", s, i)
}

func Test_Vdatai(t *testing.T) {
	data, i := Vdatai()
	t.Logf("Test_Vdatai data:%v,i:%d", data, i)
}

func Test_Vdatai2(t *testing.T) {
	_, i := Vdatai()
	t.Logf("Test_Vdatai2 ,i:%d", i)
}

func Test_Type(t *testing.T) {
	c := Black
	a := func(colour Colour) {
		t.Logf("Test_Type %v", colour)
	}
	a(c)

	//x := 1
	//a(x)

}
