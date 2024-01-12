package nldi

import (
	"io"

	"github.com/paulmach/orb/geojson"
	"github.com/program--/nldi/internal"
)

type NLDIDataSource struct {
	Source     string `json:"source"`
	SourceName string `json:"sourceName"`
	Features   string `json:"features"`

	features []*NLDIFeature
}

func (d *NLDIDataSource) GetFeatures() ([]*NLDIFeature, error) {
	if d.features == nil {
		println(NLDI_ENDPOINT + "/linked-data/" + d.Source)
		resp, err := internal.NLDIHttpClient.Get(NLDI_ENDPOINT + "/linked-data/" + d.Source)
		if err != nil {
			return nil, err
		}

		var collection *geojson.FeatureCollection

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		//FIXME: Segfault here

		collection, err = geojson.UnmarshalFeatureCollection(data)
		if err != nil {
			return nil, err
		}

		var nldiFeature *NLDIFeature
		d.features = make([]*NLDIFeature, len(collection.Features))
		for _, feature := range collection.Features {
			nldiFeature = new(NLDIFeature)
			nldiFeature.feature = *feature
			d.features = append(d.features, nldiFeature)
		}
	}

	return d.features, nil
}
