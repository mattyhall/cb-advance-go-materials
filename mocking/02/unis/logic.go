package unis

import (
	"fmt"
	"strings"
)

func CountUoLUnis(cli Client) (int, error) {
	unis, err := cli.GetUniversities()
	if err != nil {
		return 0, fmt.Errorf("could not get universities: %w", err)
	}

	count := 0
	for _, uni := range unis {
		if strings.Contains(uni.Name, "University of London") {
			count++
		}
	}

	return count, nil
}
