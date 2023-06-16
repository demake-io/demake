package demake

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	. "github.com/gospel-dev/gospel"
	"github.com/gospel-dev/gospel/css"
)

func MainMenu(c Context) Element {
	return Ul(
		Li("Spreadsheets"),
		Li("Writer"),
		Li("Tasks"),
	)
}

func MainHeader(c Context) Element {
	return Header(
		Nav(
			Ul(
				Li("Demake"),
				Li("max@mustermann.de"),
			),
		),
	)
}

func Root(c Context) Element {
	return css.CSS(
		Div(
			css.Flex,
			MainHeader(c),
			MainMenu(c),
		),
	)
}

func Run() {
	server := MakeServer(&App{
		Root:         Root,
		StaticPrefix: "/static",
	})

	server.Start()

	wait()
}

func wait() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done // Will block here until user hits ctrl+c
}
