package request

import "github.com/bsm/openrtb"

type EventTracker struct {
	EventType int               `json:"event,omitempty"`
	Method    []int             `json:"methods,omitempty"`
	Ext       openrtb.Extension `json:"ext,omitempty"`
}
