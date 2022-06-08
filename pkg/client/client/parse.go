package client

// Parse parses the command string
func Parse(input []byte) (cmd, key, value string) {
	var i int
	for i < len(input)-1 {
		if input[i] == ' ' || input[i] == 0 {
			break
		}
		cmd += string(input[i])
		i++
	}
	i++
	for i < len(input)-1 {
		if input[i] == ' ' || input[i] == 0 || input[i] == '\n' {
			break
		}
		key += string(input[i])
		i++
	}
	i++
	for i < len(input)-1 {
		if input[i] == '\n' || input[i] == 0 {
			break
		}
		value += string(input[i])
		i++
	}
	return
}
