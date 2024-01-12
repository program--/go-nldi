package nldi

import (
	"encoding/json"
	"errors"

	"github.com/program--/nldi/internal"
)

const NLDI_ENDPOINT = "https://labs.waterdata.usgs.gov/api/nldi"

type NLDIClient struct {
	data_sources []NLDIDataSource
}

func NewNLDIClient() *NLDIClient {
	client := new(NLDIClient)
	client.data_sources = nil
	return client
}

func (c *NLDIClient) GetDataSources() ([]NLDIDataSource, error) {
	if c.data_sources == nil {

		resp, err := internal.NLDIHttpClient.Get(NLDI_ENDPOINT + "/linked-data")
		if err != nil {
			return nil, err
		}

		reader := json.NewDecoder(resp.Body)
		err = reader.Decode(&c.data_sources)
		if err != nil {
			return nil, err
		}
	}

	return c.data_sources, nil
}

func (c *NLDIClient) GetDataSource(source string) (*NLDIDataSource, error) {
	if c.data_sources == nil {
		if _, err := c.GetDataSources(); err != nil {
			return nil, err
		}
	}

	for _, src := range c.data_sources {
		if src.Source == source || src.SourceName == source {
			return &src, nil
		}
	}

	return nil, errors.New("Invalid NLDI Source `" + source + "`")
}
