package main

// Indica que so recebe
func receive(ch chan<- bool, value bool) {
	ch <- value
}

// Indica que so envia
func clean(ch <-chan bool) {
	<-ch
}

func main() {
	ch := make(chan bool)
	go receive(ch, true)
	clean(ch)
}
