package main

func main() {
	ch := make(chan int)
	// 双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch //只能写，不能读
	var readCh <-chan int = ch  //只能读，不能写

	writeCh <- 666 //写

	<-readCh //读
}
