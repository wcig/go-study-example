package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dtm-labs/client/dtmcli/logger"
)

func main() {
	// dtm client logger
	EnableDtmcliLogger()

	// run rm1, rm2
	RunRM1Server()
	RunRM2Server()

	// run ap
	RunAPTx()

	// wait exit
	WaitExit()
}

func EnableDtmcliLogger() {
	logger.InitLog("debug")
}

func WaitExit() {
	// Ctrl+C exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("quit (%v)\n", <-sig)
}
