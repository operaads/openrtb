package openrtb

type UserAgent struct {
	Browsers     []*BrandVersion `json:"browsers,omitempty"`     // Each BrandVersion object identifies a browser or similar software component.
	Platform     *BrandVersion   `json:"platform,omitempty"`     // A BrandVersion object that identifies the user agent’s execution platform / OS.
	Mobile       int             `json:"mobile,omitempty"`       // 1 if the agent prefers a “mobile” version of the content, if available, i.e. optimized for small screens or touch input. 0 if the agent prefers the “desktop”
	Architecture string          `json:"architecture,omitempty"` // Device’s major binary architecture, e.g. “x86” or “arm”.
	Bitness      string          `json:"bitness,omitempty"`      // Device’s bitness, e.g. “64” for 64-bit architecture.
	Model        string          `json:"model,omitempty"`        // Device model.
	Source       int             `json:"source,omitempty"`       // The source of data used to create this object.
}

type BrandVersion struct {
	Brand   string    `json:"brand,omitempty"`   // A brand identifier, for example, “Chrome” or “Windows”.
	Version string    `json:"version,omitempty"` // g A sequence of version components, in descending hierarchical order (major, minor, micro, …).
	Ext     Extension `json:"ext,omitempty"`
}
