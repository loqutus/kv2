package client

import (
	"fmt"
	"testing"

	"github.com/rusik69/kv2/pkg/client/argparse"
)

func TestSet(t *testing.T) {
	argparse.ArgsInstance.Parse()
	ClientInstance.Init(argparse.ArgsInstance)
	ClientInstance.Connect()
	err := ClientInstance.Set("key", "value")
	if err != nil {
		t.Error(err)
	}
	value, err := ClientInstance.Get("key")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("value:", value)
	fmt.Println("value_len:", len(value))
	if value != "value" {
		t.Error("Expected value to be 'value', got:", value)
	}
}
