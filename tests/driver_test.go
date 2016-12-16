package test

import (
	"testing"
	"time"

	"github.com/codilime/contrail-windows-docker/driver"
	"github.com/docker/go-connections/sockets"
	"github.com/docker/go-plugins-helpers/network"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contrail Network Driver test suite")
}

var _ = Describe("Contrail Network Driver", func() {

	contrailDriver := &driver.ContrailDriver{}

	Context("upon starting", func() {

		It("listens on a pipe", func() {
			d := startDriver()
			defer stopDriver(d)

			Skip("TODO")
			_, err := sockets.DialPipe("//./pipe/"+driver.DriverName, time.Second*3)
			Expect(err).ToNot(HaveOccurred())
		})

		It("tries to connect to existing HNS switch", func() {
			d := startDriver()
			defer stopDriver(d)
		})

		It("if HNS switch doesn't exist, creates a new one", func() {
			net, err := driver.GetHNSNetworkByName(driver.NetworkHNSname)
			Expect(net).To(BeNil())
			Expect(err).To(HaveOccurred())

			d := startDriver()
			defer stopDriver(d)

			net, err = driver.GetHNSNetworkByName(driver.NetworkHNSname)
			Expect(net).ToNot(BeNil())
			Expect(net.Id).To(Equal(d.HnsID))
			Expect(err).ToNot(HaveOccurred())
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

	Context("upon shutting down", func() {
		It("HNS switch isn't removed", func() {})
	})

	Context("on request from docker daemon", func() {

		Context("on GetCapabilities", func() {
			It("returns local scope CapabilitiesResponse, nil", func() {
				resp, err := contrailDriver.GetCapabilities()
				Expect(resp).To(Equal(&network.CapabilitiesResponse{Scope: "local"}))
				Expect(err).ToNot(HaveOccurred())
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
				resp, err := contrailDriver.AllocateNetwork(&req)
				Expect(resp).To(Equal(&network.AllocateNetworkResponse{}))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("on DeleteNetwork", func() {
			It("responds with nil", func() {
				req := network.DeleteNetworkRequest{}
				err := contrailDriver.DeleteNetwork(&req)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("on FreeNetwork", func() {
			It("responds with nil", func() {
				req := network.FreeNetworkRequest{}
				err := contrailDriver.FreeNetwork(&req)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("on CreateEndpoint", func() {
			It("allocates Contrail resources", func() {})
			It("configures container network via HNS", func() {})
			It("configures vRouter agent", func() {})
			It("responds with proper CreateEndpointResponse, nil", func() {
				req := network.CreateEndpointRequest{}
				resp, err := contrailDriver.CreateEndpoint(&req)
				Expect(resp).To(Equal(&network.CreateEndpointResponse{}))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("on DeleteEndpoint", func() {
			It("deallocates Contrail resources", func() {})
			It("configures vRouter Agent", func() {})
			It("responds with nil", func() {
				req := network.DeleteEndpointRequest{}
				err := contrailDriver.DeleteEndpoint(&req)
				Expect(err).ToNot(HaveOccurred())
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
				It("responds with proper JoinResponse, nil", func() {})
			})

			Context("queried endpoint doesn't exist", func() {
				It("responds with err", func() {})
			})
		})

		Context("on DiscoverNew", func() {
			It("responds with nil", func() {
				req := network.DiscoveryNotification{}
				err := contrailDriver.DiscoverNew(&req)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("on DiscoverDelete", func() {
			It("responds with nil", func() {
				req := network.DiscoveryNotification{}
				err := contrailDriver.DiscoverDelete(&req)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("on ProgramExternalConnectivity", func() {
			It("responds with nil", func() {
				req := network.ProgramExternalConnectivityRequest{}
				err := contrailDriver.ProgramExternalConnectivity(&req)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("on RevokeExternalConnectivity", func() {
			It("responds with nil", func() {
				req := network.RevokeExternalConnectivityRequest{}
				err := contrailDriver.RevokeExternalConnectivity(&req)
				Expect(err).ToNot(HaveOccurred())
			})
		})

	})
})

func startDriver() *driver.ContrailDriver {
	d, err := driver.NewDriver()
	Expect(err).ToNot(HaveOccurred())
	Expect(d.HnsID).ToNot(Equal(""))
	return d
}

func stopDriver(d *driver.ContrailDriver) {
	err := d.Teardown()
	Expect(err).ToNot(HaveOccurred())
	net, err := driver.GetHNSNetworkByName(driver.NetworkHNSname)
	Expect(net).To(BeNil())
	Expect(err).To(HaveOccurred())
}
