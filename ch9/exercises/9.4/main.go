package main

func pipeline(stages int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < stages; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				//将in中的数字读到out中去,如此循环好多次
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}

func main() {
	in, out := pipeline(1000000)
	for i := 0; i < 1000000; i++ {
		in <- 1
		<-out
	}
	close(in)
}
