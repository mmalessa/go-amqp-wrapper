package main

/*
* Inspired by https://github.com/corvus-ch/rabbitmq-cli-consumer
 */

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "config, c",
		Usage: "Location of configuration file",
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
	fmt.Println("It working...")
	return nil
}

// ExitErrHandler is a global error handler registered with the application.
func ExitErrHandler(_ *cli.Context, err error) {
	if err == nil {
		return
	}

	code := 1

	// if err.Error() != "" {
	// 	stdlog.Printf("%+v\n", err)
	// }

	if exitErr, ok := err.(cli.ExitCoder); ok {
		code = exitErr.ExitCode()
	}

	os.Exit(code)
}
