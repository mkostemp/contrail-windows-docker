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
		testHnsID := ""

		subnets := []hcsshim.Subnet{
			{
				AddressPrefix:  "1.1.1.0/24",
				GatewayAddress: "1.1.1.1",
			},
		}
		configuration := &hcsshim.HNSNetwork{
			Name:               testNetName,
			Type:               "transparent",
			Subnets:            subnets,
			NetworkAdapterName: "Ethernet0",
		}

		BeforeEach(func() {
			Expect(testHnsID).To(Equal(""))
			var err error
			testHnsID, err = CreateHNSNetwork(configuration)
			Expect(err).ToNot(HaveOccurred())
			Expect(testHnsID).ToNot(Equal(""))
		})

		AfterEach(func() {
			Expect(testHnsID).ToNot(Equal(""))
			err := DeleteHNSNetwork(testHnsID)
			Expect(err).ToNot(HaveOccurred())
			_, err = GetHNSNetwork(testHnsID)
			Expect(err).To(HaveOccurred())
			testHnsID = ""
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
				if n.Id == testHnsID {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue())
		})

		Specify("getting a single HNS network works", func() {
			net, err := GetHNSNetwork(testHnsID)
			Expect(err).ToNot(HaveOccurred())
			Expect(net).ToNot(BeNil())
			Expect(net.Id).To(Equal(testHnsID))
		})

		Specify("getting a single HNS network by name works", func() {
			net, err := GetHNSNetworkByName(testNetName)
			Expect(err).ToNot(HaveOccurred())
			Expect(net).ToNot(BeNil())
			Expect(net.Id).To(Equal(testHnsID))
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
