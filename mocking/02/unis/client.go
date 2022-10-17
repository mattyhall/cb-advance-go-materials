package unis

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type University struct {
	Name    string   `json:"name"`
	Domains []string `json:"domains"`
}

type Client interface {
	GetUniversities() ([]University, error)
}

type MockClient struct {
	results []University
}

func (c *MockClient) GetUniversities() ([]University, error) {
	return c.results, nil
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
