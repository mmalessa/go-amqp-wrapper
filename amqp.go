package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

func NewConsumer(cfg Config) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     cfg.Consumer.Tag,
		done:    make(chan error),
	}

	var err error

	log.Printf("dialing %q", cfg.Connection.Uri)
	if strings.HasPrefix(cfg.Connection.Uri, "amqps:") {
		tlsCfg := new(tls.Config)
		tlsCfg.RootCAs = x509.NewCertPool()
		tlsCfg.ServerName = cfg.Connection.ServerName
		if ca, err := ioutil.ReadFile(cfg.Connection.SslCaCert); err == nil {
			tlsCfg.RootCAs.AppendCertsFromPEM(ca)
		}
		if cert, err := tls.LoadX509KeyPair(cfg.Connection.SslCert, cfg.Connection.SslKey); err == nil {
			tlsCfg.Certificates = append(tlsCfg.Certificates, cert)
		}
		c.conn, err = amqp.DialTLS(cfg.Connection.Uri, tlsCfg)
	} else {
		c.conn, err = amqp.Dial(cfg.Connection.Uri)
	}
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	go func() {
		fmt.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%q)", cfg.Exchange.Name)
	if err = c.channel.ExchangeDeclare(
		cfg.Exchange.Name,       // name of the exchange
		cfg.Exchange.Type,       // type
		cfg.Exchange.Durable,    // durable
		cfg.Exchange.AutoDelete, // delete when complete
		cfg.Exchange.Internal,   // internal
		cfg.Exchange.NoWait,     // noWait
		nil,                     // arguments
	); err != nil {
		return nil, fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue %q", cfg.Queue.Name)
	queue, err := c.channel.QueueDeclare(
		cfg.Queue.Name,       // name of the queue
		cfg.Queue.Durable,    // durable
		cfg.Queue.AutoDelete, // delete when unused
		cfg.Queue.Exclusive,  // exclusive
		cfg.Queue.NoWait,     // noWait
		nil,                  // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	// log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
	// 	queue.Name, queue.Messages, queue.Consumers, cfg.Queue.RoutingKeys[0])

	for _, routingKey := range cfg.Queue.RoutingKeys {
		log.Printf("Binding Queue to Exchange (key: %s)", routingKey)
		if err = c.channel.QueueBind(
			cfg.Queue.Name,    // name of the queue
			routingKey,        // bindingKey
			cfg.Exchange.Name, // sourceExchange
			false,             // noWait
			nil,               // arguments
		); err != nil {
			return nil, fmt.Errorf("Queue Bind: %s", err)
		}
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		queue.Name,             // name
		c.tag,                  // consumerTag,
		cfg.Consumer.NoAck,     // noAck
		cfg.Consumer.Exclusive, // exclusive
		cfg.Consumer.NoLocal,   // noLocal
		cfg.Consumer.NoWait,    // noWait
		nil,                    // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	go handle(deliveries, c.done)

	return c, nil
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}
