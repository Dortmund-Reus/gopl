package main

import "fmt"

func Parse(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	//panic(bailout{})
	panic(fmt.Sprintf("invalid suit %q", input))
	// ...parser...
}

func weird() (ret string) {
	defer func() {
		recover()
		ret = "hi"
	}()
	panic("omg")
}

func main() {
	err := Parse("hello")
	fmt.Println(err)
	fmt.Println(weird())
}
