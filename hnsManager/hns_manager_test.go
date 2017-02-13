package hnsManager

import (
	"flag"
	"fmt"
	"testing"

	"github.com/codilime/contrail-windows-docker/common"
	"github.com/codilime/contrail-windows-docker/hns"
	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var netAdapter string

func init() {
	flag.StringVar(&netAdapter, "netAdapter", "Ethernet0", "Ethernet adapter name to use")
	log.SetLevel(log.DebugLevel)
}

func TestHNSManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HNS manager test suite")
}

var _ = BeforeSuite(func() {
	err := common.HardResetHNS()
	Expect(err).ToNot(HaveOccurred())
})

var _ = Describe("HNS manager", func() {

	const (
		tenantName  = "agatka"
		networkName = "test_net"
		subnetCIDR  = "10.0.0.0/24"
		defaultGW   = "10.0.0.1"
	)

	var hnsMgr *HNSManager

	BeforeEach(func() {
		hnsMgr = &HNSManager{}
	})

	AfterEach(func() {
		err := common.HardResetHNS()
		Expect(err).ToNot(HaveOccurred())
	})

	Context("specified network does not exist", func() {
		Specify("creating a new HNS network works", func() {
			_, err := hnsMgr.CreateNetwork(netAdapter, tenantName, networkName,
				subnetCIDR, defaultGW)
			Expect(err).ToNot(HaveOccurred())
		})
		Specify("getting the HNS network returns error", func() {
			net, err := hnsMgr.GetNetwork(tenantName, networkName)
			Expect(err).To(HaveOccurred())
			Expect(net).To(BeNil())
		})
	})

	Context("specified network already exists", func() {
		var existingNetID string
		BeforeEach(func() {
			hnsNetName := fmt.Sprintf("Contrail:%s:%s", tenantName, networkName)
			existingNetID = hns.MockHNSNetwork(hnsNetName, netAdapter, subnetCIDR, defaultGW)
		})

		Specify("creating a new network with same params returns error", func() {
			net, err := hnsMgr.CreateNetwork(netAdapter, tenantName, networkName,
				subnetCIDR, defaultGW)
			Expect(err).To(HaveOccurred())
			Expect(net).To(BeNil())
		})

		Specify("getting the network returns it", func() {
			net, err := hnsMgr.GetNetwork(tenantName, networkName)
			Expect(err).ToNot(HaveOccurred())
			Expect(net.Id).To(Equal(existingNetID))
		})

		Context("network has active endpoints", func() {
			BeforeEach(func() {
				eps, err := hns.ListHNSEndpoints()
				Expect(err).ToNot(HaveOccurred())
				Expect(eps).To(BeEmpty())

				_ = hns.MockHNSEndpoint(existingNetID)

				eps, err = hns.ListHNSEndpoints()
				Expect(err).ToNot(HaveOccurred())
				Expect(eps).ToNot(BeEmpty())
			})

			Specify("deleting the network returns error", func() {
				err := hnsMgr.DeleteNetwork(tenantName, networkName)
				Expect(err).To(HaveOccurred())

				eps, err := hns.ListHNSEndpoints()
				Expect(err).ToNot(HaveOccurred())
				Expect(eps).ToNot(BeEmpty())
			})
		})

		Context("network has no active endpoints", func() {
			Specify("deleting the network removes it", func() {
				netsBefore, err := hns.ListHNSNetworks()
				Expect(err).ToNot(HaveOccurred())
				err = hnsMgr.DeleteNetwork(tenantName, networkName)
				Expect(err).ToNot(HaveOccurred())
				netsAfter, err := hns.ListHNSNetworks()
				Expect(err).ToNot(HaveOccurred())
				Expect(netsBefore).To(HaveLen(len(netsAfter) + 1))
			})
		})
	})

	Describe("Listing Contrail networks", func() {
		BeforeEach(func() {
			names := []string{
				fmt.Sprintf("Contrail:%s:%s", "tenant1", "netname1"),
				fmt.Sprintf("Contrail:%s:%s", "tenant2", "netname2"),
				fmt.Sprintf("Contrail:%s", "invalid_num_of_fields"),
				"some_other_name",
			}
			for _, n := range names {
				hns.MockHNSNetwork(n, netAdapter, subnetCIDR, defaultGW)
			}
		})
		Specify("Listing only Contrail networks works", func() {
			nets, err := hnsMgr.ListNetworks()
			Expect(err).ToNot(HaveOccurred())
			Expect(nets).To(HaveLen(2))
			for _, n := range nets {
				Expect(n.Name).To(ContainSubstring("Contrail:"))
			}
		})
	})
})
