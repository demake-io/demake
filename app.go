package demake

import (
	"fmt"
	. "github.com/gospel-dev/gospel"
	"os"
	"os/signal"
	"syscall"
)

func Root(c Context) Element {
	return Div("hey")
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
