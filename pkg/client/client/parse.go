package client

func Parse(input []byte) (cmd, key, value string) {
	//fmt.Println("input:", input)
	var i int
	for i < len(input)-1 {
		if input[i] == ' ' {
			break
		}
		cmd += string(input[i])
		i++
	}
	i++
	for i < len(input)-1 {
		if input[i] == ' ' || input[i] == 0 {
			break
		}
		key += string(input[i])
		i++
	}
	i++
	if cmd == "set" {
		for i < len(input)-1 {
			if input[i] == '\n' || input[i] == 0 {
				break
			}
			value += string(input[i])
			i++
		}
	}
	//logrus.Println("len:", len(input), len(cmd), len(key), len(value))
	//logrus.Println("tokens:", cmd, key, value)
	return
}
