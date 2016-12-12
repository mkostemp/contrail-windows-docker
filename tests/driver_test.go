package driver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Network Driver test suite")
}

var _ = Describe("Network Driver", func() {

	Context("initially", func() {
		It("listens on a pipe", func() {})
		It("tries to connect to existing HNS switch", func() {})
		It("if HNS switch doesn't exist, creates a new one", func() {})
	})

	Context("on shutdown", func() {
		It("deletes HNS switch", func() {})
	})

	Context("on request from docker daemon", func() {

		Context("on GetCapabilities", func() {
			It("returns local scope CapabilitiesResponse", func() {})
		})

		Context("on CreateNetwork", func() {

			Context("tenant or subnet don't exist in Contrail", func() {
				It("responds with error", func() {})
			})

			Context("tenant and subnet exist in Contrail", func() {
				It("responds with nil", func() {})
			})
		})

		Context("on AllocateNetwork", func() {
			// Dunno. in standard windows drivers, it throws NotImplementedError
		})

		Context("on DeleteNetwork", func() {
			It("responds with nil", func() {})
		})

		Context("on FreeNetwork", func() {
			// Dunno. in standard windows drivers, it throws NotImplementedError
		})

		Context("on CreateEndpoint", func() {
			It("allocates Contrail resources", func() {})
			It("configures container network via HNS", func() {})
			It("configures vRouter agent", func() {})
			It("responds with proper CreateEndpointResponse", func() {})
		})

		Context("on DeleteEndpoint", func() {
			It("deallocates Contrail resources", func() {})
			It("configures vRouter Agent", func() {})
			It("reponds with nil", func() {})
		})

		Context("on EndpointInfo", func() {
			It("responds with proper InfoResponse", func() {})
		})

		Context("on Join", func() {
			Context("queried endpoint exists", func() {
				It("responds with proper JoinResponse", func() {}) // nil maybe?
			})

			Context("queried endpoint doesn't exist", func() {
				It("responds with err", func() {})
			})
			// Dunno. in standard windows drivers, it's an empty func otherwise
		})

		Context("on Leave", func() {
			Context("queried endpoint exists", func() {
				It("responds with proper JoinResponse", func() {}) // nil maybe?
			})

			Context("queried endpoint doesn't exist", func() {
				It("responds with err", func() {})
			})
			// Dunno. in standard windows drivers, it's an empty func otherwise
		})

		Context("on DiscoverNew", func() {
			// Dunno. in standard windows drivers, it's an empty func
		})

		Context("on DiscoverDelete", func() {
			// Dunno. in standard windows drivers, it's an empty func
		})

		Context("on ProgramExternalConnectivity", func() {
			// Dunno. in standard windows drivers, it's an empty func
		})

		Context("on RevokeExternalConnectivity", func() {
			// Dunno. in standard windows drivers, it's an empty func
		})

	})

})
