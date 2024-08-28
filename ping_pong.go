// Importando pacote principal
package main

// importando bibliotecas
import (
	"fmt"
	"time"
)

//Declarando funções
func ping(canal chan string) {
	for i := 0; ; i++ {
		canal <- "ping"
	}
}

func pong(canal chan string) {
	for i := 0; ; i++ {
		canal <- "pong"
	}
}

//Função principal
func saida(canal chan string) {
	for {
		msg := <- canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
	
func main() {
	var canal chan string = make(chan string)

	go ping(canal)
	go saida(canal)
	go pong(canal)

	var entrada string 
	fmt.Scanln(&entrada)
}
