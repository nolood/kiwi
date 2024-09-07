package profile

import (
	"strings"
)

func processLocationName(str string, matches []*struct {
	Length int
	Start  int
}) string {

	result := make([]string, 0)

	for _, match := range matches {
		start := match.Start

		name := []byte{}

		if start == 0 || str[start-1] == ',' {
			for i := start; i < len(str); i++ {
				if str[i] == ',' {
					break
				}

				name = append(name, str[i])
			}
		}

		result = append(result, string(name))
	}

	return strings.Join(result, ", ")
}
