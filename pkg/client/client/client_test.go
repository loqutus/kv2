package client

import (
	"fmt"
	"testing"

	"github.com/rusik69/kv2/pkg/client/argparse"
)

func TestClient(t *testing.T) {
	argparse.ArgsInstance.Parse()
	ClientInstance.Init(argparse.ArgsInstance)
	ClientInstance.Connect()
	t.Run("set", func(t *testing.T) {
		err := ClientInstance.Set("key", []byte("value"))
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("get", func(t *testing.T) {
		value, err := ClientInstance.Get("key")
		if err != nil {
			t.Error(err)
		}
		expected := "value"
		if len(value) != len(expected) {
			t.Error("Expected value to be 'value', got:", string(value))
		}
		for i := 0; i < len(value); i++ {
			if value[i] != expected[i] {
				t.Error("Expected value to be 'value', got:", string(value))
			}
		}
	})
	t.Run("info", func(t *testing.T) {
		info, err := ClientInstance.Info()
		if err != nil {
			t.Error(err)
		}
	})
}

func BenchmarkClient(b *testing.B) {
	ClientInstance.Connect()
	for n := 0; n < 1000000; n++ {
		err := ClientInstance.Set(fmt.Sprint(n), []byte{byte(n + 1)})
		if err != nil {
			b.Error(err)
		}
	}
	for n := 0; n < 1000000; n++ {
		_, err := ClientInstance.Get(fmt.Sprint(n))
		if err != nil {
			b.Error(err)
		}
	}
	for n := 0; n < 1000000; n++ {
		err := ClientInstance.Del(fmt.Sprint(n))
		if err != nil {
			b.Error(err)
		}
	}
}
