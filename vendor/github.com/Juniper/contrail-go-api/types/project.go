//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	project_quota uint64 = 1 << iota
	project_id_perms
	project_display_name
	project_namespace_refs
	project_security_groups
	project_virtual_networks
	project_qos_queues
	project_qos_forwarding_classs
	project_network_ipams
	project_network_policys
	project_virtual_machine_interfaces
	project_floating_ip_pool_refs
	project_service_instances
	project_route_tables
	project_interface_route_tables
	project_logical_routers
	project_loadbalancer_pools
	project_loadbalancer_healthmonitors
	project_virtual_ips
	project_floating_ip_back_refs
)

type Project struct {
        contrail.ObjectBase
	quota QuotaType
	id_perms IdPermsType
	display_name string
	namespace_refs contrail.ReferenceList
	security_groups contrail.ReferenceList
	virtual_networks contrail.ReferenceList
	qos_queues contrail.ReferenceList
	qos_forwarding_classs contrail.ReferenceList
	network_ipams contrail.ReferenceList
	network_policys contrail.ReferenceList
	virtual_machine_interfaces contrail.ReferenceList
	floating_ip_pool_refs contrail.ReferenceList
	service_instances contrail.ReferenceList
	route_tables contrail.ReferenceList
	interface_route_tables contrail.ReferenceList
	logical_routers contrail.ReferenceList
	loadbalancer_pools contrail.ReferenceList
	loadbalancer_healthmonitors contrail.ReferenceList
	virtual_ips contrail.ReferenceList
	floating_ip_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *Project) GetType() string {
        return "project"
}

func (obj *Project) GetDefaultParent() []string {
        name := []string{"default-domain"}
        return name
}

func (obj *Project) GetDefaultParentType() string {
        return "domain"
}

func (obj *Project) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *Project) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *Project) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *Project) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *Project) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *Project) GetQuota() QuotaType {
        return obj.quota
}

func (obj *Project) SetQuota(value *QuotaType) {
        obj.quota = *value
        obj.modified |= project_quota
}

func (obj *Project) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *Project) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= project_id_perms
}

func (obj *Project) GetDisplayName() string {
        return obj.display_name
}

func (obj *Project) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= project_display_name
}

func (obj *Project) readSecurityGroups() error {
        if !obj.IsTransient() &&
                (obj.valid & project_security_groups == 0) {
                err := obj.GetField(obj, "security_groups")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetSecurityGroups() (
        contrail.ReferenceList, error) {
        err := obj.readSecurityGroups()
        if err != nil {
                return nil, err
        }
        return obj.security_groups, nil
}

func (obj *Project) readVirtualNetworks() error {
        if !obj.IsTransient() &&
                (obj.valid & project_virtual_networks == 0) {
                err := obj.GetField(obj, "virtual_networks")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetVirtualNetworks() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualNetworks()
        if err != nil {
                return nil, err
        }
        return obj.virtual_networks, nil
}

func (obj *Project) readQosQueues() error {
        if !obj.IsTransient() &&
                (obj.valid & project_qos_queues == 0) {
                err := obj.GetField(obj, "qos_queues")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetQosQueues() (
        contrail.ReferenceList, error) {
        err := obj.readQosQueues()
        if err != nil {
                return nil, err
        }
        return obj.qos_queues, nil
}

func (obj *Project) readQosForwardingClasss() error {
        if !obj.IsTransient() &&
                (obj.valid & project_qos_forwarding_classs == 0) {
                err := obj.GetField(obj, "qos_forwarding_classs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetQosForwardingClasss() (
        contrail.ReferenceList, error) {
        err := obj.readQosForwardingClasss()
        if err != nil {
                return nil, err
        }
        return obj.qos_forwarding_classs, nil
}

func (obj *Project) readNetworkIpams() error {
        if !obj.IsTransient() &&
                (obj.valid & project_network_ipams == 0) {
                err := obj.GetField(obj, "network_ipams")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetNetworkIpams() (
        contrail.ReferenceList, error) {
        err := obj.readNetworkIpams()
        if err != nil {
                return nil, err
        }
        return obj.network_ipams, nil
}

func (obj *Project) readNetworkPolicys() error {
        if !obj.IsTransient() &&
                (obj.valid & project_network_policys == 0) {
                err := obj.GetField(obj, "network_policys")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetNetworkPolicys() (
        contrail.ReferenceList, error) {
        err := obj.readNetworkPolicys()
        if err != nil {
                return nil, err
        }
        return obj.network_policys, nil
}

func (obj *Project) readVirtualMachineInterfaces() error {
        if !obj.IsTransient() &&
                (obj.valid & project_virtual_machine_interfaces == 0) {
                err := obj.GetField(obj, "virtual_machine_interfaces")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetVirtualMachineInterfaces() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaces()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interfaces, nil
}

func (obj *Project) readServiceInstances() error {
        if !obj.IsTransient() &&
                (obj.valid & project_service_instances == 0) {
                err := obj.GetField(obj, "service_instances")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetServiceInstances() (
        contrail.ReferenceList, error) {
        err := obj.readServiceInstances()
        if err != nil {
                return nil, err
        }
        return obj.service_instances, nil
}

func (obj *Project) readRouteTables() error {
        if !obj.IsTransient() &&
                (obj.valid & project_route_tables == 0) {
                err := obj.GetField(obj, "route_tables")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetRouteTables() (
        contrail.ReferenceList, error) {
        err := obj.readRouteTables()
        if err != nil {
                return nil, err
        }
        return obj.route_tables, nil
}

func (obj *Project) readInterfaceRouteTables() error {
        if !obj.IsTransient() &&
                (obj.valid & project_interface_route_tables == 0) {
                err := obj.GetField(obj, "interface_route_tables")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetInterfaceRouteTables() (
        contrail.ReferenceList, error) {
        err := obj.readInterfaceRouteTables()
        if err != nil {
                return nil, err
        }
        return obj.interface_route_tables, nil
}

func (obj *Project) readLogicalRouters() error {
        if !obj.IsTransient() &&
                (obj.valid & project_logical_routers == 0) {
                err := obj.GetField(obj, "logical_routers")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetLogicalRouters() (
        contrail.ReferenceList, error) {
        err := obj.readLogicalRouters()
        if err != nil {
                return nil, err
        }
        return obj.logical_routers, nil
}

func (obj *Project) readLoadbalancerPools() error {
        if !obj.IsTransient() &&
                (obj.valid & project_loadbalancer_pools == 0) {
                err := obj.GetField(obj, "loadbalancer_pools")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetLoadbalancerPools() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerPools()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_pools, nil
}

func (obj *Project) readLoadbalancerHealthmonitors() error {
        if !obj.IsTransient() &&
                (obj.valid & project_loadbalancer_healthmonitors == 0) {
                err := obj.GetField(obj, "loadbalancer_healthmonitors")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetLoadbalancerHealthmonitors() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerHealthmonitors()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_healthmonitors, nil
}

func (obj *Project) readVirtualIps() error {
        if !obj.IsTransient() &&
                (obj.valid & project_virtual_ips == 0) {
                err := obj.GetField(obj, "virtual_ips")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetVirtualIps() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualIps()
        if err != nil {
                return nil, err
        }
        return obj.virtual_ips, nil
}

func (obj *Project) readNamespaceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & project_namespace_refs == 0) {
                err := obj.GetField(obj, "namespace_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetNamespaceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readNamespaceRefs()
        if err != nil {
                return nil, err
        }
        return obj.namespace_refs, nil
}

func (obj *Project) AddNamespace(
        rhs *Namespace, data SubnetType) error {
        err := obj.readNamespaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & project_namespace_refs == 0 {
                obj.storeReferenceBase("namespace", obj.namespace_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
        obj.namespace_refs = append(obj.namespace_refs, ref)
        obj.modified |= project_namespace_refs
        return nil
}

func (obj *Project) DeleteNamespace(uuid string) error {
        err := obj.readNamespaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & project_namespace_refs == 0 {
                obj.storeReferenceBase("namespace", obj.namespace_refs)
        }

        for i, ref := range obj.namespace_refs {
                if ref.Uuid == uuid {
                        obj.namespace_refs = append(
                                obj.namespace_refs[:i],
                                obj.namespace_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= project_namespace_refs
        return nil
}

func (obj *Project) ClearNamespace() {
        if (obj.valid & project_namespace_refs != 0) &&
           (obj.modified & project_namespace_refs == 0) {
                obj.storeReferenceBase("namespace", obj.namespace_refs)
        }
        obj.namespace_refs = make([]contrail.Reference, 0)
        obj.valid |= project_namespace_refs
        obj.modified |= project_namespace_refs
}

func (obj *Project) SetNamespaceList(
        refList []contrail.ReferencePair) {
        obj.ClearNamespace()
        obj.namespace_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.namespace_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *Project) readFloatingIpPoolRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & project_floating_ip_pool_refs == 0) {
                err := obj.GetField(obj, "floating_ip_pool_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetFloatingIpPoolRefs() (
        contrail.ReferenceList, error) {
        err := obj.readFloatingIpPoolRefs()
        if err != nil {
                return nil, err
        }
        return obj.floating_ip_pool_refs, nil
}

func (obj *Project) AddFloatingIpPool(
        rhs *FloatingIpPool) error {
        err := obj.readFloatingIpPoolRefs()
        if err != nil {
                return err
        }

        if obj.modified & project_floating_ip_pool_refs == 0 {
                obj.storeReferenceBase("floating-ip-pool", obj.floating_ip_pool_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.floating_ip_pool_refs = append(obj.floating_ip_pool_refs, ref)
        obj.modified |= project_floating_ip_pool_refs
        return nil
}

func (obj *Project) DeleteFloatingIpPool(uuid string) error {
        err := obj.readFloatingIpPoolRefs()
        if err != nil {
                return err
        }

        if obj.modified & project_floating_ip_pool_refs == 0 {
                obj.storeReferenceBase("floating-ip-pool", obj.floating_ip_pool_refs)
        }

        for i, ref := range obj.floating_ip_pool_refs {
                if ref.Uuid == uuid {
                        obj.floating_ip_pool_refs = append(
                                obj.floating_ip_pool_refs[:i],
                                obj.floating_ip_pool_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= project_floating_ip_pool_refs
        return nil
}

func (obj *Project) ClearFloatingIpPool() {
        if (obj.valid & project_floating_ip_pool_refs != 0) &&
           (obj.modified & project_floating_ip_pool_refs == 0) {
                obj.storeReferenceBase("floating-ip-pool", obj.floating_ip_pool_refs)
        }
        obj.floating_ip_pool_refs = make([]contrail.Reference, 0)
        obj.valid |= project_floating_ip_pool_refs
        obj.modified |= project_floating_ip_pool_refs
}

func (obj *Project) SetFloatingIpPoolList(
        refList []contrail.ReferencePair) {
        obj.ClearFloatingIpPool()
        obj.floating_ip_pool_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.floating_ip_pool_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *Project) readFloatingIpBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & project_floating_ip_back_refs == 0) {
                err := obj.GetField(obj, "floating_ip_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) GetFloatingIpBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readFloatingIpBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.floating_ip_back_refs, nil
}

func (obj *Project) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & project_quota != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.quota)
                if err != nil {
                        return nil, err
                }
                msg["quota"] = &value
        }

        if obj.modified & project_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & project_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.namespace_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.namespace_refs)
                if err != nil {
                        return nil, err
                }
                msg["namespace_refs"] = &value
        }

        if len(obj.floating_ip_pool_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_pool_refs)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_pool_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *Project) UnmarshalJSON(body []byte) error {
        var m map[string]json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
                return err
        }
        err = obj.UnmarshalCommon(m)
        if err != nil {
                return err
        }
        for key, value := range m {
                switch key {
                case "quota":
                        err = json.Unmarshal(value, &obj.quota)
                        if err == nil {
                                obj.valid |= project_quota
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= project_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= project_display_name
                        }
                        break
                case "security_groups":
                        err = json.Unmarshal(value, &obj.security_groups)
                        if err == nil {
                                obj.valid |= project_security_groups
                        }
                        break
                case "virtual_networks":
                        err = json.Unmarshal(value, &obj.virtual_networks)
                        if err == nil {
                                obj.valid |= project_virtual_networks
                        }
                        break
                case "qos_queues":
                        err = json.Unmarshal(value, &obj.qos_queues)
                        if err == nil {
                                obj.valid |= project_qos_queues
                        }
                        break
                case "qos_forwarding_classs":
                        err = json.Unmarshal(value, &obj.qos_forwarding_classs)
                        if err == nil {
                                obj.valid |= project_qos_forwarding_classs
                        }
                        break
                case "network_ipams":
                        err = json.Unmarshal(value, &obj.network_ipams)
                        if err == nil {
                                obj.valid |= project_network_ipams
                        }
                        break
                case "network_policys":
                        err = json.Unmarshal(value, &obj.network_policys)
                        if err == nil {
                                obj.valid |= project_network_policys
                        }
                        break
                case "virtual_machine_interfaces":
                        err = json.Unmarshal(value, &obj.virtual_machine_interfaces)
                        if err == nil {
                                obj.valid |= project_virtual_machine_interfaces
                        }
                        break
                case "floating_ip_pool_refs":
                        err = json.Unmarshal(value, &obj.floating_ip_pool_refs)
                        if err == nil {
                                obj.valid |= project_floating_ip_pool_refs
                        }
                        break
                case "service_instances":
                        err = json.Unmarshal(value, &obj.service_instances)
                        if err == nil {
                                obj.valid |= project_service_instances
                        }
                        break
                case "route_tables":
                        err = json.Unmarshal(value, &obj.route_tables)
                        if err == nil {
                                obj.valid |= project_route_tables
                        }
                        break
                case "interface_route_tables":
                        err = json.Unmarshal(value, &obj.interface_route_tables)
                        if err == nil {
                                obj.valid |= project_interface_route_tables
                        }
                        break
                case "logical_routers":
                        err = json.Unmarshal(value, &obj.logical_routers)
                        if err == nil {
                                obj.valid |= project_logical_routers
                        }
                        break
                case "loadbalancer_pools":
                        err = json.Unmarshal(value, &obj.loadbalancer_pools)
                        if err == nil {
                                obj.valid |= project_loadbalancer_pools
                        }
                        break
                case "loadbalancer_healthmonitors":
                        err = json.Unmarshal(value, &obj.loadbalancer_healthmonitors)
                        if err == nil {
                                obj.valid |= project_loadbalancer_healthmonitors
                        }
                        break
                case "virtual_ips":
                        err = json.Unmarshal(value, &obj.virtual_ips)
                        if err == nil {
                                obj.valid |= project_virtual_ips
                        }
                        break
                case "floating_ip_back_refs":
                        err = json.Unmarshal(value, &obj.floating_ip_back_refs)
                        if err == nil {
                                obj.valid |= project_floating_ip_back_refs
                        }
                        break
                case "namespace_refs": {
                        type ReferenceElement struct {
                                To []string
                                Uuid string
                                Href string
                                Attr SubnetType
                        }
                        var array []ReferenceElement
                        err = json.Unmarshal(value, &array)
                        if err != nil {
                            break
                        }
                        obj.valid |= project_namespace_refs
                        obj.namespace_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.namespace_refs = append(obj.namespace_refs, ref)
                        }
                        break
                }
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Project) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & project_quota != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.quota)
                if err != nil {
                        return nil, err
                }
                msg["quota"] = &value
        }

        if obj.modified & project_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & project_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & project_namespace_refs != 0 {
                if len(obj.namespace_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["namespace_refs"] = &value
                } else if !obj.hasReferenceBase("namespace") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.namespace_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["namespace_refs"] = &value
                }
        }


        if obj.modified & project_floating_ip_pool_refs != 0 {
                if len(obj.floating_ip_pool_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["floating_ip_pool_refs"] = &value
                } else if !obj.hasReferenceBase("floating-ip-pool") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.floating_ip_pool_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["floating_ip_pool_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *Project) UpdateReferences() error {

        if (obj.modified & project_namespace_refs != 0) &&
           len(obj.namespace_refs) > 0 &&
           obj.hasReferenceBase("namespace") {
                err := obj.UpdateReference(
                        obj, "namespace",
                        obj.namespace_refs,
                        obj.baseMap["namespace"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & project_floating_ip_pool_refs != 0) &&
           len(obj.floating_ip_pool_refs) > 0 &&
           obj.hasReferenceBase("floating-ip-pool") {
                err := obj.UpdateReference(
                        obj, "floating-ip-pool",
                        obj.floating_ip_pool_refs,
                        obj.baseMap["floating-ip-pool"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func ProjectByName(c contrail.ApiClient, fqn string) (*Project, error) {
    obj, err := c.FindByName("project", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*Project), nil
}

func ProjectByUuid(c contrail.ApiClient, uuid string) (*Project, error) {
    obj, err := c.FindByUuid("project", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*Project), nil
}
