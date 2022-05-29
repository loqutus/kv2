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
		fmt.Println(client.ClientInstance.Host + ":" + client.ClientInstance.Port + "> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		textSplit := strings.Split(text, " ")
		command := textSplit[0]
		switch strings.ToLower(command) {
		case "set":
			if len(textSplit) < 3 {
				fmt.Println("Usage: set <key> <value>")
				continue
			}
			key := textSplit[1]
			value := strings.Join(textSplit[2:], " ")
			client.ClientInstance.Set(key, value)
		case "get":
			if len(textSplit) < 2 {
				fmt.Println("Usage: get <key>")
				continue
			}
			key := textSplit[1]
			value, err := client.ClientInstance.Get(key)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(value)
		case "exit":
			return
		}
	}
}
