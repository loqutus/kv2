package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/rusik69/kv2/pkg/client/argparse"
)

func TestClient(t *testing.T) {
	var c1 Client
	argparse.ArgsInstance.ServerHost = "kv2"
	argparse.ArgsInstance.ServerPort = "6969"
	c1.Init(argparse.ArgsInstance)
	err := c1.Connect()
	if err != nil {
		t.Error(err)
	}
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
	t.Run("upload", func(t *testing.T) {
		f, err := os.CreateTemp("", "kv2-test")
		if err != nil {
			t.Error(err)
		}
		defer f.Close()
		defer os.Remove(f.Name())
		_, err = f.WriteString("test")
		if err != nil {
			t.Error(err)
		}
		err = c1.Upload(f.Name(), "test")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("download", func(t *testing.T) {
		err := c1.Download("test")
		if err != nil {
			t.Error(err)
		}
		f, err := os.Open("test")
		if err != nil {
			t.Error(err)
		}
		defer f.Close()
		defer os.Remove(f.Name())
		b := make([]byte, 4)
		l, err := f.Read(b)
		if err != nil {
			t.Error(err)
		} else if l != 4 {
			t.Error("Expected 4 bytes, got:", l)
		} else if string(b) != "test" {
			t.Error("Expected file to contain 'test', got:", b)
		}
	})
	// t.Run("addnode", func(t *testing.T) {
	// 	err := c1.AddNode("localhost", "6969")
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// })
	// t.Run("delnode", func(t *testing.T) {
	// 	err := c1.DelNode("localhost", "6969")
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// })
	// t.Run("setreplicas", func(t *testing.T) {
	// 	err := c1.SetReplicas("2")
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// })
	t.Run("info", func(t *testing.T) {
		infoString, err := c1.Info()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(infoString)
	})
	t.Run("getid", func(t *testing.T) {
		_, err := c1.GetID()
		if err != nil {
			t.Error(err)
		}
	})
}

func BenchmarkClient(b *testing.B) {
	var c Client
	c.Init(argparse.ArgsInstance)
	err := c.Connect()
	if err != nil {
		b.Error(err)
	}
	// err = c.AddNode("kv2-1", "6969")
	// if err != nil {
	// 	b.Error(err)
	// }
	// err = c.AddNode("kv2-2", "6969")
	// if err != nil {
	// 	b.Error(err)
	// }
	// err = c.SetReplicas("3")
	// if err != nil {
	// 	b.Error(err)
	// }
	for n := 0; n < 100000; n++ {
		err := c.Set(fmt.Sprint(n), []byte(fmt.Sprintf("%d", n)))
		if err != nil {
			b.Error(err)
		}
	}
	for n := 0; n < 100000; n++ {
		_, err := c.Get(fmt.Sprint(n))
		if err != nil {
			b.Error(err)
		}
	}
	info1, err := c.Info()
	if err != nil {
		b.Error(err)
	}
	fmt.Println(info1)
	// var c2, c3 Client
	// c2.Init(argparse.Args{ServerHost: "127.0.0.1", ServerPort: "6971"})
	// c3.Init(argparse.Args{ServerHost: "127.0.0.1", ServerPort: "6973"})
	// err = c2.Connect()
	// if err != nil {
	// 	b.Error(err)
	// }
	// err = c3.Connect()
	// if err != nil {
	// 	b.Error(err)
	// }
	// info2, err := c2.Info()
	// if err != nil {
	// 	b.Error(err)
	// }
	// fmt.Println(info2)
	// info3, err := c3.Info()
	// if err != nil {
	// 	b.Error(err)
	// }
	// fmt.Println(info3)
}
