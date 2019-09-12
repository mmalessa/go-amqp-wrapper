package main

/*
"crypto/tls"
"crypto/x509"
"flag"
"io/ioutil"
"log"
"github.com/streadway/amqp"
*/
import (
	"flag"
	"fmt"
	"log"
	"os"
)

var config *Config
var err error

var (
	uri = flag.String("uri", "amqps://testuser:testuser@localhost:5671/test", "AMQP URI")
	// uri          = flag.String("uri", "amqp://testuser:testuser@localhost:5672/test", "AMQP URI")
	exchange     = flag.String("exchange", "test_exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "topic", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "test_queue", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "#", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
)

func init() {
	flag.Parse()
}

func main() {

	configFile := "config.yaml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	config, err = loadConfig(configFile)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: %v", err))
		return
	}

	connection, err := getConnection(config.Connection)
	log.Printf("conn: %v, err: %v", connection, err)

	// channel, err := conn.Channel()
	// if err != nil {
	// 	//fmt.Errorf("Channel: %s", err)
	// }

	// if err = channel.ExchangeDeclare(
	// 	exchange,     // name of the exchange
	// 	exchangeType, // type
	// 	false,        // durable
	// 	false,        // delete when complete
	// 	false,        // internal
	// 	false,        // noWait
	// 	nil,          // arguments
	// ); err != nil {
	// 	//return nil, fmt.Errorf("Exchange Declare: %s", err)
	// }

	// for {
	// }
}
