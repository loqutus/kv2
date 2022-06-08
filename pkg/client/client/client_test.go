package client

import (
	"testing"

	"github.com/rusik69/kv2/pkg/client/argparse"
)

func TestClient(t *testing.T) {
	argparse.ArgsInstance.Parse()
	ClientInstance.Init(argparse.ArgsInstance)
	ClientInstance.Connect()
	t.Run("set", func(t *testing.T) {
		err := ClientInstance.Set("key", "value")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("get", func(t *testing.T) {
		value, err := ClientInstance.Get("key")
		if err != nil {
			t.Error(err)
		}
		if value != "value" {
			t.Error("Expected value to be 'value', got:", value)
		}
	})
	t.Run("info", func(t *testing.T) {
		info, err := ClientInstance.Info()
		if err != nil {
			t.Error(err)
		}
		if info != "keys: 1" {
			t.Error("Expected info to be 'keys: 1', got:", info)
		}
	})
}

func BenchmarkClient(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ClientInstance.Set("key", "value")
	}
}
