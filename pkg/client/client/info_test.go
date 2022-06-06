package client

import (
	"testing"

	"github.com/rusik69/kv2/pkg/client/argparse"
)

func TestInfo(t *testing.T) {
	argparse.ArgsInstance.Parse()
	ClientInstance.Init(argparse.ArgsInstance)
	ClientInstance.Connect()
	info, err := ClientInstance.Info()
	if err != nil {
		t.Error(err)
	}
	if info != "keys: 1" {
		t.Error("Expected value to be 'keys: 1', got:", info)
	}
}
