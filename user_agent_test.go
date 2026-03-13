package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserAgent", func() {
	var subject *UserAgent

	BeforeEach(func() {
		err := fixture("user_agent", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&UserAgent{
			Browsers: []*BrandVersion{
				{
					Brand:   "Chrome",
					Version: []string{"68.0.2704.79"},
				},
			},
			Platform: &BrandVersion{
				Brand:   "iOS",
				Version: []string{"16.2"},
			},
			Mobile:       1,
			Architecture: "ARM",
			Bitness:      "64",
			Model:        "iPhone 13",
			Source:       3,
		}))
	})
})
