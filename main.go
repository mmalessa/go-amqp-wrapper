package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var config *Config
var err error
var configFile = flag.String("config", "config.yaml", "Config file")

func init() {
	flag.Parse()
}

func main() {
	config, err = loadConfig(*configFile)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: %v", err))
		return
	}
	c, err := NewConsumer(*config)
	if err != nil {
		log.Fatalf("%s", err)
	}

	stopServerChannel := make(chan os.Signal, 1)
	signal.Notify(stopServerChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-stopServerChannel
	log.Printf("Ask for stop with signal: %T %s\n", sig, sig)

	if err := c.Shutdown(); err != nil {
		log.Fatalf("error during shutdown: %s", err)
	}
}
