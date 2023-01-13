package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SupplyChain", func() {
	var subject *SupplyChain

	BeforeEach(func() {
		err := fixture("supply_chain", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&SupplyChain{
			Complete: 1,
			Ver:      "1.0",
			Nodes: []SupplyChainNode{
				{
					Asi:    "exchange1.com",
					Sid:    "1234",
					Hp:     1,
					Rid:    "bid-request-1",
					Name:   "publisher",
					Domain: "publisher.com",
				},
			},
			Ext: Extension("{}"),
		}))
	})
})
