package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type University struct {
	Name    string   `json:"name"`
	Domains []string `json:"domains"`
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/client.go github.com/mattyhall/cb-advance-go-materials/mocking/03/client Client
type Client interface {
	GetUniversities() ([]University, error)
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
