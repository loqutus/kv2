package client

import (
	"fmt"
	"testing"

	"github.com/rusik69/kv2/pkg/client/argparse"
)

func TestClient(t *testing.T) {
	var c1 Client
	c1.Init(argparse.ArgsInstance)
	c1.Connect()
	t.Run("set", func(t *testing.T) {
		err := c1.Set("key", []byte("value"))
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("get", func(t *testing.T) {
		value, err := c1.Get("key")
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
	t.Run("addnode", func(t *testing.T) {
		err := c1.AddNode("127.0.0.1", "6972")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("delnode", func(t *testing.T) {
		err := c1.DelNode("127.0.0.1", "6972")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("setreplicas", func(t *testing.T) {
		err := c1.SetReplicas("2")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("info", func(t *testing.T) {
		_, err := c1.Info()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("getid", func(t *testing.T) {
		id, err := c1.GetID()
		if err != nil {
			t.Error(err)
		}
	})
}

func BenchmarkClient(b *testing.B) {
	argparse.ArgsInstance.Parse()
	var c Client
	c.Init(argparse.ArgsInstance)
	c.Connect()
	for n := 0; n < 1000000; n++ {
		err := c.Set(fmt.Sprint(n), []byte(fmt.Sprintf("%d", n)))
		if err != nil {
			b.Error(err)
		}
	}
	for n := 0; n < 1000000; n++ {
		_, err := c.Get(fmt.Sprint(n))
		if err != nil {
			b.Error(err)
		}
	}
	for n := 0; n < 1000000; n++ {
		err := c.Del(fmt.Sprint(n))
		if err != nil {
			b.Error(err)
		}
	}
}
