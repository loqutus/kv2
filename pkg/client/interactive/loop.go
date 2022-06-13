package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rusik69/kv2/pkg/client/client"
)

func Loop() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(client.ClientInstance.Host + ":" + client.ClientInstance.Port + "> ")
		text, _ := reader.ReadBytes('\n')
		cmd, key, value := client.Parse(text)
		switch strings.ToLower(cmd) {
		case "set":
			if key == "" || len(value) == 0 {
				fmt.Println("Usage: set <key> <value>")
				continue
			}
			err := client.ClientInstance.Set(key, value)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("OK")
		case "get":
			if key == "" {
				fmt.Println("Usage: get <key>")
				continue
			}
			value, err := client.ClientInstance.Get(key)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(string(value))
		case "del":
			if key == "" {
				fmt.Println("Usage: del <key>")
				continue
			}
			err := client.ClientInstance.Del(key)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("OK")
		case "info":
			info, err := client.ClientInstance.Info()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(info)
		case "addnode":
			if key == "" || len(value) == 0 {
				fmt.Println("Usage: addnode <host> <port>")
				continue
			}
			err := client.ClientInstance.AddNode(key, string(value))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("OK")
		case "delnode":
			if key == "" {
				fmt.Println("Usage: delnode <host>")
				continue
			}
			err := client.ClientInstance.DelNode(key)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("OK")
		case "setreplicas":
			if key == "" {
				fmt.Println("Usage: setreplicas <replicas>")
				continue
			}
			err := client.ClientInstance.SetReplicas(key)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("OK")
		case "exit":
			return
		default:
			fmt.Println("unknown command")
		}
	}
}
