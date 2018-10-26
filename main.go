package main

/*
* Inspired by https://github.com/corvus-ch/rabbitmq-cli-consumer
 */

import (
	"fmt"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/cli"
	"github.com/bketelsen/logr"
	"github.com/corvus-ch/rabbitmq-cli-consumer/config"
	"github.com/corvus-ch/rabbitmq-cli-consumer/log"
	"github.com/corvus-ch/rabbitmq-cli-consumer/command"
	"github.com/corvus-ch/rabbitmq-cli-consumer/consumer"	
	"github.com/corvus-ch/rabbitmq-cli-consumer/acknowledger"
	"github.com/corvus-ch/rabbitmq-cli-consumer/processor"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "Location of configuration file",
	},
	cli.StringFlag{
		Name:  "executable, e",
		Usage: "Location of executable",
	},
}

func main() {
	NewApp().Run(os.Args)
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "go-amqp-wrapper"
	app.Usage = "RabbitMQ consumer-wrapper"
	app.Authors = []cli.Author{
		{"Marcin Malessa", "marcin@malessa.pl"},
	}
	app.Version = "experimental"
	app.Flags = flags
	app.Action = Action
	app.ExitErrHandler = ExitErrHandler
	return app
}

// Action is the function being run when the application gets executed.
func Action(c *cli.Context) error {
	cfg, err := LoadConfiguration(c)
	if err != nil {
		return err
	}

	
	l, infW, errW, err := log.NewFromConfig(cfg)
	if err != nil {
		return err
	}

	// b := CreateBuilder(c.Bool("pipe"), cfg.RabbitMq.Compression, c.Bool("include"))
	b := CreateBuilder(true, false, false)

	// builder, err := command.NewBuilder(b, c.String("executable"), c.Bool("output"), l, infW, errW)
	builder, err := command.NewBuilder(b, c.String("executable"), true, l, infW, errW)

	if err != nil {
		return fmt.Errorf("failed to create command builder: %v", err)
	}

	ack := acknowledger.NewFromConfig(cfg)
	p := processor.New(builder, ack, l)

	client, err := consumer.NewFromConfig(cfg, p, l)
	if err != nil {
		return err
	}
	defer client.Close()

	return consume(client, l)
	return nil
}

func consume(client *consumer.Consumer, l logr.Logger) error {
	fmt.Println("Start consume...")
	done := make(chan error)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		done <- client.Consume(ctx)
	}()

	 select {
	 	case err := <-client.NotifyClose(make(chan error)):
	 		return cli.NewExitError(fmt.Sprintf("connection closed: %v", err), 10)

	 	case <-sig:
	 		l.Info("Cancel consumption of messages.")
	 		cancel()
	 		return checkConsumeError(<-done)

		case err := <-done:
			return checkConsumeError(err)
	}
	return nil
}

func checkConsumeError(err error) error {
	switch err.(type) {
	case *processor.AcknowledgmentError:
		return cli.NewExitError(err, 11)

	default:
		return err
	}
}

// ExitErrHandler is a global error handler registered with the application.
func ExitErrHandler(_ *cli.Context, err error) {
	if err == nil {
		return
	}

	code := 1

	if err.Error() != "" {
		// stdlog.Printf("%+v\n", err)
		fmt.Printf("%+v\n", err)
	}

	if exitErr, ok := err.(cli.ExitCoder); ok {
		code = exitErr.ExitCode()
	}

	os.Exit(code)
}

func CreateBuilder(pipe, compression, metadata bool) command.Builder {
	if pipe {
		return &command.PipeBuilder{}
	}

	return &command.ArgumentBuilder{
		Compressed:   compression,
		WithMetadata: metadata,
	}
}

func LoadConfiguration(c *cli.Context) (*config.Config, error) {
	file := c.String("config")

	if file == "" {
		file = "default.conf"
	}
	fmt.Println("Loading configuration from file: ", file)
	
	cfg, err := config.LoadAndParse(file)

	if err != nil {
		return nil, fmt.Errorf("failed parsing configuration: %s", err)
	}
	return cfg, err
}