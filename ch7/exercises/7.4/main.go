package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type stringReader struct {
	s string
}

//func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
//	b.lastRead = opInvalid
//	for {
//		i := b.grow(MinRead)
//		b.buf = b.buf[:i]
//		m, e := r.Read(b.buf[i:cap(b.buf)])
//		if m < 0 {
//			panic(errNegativeRead)
//		}
//
//		b.buf = b.buf[:i+m]
//		n += int64(m)
//		if e == io.EOF {
//			return n, nil // e is EOF, so return nil explicitly
//		}
//		if e != nil {
//			return n, e
//		}
//	}
//}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s) //为什么copy的方向是这样的？可以看看上面ReadFrom的实现
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}

func TestNewReader() {
	s := "hi there"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		fmt.Errorf("n=%d err=%s", n, err)
		//t.Fail()
	}
	if b.String() != s {
		fmt.Errorf(`"%s" != "%s"`, b.String(), s)
	}
}

func TestNewReaderWithHTML() {
	s := "<html><body><p>hi</p></body></html>"
	_, err := html.Parse(NewReader(s))
	if err != nil {
		fmt.Errorf("Parse failed!")
	}
}


func main() {
	TestNewReader()
}
