package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rusik69/kv2/pkg/client/client"
)

func Loop(c client.Client) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(c.Host + ":" + c.Port + "> ")
		text, _ := reader.ReadBytes('\n')
		cmd, key, value := client.Parse(text)
		switch strings.ToLower(cmd) {
		case "set":
			if key == "" || len(value) == 0 {
				fmt.Println("Usage: set <key> <value>")
				continue
			}
			err := c.Set(key, value)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("OK")
			}
		case "get":
			if key == "" {
				fmt.Println("Usage: get <key>")
				continue
			}
			value, err := c.Get(key)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(value))
			}
		case "del":
			if key == "" {
				fmt.Println("Usage: del <key>")
				continue
			}
			err := c.Del(key)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("OK")
			}
		case "info":
			info, err := c.Info()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(info)
			}
		case "addnode":
			if key == "" || len(value) == 0 {
				fmt.Println("Usage: addnode <host> <port>")
				continue
			}
			err := c.AddNode(key, string(value))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("OK")
			}
		case "delnode":
			if key == "" || len(value) == 0 {
				fmt.Println("Usage: delnode <host> <port>")
				continue
			}
			err := c.DelNode(key, string(value))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("OK")
			}
		case "setreplicas":
			if key == "" {
				fmt.Println("Usage: setreplicas <replicas>")
				continue
			}
			err := c.SetReplicas(key)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("OK")
			}
		case "id":
			id, err := c.GetID()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(id)
			}
		case "ping":
			err := c.Ping()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("pong")
			}
		case "exit", "quit", "q":
			return
		case "":
			continue
		default:
			fmt.Println("unknown command")
		}
	}
}
