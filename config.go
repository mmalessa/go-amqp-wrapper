package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigConnection struct {
	Uri        string
	ServerName string
	SslCaCert  string
	SslCert    string
	SslKey     string
}
type ConfigExchange struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	// Arguments  []string
}
type ConfigQueue struct {
	Name        string
	Durable     bool
	AutoDelete  bool
	Exclusive   bool
	NoWait      bool
	RoutingKeys []string
}

type ConfigConsumer struct {
	Tag       string
	NoAck     bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
}

type Config struct {
	Connection ConfigConnection
	Exchange   ConfigExchange
	Queue      ConfigQueue
	Consumer   ConfigConsumer
}

func loadConfig(location string) (*Config, error) {

	cfg := &Config{}
	_, err := os.Stat(location)

	if os.IsNotExist(err) {
		return cfg, fmt.Errorf("Config file not found: %s", location)
	}

	yamlFile, err := ioutil.ReadFile(location)
	if err != nil {
		return cfg, fmt.Errorf("Load config error %v ", err)
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return cfg, fmt.Errorf("Unmarshal: %v", err)
	}

	// FOR DEBUG ONLY
	cfgM, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Print("CONFIG: ")
	fmt.Println(string(cfgM))

	return cfg, err
}
