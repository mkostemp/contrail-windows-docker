package hns

import (
	"strings"
	"testing"

	"github.com/Microsoft/hcsshim"
	"github.com/codilime/contrail-windows-docker/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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

	var originalNumNetworks int

	BeforeEach(func() {
		nets, err := ListHNSNetworks()
		Expect(err).ToNot(HaveOccurred())
		originalNumNetworks = len(nets)
	})

	Context("HNS network exists", func() {
		/*
			There's an issue with HNS where deleting a network and then creating one
			immediately after doesn't work (https://github.com/Microsoft/hcsshim/issues/95).
			A way to fix this is to have a long timeout before creating a network (like 20
			seconds), which is way too long for test suite such as this one.
			Ideally, we would like to call Create/Delete a test network for each of the
			following test cases, but it would take too long. So a single test network will
			be shared for all of them.
		*/

		testNetName := "TestNetwork"
		testHnsNetID := ""

		subnets := []hcsshim.Subnet{
			{
				AddressPrefix:  "1.1.1.0/24",
				GatewayAddress: "1.1.1.1",
			},
		}
		netConfiguration := &hcsshim.HNSNetwork{
			Name:               testNetName,
			Type:               "transparent",
			Subnets:            subnets,
			NetworkAdapterName: "Ethernet0",
		}

		BeforeEach(func() {
			Expect(testHnsNetID).To(Equal(""))
			var err error
			testHnsNetID, err = CreateHNSNetwork(netConfiguration)
			Expect(err).ToNot(HaveOccurred())
			Expect(testHnsNetID).ToNot(Equal(""))
		})

		AfterEach(func() {
			Expect(testHnsNetID).ToNot(Equal(""))
			err := DeleteHNSNetwork(testHnsNetID)
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
			}

			endpointID, err := CreateHNSEndpoint(hnsEndpointConfig)
			Expect(err).ToNot(HaveOccurred())
			Expect(endpointID).ToNot(Equal(""))

			endpoint, err := GetHNSEndpoint(endpointID)
			Expect(err).ToNot(HaveOccurred())
			Expect(endpoint).ToNot(BeNil())

			err = DeleteHNSEndpoint(endpointID)
			Expect(err).ToNot(HaveOccurred())

			endpoint, err = GetHNSEndpoint(endpointID)
			Expect(err).To(HaveOccurred())
			Expect(endpoint).To(BeNil())
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

		Specify("getting single HNS network by name returns error", func() {
			net, err := GetHNSNetworkByName("asdf")
			Expect(err).To(HaveOccurred())
			Expect(net).To(BeNil())
		})
	})
})
