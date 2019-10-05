package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/streadway/amqp"

	"gopkg.in/yaml.v3"
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
	Arguments  amqp.Table
}

type ConfigQueue struct {
	Name        string
	Durable     bool
	AutoDelete  bool
	Exclusive   bool
	NoWait      bool
	RoutingKeys []string
	Arguments   amqp.Table
}

type ConfigConsumer struct {
	Tag           string
	NoAck         bool
	Exclusive     bool
	NoWait        bool
	PrefetchCount int
	PrefetchSize  int
	Global        bool
	Executable    string
}

type Config struct {
	DebugMode  bool
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

	cfg.Exchange.Arguments = castConfigArguments(cfg.Exchange.Arguments)
	cfg.Queue.Arguments = castConfigArguments(cfg.Queue.Arguments)

	// FOR DEBUG ONLY
	// cfgM, _ := json.MarshalIndent(cfg, "", "  ")
	// fmt.Print("CONFIG: ")
	// fmt.Println(string(cfgM))

	return cfg, err
}

// RabbitMQ expects int32 for integer values.
func castConfigArguments(arguments amqp.Table) amqp.Table {
	for k, v := range arguments {
		switch v.(type) {
		case int:
			arguments[k] = int32(v.(int))
		default:
		}
	}
	return arguments
}
