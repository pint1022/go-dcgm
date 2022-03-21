package main

import (
	"github.com/pint1022/go-dcgm/pkg/dcgm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// res: curl localhost:8070/dcgm/device/info/id/0

func main() {
	stopSig := make(chan os.Signal, 1)
	signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM)

	cleanup, err := dcgm.Init(dcgm.Embedded)
	if err != nil {
		log.Panicln(err)
	}
	defer cleanup()

	addr := ":8070"
	server := newHttpServer(addr)

	go func() {
		log.Printf("Running http server on localhost%s", addr)
		server.serve()
	}()
	defer server.stop()

	<-stopSig
	return
}
