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
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/streadway/amqp"
)

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

	if strings.HasPrefix(*uri, "amqps:") {
		// SSL connection
		cfg := new(tls.Config)
		cfg.RootCAs = x509.NewCertPool()
		cfg.ServerName = "PmServer"
		if ca, err := ioutil.ReadFile("ssl/cacert.pem"); err == nil {
			cfg.RootCAs.AppendCertsFromPEM(ca)
		}
		if cert, err := tls.LoadX509KeyPair("ssl/cert.pem", "ssl/key.pem"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		}
		conn, err := amqp.DialTLS(*uri, cfg)
		log.Printf("SSL conn: %v, err: %v", conn, err)
	} else {
		// connection
		conn, err := amqp.Dial(*uri)
		if err != nil {
			//fmt.Errorf("Channel: %s", err)
		}
		log.Printf("conn: %v, err: %v", conn, err)
	}

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

	for {
	}
}
