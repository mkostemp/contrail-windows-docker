package hns

import (
	"flag"
	"net"
	"strings"
	"testing"

	log "github.com/Sirupsen/logrus"

	"github.com/Microsoft/hcsshim"
	"github.com/codilime/contrail-windows-docker/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var netAdapter string

func init() {
	flag.StringVar(&netAdapter, "netAdapter", "Ethernet0",
		"Network adapter to connect HNS switch to")
}

func TestHNS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HNS wrapper test suite")
}

var _ = BeforeSuite(func() {
	err := common.HardResetHNS()
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	err := common.HardResetHNS()
	Expect(err).ToNot(HaveOccurred())
})

var _ = Describe("HNS wrapper", func() {

	const (
		tenantName  = "agatka"
		networkName = "test_net"
		subnetCIDR  = "10.0.0.0/24"
		defaultGW   = "10.0.0.1"
	)

	var originalNumNetworks int

	BeforeEach(func() {
		nets, err := ListHNSNetworks()
		Expect(err).ToNot(HaveOccurred())
		originalNumNetworks = len(nets)
	})

	Context("HNS network exists", func() {

		testNetName := "TestNetwork"
		testHnsNetID := ""

		BeforeEach(func() {
			expectNumberOfEndpoints(0)

			Expect(testHnsNetID).To(Equal(""))
			testHnsNetID = MockHNSNetwork(testNetName, netAdapter, subnetCIDR, defaultGW)
			Expect(testHnsNetID).ToNot(Equal(""))

			net, err := GetHNSNetwork(testHnsNetID)
			Expect(err).ToNot(HaveOccurred())
			Expect(net).ToNot(BeNil())
		})

		AfterEach(func() {
			endpoints, err := ListHNSEndpoints()
			Expect(err).ToNot(HaveOccurred())
			if len(endpoints) > 0 {
				// Cleanup lingering endpoints.
				for _, ep := range endpoints {
					err = DeleteHNSEndpoint(ep.Id)
					Expect(err).ToNot(HaveOccurred())
				}
				expectNumberOfEndpoints(0)
			}

			Expect(testHnsNetID).ToNot(Equal(""))
			err = DeleteHNSNetwork(testHnsNetID)
			Expect(err).ToNot(HaveOccurred())
			_, err = GetHNSNetwork(testHnsNetID)
			Expect(err).To(HaveOccurred())
			testHnsNetID = ""
			nets, err := ListHNSNetworks()
			Expect(err).ToNot(HaveOccurred())
			Expect(nets).ToNot(BeNil())
			Expect(len(nets)).To(Equal(originalNumNetworks))
		})

		Specify("listing all HNS networks works", func() {
			nets, err := ListHNSNetworks()
			Expect(err).ToNot(HaveOccurred())
			Expect(nets).ToNot(BeNil())
			Expect(len(nets)).To(Equal(originalNumNetworks + 1))
			found := false
			for _, n := range nets {
				if n.Id == testHnsNetID {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue())
		})

		Specify("getting a single HNS network works", func() {
			net, err := GetHNSNetwork(testHnsNetID)
			Expect(err).ToNot(HaveOccurred())
			Expect(net).ToNot(BeNil())
			Expect(net.Id).To(Equal(testHnsNetID))
		})

		Specify("getting a single HNS network by name works", func() {
			net, err := GetHNSNetworkByName(testNetName)
			Expect(err).ToNot(HaveOccurred())
			Expect(net).ToNot(BeNil())
			Expect(net.Id).To(Equal(testHnsNetID))
		})

		Specify("HNS endpoint operations work", func() {
			hnsEndpointConfig := &hcsshim.HNSEndpoint{
				VirtualNetwork: testHnsNetID,
				Name:           "ep_name",
			}

			endpointID, err := CreateHNSEndpoint(hnsEndpointConfig)
			Expect(err).ToNot(HaveOccurred())
			Expect(endpointID).ToNot(Equal(""))

			endpoint, err := GetHNSEndpoint(endpointID)
			Expect(err).ToNot(HaveOccurred())
			Expect(endpoint).ToNot(BeNil())

			expectNumberOfEndpoints(1)

			log.Infoln(endpoint)

			err = DeleteHNSEndpoint(endpointID)
			Expect(err).ToNot(HaveOccurred())

			endpoint, err = GetHNSEndpoint(endpointID)
			Expect(err).To(HaveOccurred())
			Expect(endpoint).To(BeNil())
		})

		Specify("Listing HNS endpoints works", func() {
			hnsEndpointConfig := &hcsshim.HNSEndpoint{
				VirtualNetwork: testHnsNetID,
			}

			endpointsList, err := ListHNSEndpoints()
			Expect(err).ToNot(HaveOccurred())
			numEndpointsOriginal := len(endpointsList)

			var endpoints [2]string
			for i := 0; i < 2; i++ {
				endpoints[i], err = CreateHNSEndpoint(hnsEndpointConfig)
				Expect(err).ToNot(HaveOccurred())
				Expect(endpoints[i]).ToNot(Equal(""))
			}

			expectNumberOfEndpoints(numEndpointsOriginal + 2)

			for _, ep := range endpoints {
				err = DeleteHNSEndpoint(ep)
				Expect(err).ToNot(HaveOccurred())
			}

			expectNumberOfEndpoints(numEndpointsOriginal)
		})

		Specify("Getting HNS endpoint by name works", func() {
			names := []string{"name1", "name2", "name3"}
			for _, name := range names {
				hnsEndpointConfig := &hcsshim.HNSEndpoint{
					VirtualNetwork: testHnsNetID,
					Name:           name,
				}
				_, err := CreateHNSEndpoint(hnsEndpointConfig)
				Expect(err).ToNot(HaveOccurred())
			}

			ep, err := GetHNSEndpointByName("name2")
			Expect(err).ToNot(HaveOccurred())
			Expect(ep.Name).To(Equal("name2"))
		})

		Context("There's a second HNS network", func() {
			secondHNSNetID := ""
			BeforeEach(func() {
				secondHNSNetID = MockHNSNetwork("other_net_name", netAdapter, subnetCIDR,
					defaultGW)
			})
			AfterEach(func() {
				err := DeleteHNSNetwork(secondHNSNetID)
				Expect(err).ToNot(HaveOccurred())
			})
			Specify("Listing HNS endpoints of specific network works", func() {
				config1 := &hcsshim.HNSEndpoint{
					VirtualNetwork: testHnsNetID,
				}
				config2 := &hcsshim.HNSEndpoint{
					VirtualNetwork: secondHNSNetID,
				}

				var epsInFirstNet []string
				var epsInSecondNet []string

				// create 3 endpoints in each network
				for i := 0; i < 3; i++ {
					ep1, err := CreateHNSEndpoint(config1)
					Expect(err).ToNot(HaveOccurred())

					epsInFirstNet = append(epsInFirstNet, ep1)

					ep2, err := CreateHNSEndpoint(config2)
					Expect(err).ToNot(HaveOccurred())

					epsInSecondNet = append(epsInSecondNet, ep2)
				}

				foundEpsOfFirstNet, err := ListHNSEndpointsOfNetwork(testHnsNetID)
				Expect(err).ToNot(HaveOccurred())
				Expect(foundEpsOfFirstNet).To(HaveLen(3))
				for _, ep := range foundEpsOfFirstNet {
					Expect(epsInFirstNet).To(ContainElement(ep.Id))
					Expect(epsInSecondNet).ToNot(ContainElement(ep.Id))
				}

				foundEpsOfSecondNet, err := ListHNSEndpointsOfNetwork(secondHNSNetID)
				Expect(err).ToNot(HaveOccurred())
				Expect(foundEpsOfSecondNet).To(HaveLen(3))
				for _, ep := range foundEpsOfSecondNet {
					Expect(epsInSecondNet).To(ContainElement(ep.Id))
					Expect(epsInFirstNet).ToNot(ContainElement(ep.Id))
				}
			})
		})

		Specify("Creating endpoint in same subnet works", func() {
			_, err := CreateHNSEndpoint(&hcsshim.HNSEndpoint{
				VirtualNetwork: testHnsNetID,
				IPAddress:      net.ParseIP("10.0.0.4"),
			})
			Expect(err).ToNot(HaveOccurred())
			expectNumberOfEndpoints(1)
		})

		Specify("Creating endpoint in different subnet fails", func() {
			_, err := CreateHNSEndpoint(&hcsshim.HNSEndpoint{
				VirtualNetwork: testHnsNetID,
				IPAddress:      net.ParseIP("10.1.0.4"),
			})
			Expect(err).To(HaveOccurred())
			expectNumberOfEndpoints(0)
		})

		Specify("Creating two endpoints with same IP works in same subnet fails", func() {
			_, err := CreateHNSEndpoint(&hcsshim.HNSEndpoint{
				VirtualNetwork: testHnsNetID,
				IPAddress:      net.ParseIP("10.0.0.4"),
			})
			Expect(err).ToNot(HaveOccurred())

			_, err = CreateHNSEndpoint(&hcsshim.HNSEndpoint{
				VirtualNetwork: testHnsNetID,
				IPAddress:      net.ParseIP("10.0.0.4"),
			})
			Expect(err).To(HaveOccurred())

			expectNumberOfEndpoints(1)
		})

		type MACTestCase struct {
			MAC        string
			shouldFail bool
		}
		DescribeTable("Creating an endpoint with specific MACs",
			func(t MACTestCase) {
				epID, err := CreateHNSEndpoint(&hcsshim.HNSEndpoint{
					VirtualNetwork: testHnsNetID,
					MacAddress:     t.MAC,
				})
				if t.shouldFail {
					Expect(err).To(HaveOccurred())
					expectNumberOfEndpoints(0)
				} else {
					Expect(err).ToNot(HaveOccurred())
					ep, err := GetHNSEndpoint(epID)
					Expect(err).ToNot(HaveOccurred())
					Expect(ep.MacAddress).To(Equal(t.MAC))
					expectNumberOfEndpoints(1)
				}
			},
			Entry("11-22-33-44-55-66 works", MACTestCase{
				MAC:        "11-22-33-44-55-66",
				shouldFail: false,
			}),
			Entry("AA-BB-CC-DD-EE-FF works", MACTestCase{
				MAC:        "AA-BB-CC-DD-EE-FF",
				shouldFail: false,
			}),
			Entry("XX-YY-11-22-33-44 fails", MACTestCase{
				MAC:        "XX-YY-11-22-33-44",
				shouldFail: true,
			}),
		)

		Specify("Creating multiple endpoints with conflicting MACs works", func() {
			cfg := &hcsshim.HNSEndpoint{
				VirtualNetwork: testHnsNetID,
				MacAddress:     "11-22-33-44-55-66",
			}
			for i := 0; i < 3; i++ {
				_, err := CreateHNSEndpoint(cfg)
				Expect(err).ToNot(HaveOccurred())
			}
			expectNumberOfEndpoints(3)
		})

		Specify("Creating endpoint with name containing special characters works", func() {
			cfg := &hcsshim.HNSEndpoint{
				VirtualNetwork: testHnsNetID,
				Name:           "A:B123/123",
			}
			_, err := CreateHNSEndpoint(cfg)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("HNS network doesn't exist", func() {

		BeforeEach(func() {
			nets, err := ListHNSNetworks()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(nets)).To(Equal(originalNumNetworks))
		})

		AfterEach(func() {
			nets, err := ListHNSNetworks()
			Expect(err).ToNot(HaveOccurred())
			for _, n := range nets {
				if strings.Contains(n.Name, "nat") {
					continue
				}
				err = DeleteHNSNetwork(n.Id)
				Expect(err).ToNot(HaveOccurred())
			}
		})

		Specify("getting single HNS network returns error", func() {
			net, err := GetHNSNetwork("1234abcd")
			Expect(err).To(HaveOccurred())
			Expect(net).To(BeNil())
		})

		Specify("getting single HNS network by name returns nil, nil", func() {
			net, err := GetHNSNetworkByName("asdf")
			Expect(err).To(BeNil())
			Expect(net).To(BeNil())
		})
	})
})

func expectNumberOfEndpoints(num int) {
	eps, err := ListHNSEndpoints()
	Expect(err).ToNot(HaveOccurred())
	Expect(eps).To(HaveLen(num))
}
