package kitsu

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/Jeffail/gabs"
)

// Producers struct with all the producers data from kitsu
type Producers struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		Slug      string `json:"slug"`
		Name      string `json:"name"`
	} `json:"attributes"`
	Relationships struct {
		AnimeProductions struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"animeProductions"`
	} `json:"relationships"`
}

// SearchProducer search for a producer on kitsu.io
// query being the producer to search for
func SearchProducer(query string) (*Producers, error) {
	newQuery := url.QueryEscape(query)
	uri := fmt.Sprintf("%s/producers?filter[slug]=%s", baseURL, newQuery)
	byt, er := get(uri)
	if er != nil {
		return nil, er
	}
	parJSON, e := gabs.ParseJSON(byt)
	if e != nil {
		return nil, e
	}
	producer := parJSON.Path("data").Data().([]interface{})
	if len(producer) <= 0 {
		err := fmt.Errorf("Could not find any producers with the query: %s", query)
		return nil, err
	}
	resJSON, er := json.Marshal(producer[0])
	if er != nil {
		return nil, er
	}
	pro := new(Producers)
	err := json.Unmarshal(resJSON, &pro)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
