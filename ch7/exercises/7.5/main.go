package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Write writes len(p) bytes from p to the underlying data stream
//io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
//并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += n
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}

func TestLimitReader() {
	s := "hi there"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 4)
	n, _ := b.ReadFrom(r)
	if n != 4 {
		fmt.Errorf("n=%d", n)
		//t.Fail()
	}
	if b.String() != "hi t" {
		fmt.Errorf(`"%s" != "%s"`, b.String(), s)
	}
}

func main() {
	TestLimitReader()
}
