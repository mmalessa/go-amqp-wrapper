package main

import (
	"flag"
	"fmt"
	"log"
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

	// FIXME - add channel to nice shutt down
	for {
	}
	log.Printf("shutting down")

	if err := c.Shutdown(); err != nil {
		log.Fatalf("error during shutdown: %s", err)
	}
}
