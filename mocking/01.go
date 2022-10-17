package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type University struct {
	Name    string   `json:"name"`
	Domains []string `json:"domains"`
}

type UniversityClient struct {
	country string
}

func (c *UniversityClient) GetUniversities() ([]University, error) {
	res, err := http.Get(fmt.Sprintf("http://universities.hipolabs.com/search?country=%s", c.country))
	if err != nil {
		return nil, fmt.Errorf("could not GET universities API: %w", err)
	}

	defer res.Body.Close()

	var unis []University

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&unis)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	return unis, nil
}

func CountUoLUnis(cli UniversityClient) (int, error) {
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

func main() {
	cli := UniversityClient{ country: "United+Kingdom" }

	count, err := CountUoLUnis(cli)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(count)
}
