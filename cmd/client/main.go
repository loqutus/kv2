package main

import (
	"fmt"
	"os"

	"github.com/rusik69/kv2/pkg/client/argparse"
	"github.com/rusik69/kv2/pkg/client/client"
	"github.com/rusik69/kv2/pkg/client/interactive"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetReportCaller(true)
	argparse.ArgsInstance.Parse()
	var c client.Client
	c.Init(argparse.ArgsInstance)
	err := c.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	interactive.Loop(c)
}
