package main

func main() {
	//Tamanho do espaco do canal é 2
	ch := make(chan string, 2)
	ch <- "Teste"
	ch <- "Teste"

	println(<-ch)
	println(<-ch)
}
