package argparse

import "strings"

func parseNodes(nodesString string) [][]string {
	if nodesString == "" {
		return nil
	} else {
		res := [][]string{}
		nodes := strings.Split(nodesString, ",")
		for _, node := range nodes {
			s := strings.Split(node, ":")
			if len(s) != 2 {
				continue
			}
			res = append(res, []string{s[0], s[1]})
		}
		return res
	}
}
