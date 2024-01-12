package nldi

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type NLDIFeature struct {
	feature geojson.Feature
}

func (f *NLDIFeature) Geometry() orb.Geometry {
	return f.feature.Geometry
}

func (f *NLDIFeature) Properties() geojson.Properties {
	return f.feature.Properties
}

func (f *NLDIFeature) BBox() geojson.BBox {
	return f.feature.BBox
}

func (f *NLDIFeature) Basin() {

}
