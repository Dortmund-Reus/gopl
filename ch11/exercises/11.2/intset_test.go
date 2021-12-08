package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct{
		input int
		want string
	}{
		{1, "{1}"},
		{2, "{1 2}"},
		{6, "{1 2 6}"},
		{1, "{1 2 6}"},
	}

	var x IntSet
	for _, test := range tests {
		x.Add(test.input)
		if got := x.String(); got != test.want {
			t.Errorf("After Add(%d), x = %v", test.input, got)
		}
	}
}

func TestHas(t *testing.T) {
	tests := []struct{
		input int
		want bool
	}{
		{1, true},
		{2, true},
		{6, true},
		{99, false},
	}

	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(22)
	x.Add(6)
	for _, test := range tests {
		//fmt.Println(x.String())
		if got := x.Has(test.input); got != test.want {
			t.Errorf("Has(%d) = %v", test.input, got)
		}
	}
}

func TestUnion(t *testing.T) {

	//tests := []struct{
	//	input int
	//	want s
	//}{
	//	{1, true},
	//}

	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(22)
	x.Add(6)

	y.Add(9)
	y.Add(5)
	y.Add(1)
	y.Add(886)
	x.UnionWith(&y)
	want := "{1 2 5 6 9 22 886}"
	if x.String() != want {
		t.Errorf("UnionWith(%q) = %v", y.String(), x.String())
	}
	//for _, test := range tests {
	//	//fmt.Println(x.String())
	//	if got := x.String(); got != test.want {
	//		t.Errorf("Has(%d) = %v", test.input, got)
	//	}
	//}
}
