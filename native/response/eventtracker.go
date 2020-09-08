package response

import "github.com/bsm/openrtb"

const (
	EventTypeImpression = 1

	EventTrackMethodImg = 1
	EventTrackMethodJs = 2
)

type EventTracker struct {
	EventType int `json:"event,omitempty"`
	Method int `json:"method,omitempty"`
	Url string `json:"url,omitempty"`
	CustomData openrtb.Extension `json:"customdata,omitempty"`
	Ext openrtb.Extension `json:"ext,omitempty"`
}
