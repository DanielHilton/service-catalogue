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

func NewCache(url string) *Cache {
	c := new(Cache)
	c.url = url

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &c.Entries)
	}

	return c
}
