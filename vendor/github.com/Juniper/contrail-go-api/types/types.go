
package types

import (
        "reflect"

        "github.com/Juniper/contrail-go-api"
)

var (
        TypeMap = map[string]reflect.Type {
		"domain": reflect.TypeOf(Domain{}),
		"global-vrouter-config": reflect.TypeOf(GlobalVrouterConfig{}),
		"instance-ip": reflect.TypeOf(InstanceIp{}),
		"network-policy": reflect.TypeOf(NetworkPolicy{}),
		"loadbalancer-pool": reflect.TypeOf(LoadbalancerPool{}),
		"virtual-DNS-record": reflect.TypeOf(VirtualDnsRecord{}),
		"route-target": reflect.TypeOf(RouteTarget{}),
		"floating-ip": reflect.TypeOf(FloatingIp{}),
		"floating-ip-pool": reflect.TypeOf(FloatingIpPool{}),
		"physical-router": reflect.TypeOf(PhysicalRouter{}),
		"bgp-router": reflect.TypeOf(BgpRouter{}),
		"virtual-router": reflect.TypeOf(VirtualRouter{}),
		"config-root": reflect.TypeOf(ConfigRoot{}),
		"subnet": reflect.TypeOf(Subnet{}),
		"global-system-config": reflect.TypeOf(GlobalSystemConfig{}),
		"service-appliance": reflect.TypeOf(ServiceAppliance{}),
		"service-instance": reflect.TypeOf(ServiceInstance{}),
		"namespace": reflect.TypeOf(Namespace{}),
		"logical-interface": reflect.TypeOf(LogicalInterface{}),
		"route-table": reflect.TypeOf(RouteTable{}),
		"physical-interface": reflect.TypeOf(PhysicalInterface{}),
		"access-control-list": reflect.TypeOf(AccessControlList{}),
		"analytics-node": reflect.TypeOf(AnalyticsNode{}),
		"virtual-DNS": reflect.TypeOf(VirtualDns{}),
		"customer-attachment": reflect.TypeOf(CustomerAttachment{}),
		"service-appliance-set": reflect.TypeOf(ServiceApplianceSet{}),
		"config-node": reflect.TypeOf(ConfigNode{}),
		"qos-queue": reflect.TypeOf(QosQueue{}),
		"virtual-machine": reflect.TypeOf(VirtualMachine{}),
		"interface-route-table": reflect.TypeOf(InterfaceRouteTable{}),
		"service-template": reflect.TypeOf(ServiceTemplate{}),
		"virtual-ip": reflect.TypeOf(VirtualIp{}),
		"loadbalancer-member": reflect.TypeOf(LoadbalancerMember{}),
		"security-group": reflect.TypeOf(SecurityGroup{}),
		"provider-attachment": reflect.TypeOf(ProviderAttachment{}),
		"virtual-machine-interface": reflect.TypeOf(VirtualMachineInterface{}),
		"loadbalancer-healthmonitor": reflect.TypeOf(LoadbalancerHealthmonitor{}),
		"virtual-network": reflect.TypeOf(VirtualNetwork{}),
		"project": reflect.TypeOf(Project{}),
		"qos-forwarding-class": reflect.TypeOf(QosForwardingClass{}),
		"database-node": reflect.TypeOf(DatabaseNode{}),
		"routing-instance": reflect.TypeOf(RoutingInstance{}),
		"network-ipam": reflect.TypeOf(NetworkIpam{}),
		"logical-router": reflect.TypeOf(LogicalRouter{}),

        }
)

func init() {
        contrail.RegisterTypeMap(TypeMap)
}
