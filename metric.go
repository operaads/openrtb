package openrtb

type Metric struct {
	Type string `json:"type"`
	Value float64 `json:"value,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Ext Extension `json:"ext,omitempty"`
}
