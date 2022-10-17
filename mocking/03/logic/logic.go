package logic

import (
	"fmt"
	"strings"

	"github.com/mattyhall/cb-advance-go-materials/mocking/03/client"
)

func CountUoLUnis(cli client.Client) (int, error) {
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
