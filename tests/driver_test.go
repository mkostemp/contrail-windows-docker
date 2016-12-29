package test

import (
	"testing"

	"github.com/codilime/contrail-windows-docker/driver"
	"github.com/docker/go-plugins-helpers/network"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Network Driver test suite")
}

var _ = Describe("Network Driver", func() {

	contrail_driver := &driver.ContrailDriver{}

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
			It("returns local scope CapabilitiesResponse, nil", func() {
				resp, err := contrail_driver.GetCapabilities()
				Expect(resp).To(Equal(&network.CapabilitiesResponse{Scope: "local"}))
				Expect(err).ShouldNot(HaveOccurred())
			})
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
			It("responds with empty AllocateNetworkResponse, nil", func() {
				req := network.AllocateNetworkRequest{}
				resp, err := contrail_driver.AllocateNetwork(&req)
				Expect(resp).To(Equal(&network.AllocateNetworkResponse{}))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("on DeleteNetwork", func() {
			It("responds with nil", func() {
				req := network.DeleteNetworkRequest{}
				err := contrail_driver.DeleteNetwork(&req)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("on FreeNetwork", func() {
			It("responds with nil", func() {
				req := network.FreeNetworkRequest{}
				err := contrail_driver.FreeNetwork(&req)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("on CreateEndpoint", func() {
			It("allocates Contrail resources", func() {})
			It("configures container network via HNS", func() {})
			It("configures vRouter agent", func() {})
			It("responds with proper CreateEndpointResponse, nil", func() {
				req := network.CreateEndpointRequest{}
				resp, err := contrail_driver.CreateEndpoint(&req)
				Expect(resp).To(Equal(&network.CreateEndpointResponse{}))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("on DeleteEndpoint", func() {
			It("deallocates Contrail resources", func() {})
			It("configures vRouter Agent", func() {})
			It("responds with nil", func() {
				req := network.DeleteEndpointRequest{}
				err := contrail_driver.DeleteEndpoint(&req)
				Expect(err).ShouldNot(HaveOccurred())
			})
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
		})

		Context("on Leave", func() {

			Context("queried endpoint exists", func() {
				It("responds with proper JoinResponse, nil", func() {}) // nil maybe?
			})

			Context("queried endpoint doesn't exist", func() {
				It("responds with err", func() {})
			})
		})

		Context("on DiscoverNew", func() {
			It("responds with nil", func() {
				req := network.DiscoveryNotification{}
				err := contrail_driver.DiscoverNew(&req)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("on DiscoverDelete", func() {
			It("responds with nil", func() {
				req := network.DiscoveryNotification{}
				err := contrail_driver.DiscoverDelete(&req)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("on ProgramExternalConnectivity", func() {
			It("responds with nil", func() {
				req := network.ProgramExternalConnectivityRequest{}
				err := contrail_driver.ProgramExternalConnectivity(&req)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("on RevokeExternalConnectivity", func() {
			It("responds with nil", func() {
				req := network.RevokeExternalConnectivityRequest{}
				err := contrail_driver.RevokeExternalConnectivity(&req)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

	})

	Describe("allocating resources in Contrail Controller", func() {

		Context("given correct tenant and subnet id", func() {
			It("works", func() {})
		})

		Context("given incorrect tenant and subnet id", func() {
			It("returns proper error message", func() {})
		})

	})

})
