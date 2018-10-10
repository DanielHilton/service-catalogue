package catalogue

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Cache struct {
	url     string
	Entries []Entry
}

func NewCache(url string) (c *Cache, e error) {
	c = new(Cache)
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
