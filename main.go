package main

import (
	"flag"
	"crypto/tls"
	"crypto/x509"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
)

var (
	uri          = flag.String("uri", "amqps://testuser:testuser@localhost:5671/test", "AMQP URI")
// 	uri          = flag.String("uri", "amqp://testuser:testuser@localhost:5672/test", "AMQP URI")
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

    // make connection
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
	log.Printf("conn: %v, err: %v", conn, err)

    for {}
}
