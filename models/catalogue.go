package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Catalogue struct {
	url     string
	Entries []Entry
}

func NewCatalogue(url string) (c *Catalogue, e error) {
	c = new(Catalogue)
	c.url = url

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(body, &c.Entries)
	}

	return c, nil
}

func GetEntryFromCatalogue(name string, c *Catalogue) (e Entry) {
	for _, entry := range c.Entries {
		if entry.Name == name {
			e = entry
			break
		}
	}

	return e
}
