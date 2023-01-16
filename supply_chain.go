package openrtb

type SupplyChain struct {
	Complete int                `json:"complete,omitempty"` // Flag indicating whether the chain contains all nodes involved in the transaction leading back to the owner of the site, app or other medium of the inventory, where 0 = no, 1 = yes.
	Nodes    []*SupplyChainNode `json:"nodes,omitempty"`    // Array of SupplyChainNode objects in the order of the chain.
	Ver      string             `json:"ver,omitempty"`      // Version of the supply chain specification in use, in the format of “major.minor”.
	Ext      Extension          `json:"ext,omitempty"`
}

type SupplyChainNode struct {
	Asi    string    `json:"asi,omitempty"`    // The canonical domain name of the SSP, Exchange, Header Wrapper, etc system that bidders connect to.
	Sid    string    `json:"sid,omitempty"`    // The identifier associated with the seller or reseller account within the advertising system.
	Rid    string    `json:"rid,omitempty"`    // The OpenRTB RequestId of the request as issued by this seller.
	Name   string    `json:"name,omitempty"`   // The name of the company (the legal entity) that is paid for inventory transacted under the given seller_ID.
	Domain string    `json:"domain,omitempty"` // The business domain name of the entity represented by this node.
	Hp     int       `json:"hp,omitempty"`     // Indicates whether this node will be involved in the flow of payment for the inventory.
	Ext    Extension `json:"ext,omitempty"`
}
