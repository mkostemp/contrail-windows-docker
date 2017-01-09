//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	qos_forwarding_class_dscp uint64 = 1 << iota
	qos_forwarding_class_trusted
	qos_forwarding_class_id_perms
	qos_forwarding_class_display_name
	qos_forwarding_class_qos_queue_refs
	qos_forwarding_class_virtual_network_back_refs
	qos_forwarding_class_virtual_machine_interface_back_refs
)

type QosForwardingClass struct {
        contrail.ObjectBase
	dscp int
	trusted bool
	id_perms IdPermsType
	display_name string
	qos_queue_refs contrail.ReferenceList
	virtual_network_back_refs contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *QosForwardingClass) GetType() string {
        return "qos-forwarding-class"
}

func (obj *QosForwardingClass) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *QosForwardingClass) GetDefaultParentType() string {
        return "project"
}

func (obj *QosForwardingClass) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *QosForwardingClass) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *QosForwardingClass) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *QosForwardingClass) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *QosForwardingClass) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *QosForwardingClass) GetDscp() int {
        return obj.dscp
}

func (obj *QosForwardingClass) SetDscp(value int) {
        obj.dscp = value
        obj.modified |= qos_forwarding_class_dscp
}

func (obj *QosForwardingClass) GetTrusted() bool {
        return obj.trusted
}

func (obj *QosForwardingClass) SetTrusted(value bool) {
        obj.trusted = value
        obj.modified |= qos_forwarding_class_trusted
}

func (obj *QosForwardingClass) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *QosForwardingClass) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= qos_forwarding_class_id_perms
}

func (obj *QosForwardingClass) GetDisplayName() string {
        return obj.display_name
}

func (obj *QosForwardingClass) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= qos_forwarding_class_display_name
}

func (obj *QosForwardingClass) readQosQueueRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & qos_forwarding_class_qos_queue_refs == 0) {
                err := obj.GetField(obj, "qos_queue_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosForwardingClass) GetQosQueueRefs() (
        contrail.ReferenceList, error) {
        err := obj.readQosQueueRefs()
        if err != nil {
                return nil, err
        }
        return obj.qos_queue_refs, nil
}

func (obj *QosForwardingClass) AddQosQueue(
        rhs *QosQueue) error {
        err := obj.readQosQueueRefs()
        if err != nil {
                return err
        }

        if obj.modified & qos_forwarding_class_qos_queue_refs == 0 {
                obj.storeReferenceBase("qos-queue", obj.qos_queue_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.qos_queue_refs = append(obj.qos_queue_refs, ref)
        obj.modified |= qos_forwarding_class_qos_queue_refs
        return nil
}

func (obj *QosForwardingClass) DeleteQosQueue(uuid string) error {
        err := obj.readQosQueueRefs()
        if err != nil {
                return err
        }

        if obj.modified & qos_forwarding_class_qos_queue_refs == 0 {
                obj.storeReferenceBase("qos-queue", obj.qos_queue_refs)
        }

        for i, ref := range obj.qos_queue_refs {
                if ref.Uuid == uuid {
                        obj.qos_queue_refs = append(
                                obj.qos_queue_refs[:i],
                                obj.qos_queue_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= qos_forwarding_class_qos_queue_refs
        return nil
}

func (obj *QosForwardingClass) ClearQosQueue() {
        if (obj.valid & qos_forwarding_class_qos_queue_refs != 0) &&
           (obj.modified & qos_forwarding_class_qos_queue_refs == 0) {
                obj.storeReferenceBase("qos-queue", obj.qos_queue_refs)
        }
        obj.qos_queue_refs = make([]contrail.Reference, 0)
        obj.valid |= qos_forwarding_class_qos_queue_refs
        obj.modified |= qos_forwarding_class_qos_queue_refs
}

func (obj *QosForwardingClass) SetQosQueueList(
        refList []contrail.ReferencePair) {
        obj.ClearQosQueue()
        obj.qos_queue_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.qos_queue_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *QosForwardingClass) readVirtualNetworkBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & qos_forwarding_class_virtual_network_back_refs == 0) {
                err := obj.GetField(obj, "virtual_network_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosForwardingClass) GetVirtualNetworkBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualNetworkBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_network_back_refs, nil
}

func (obj *QosForwardingClass) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & qos_forwarding_class_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosForwardingClass) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *QosForwardingClass) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & qos_forwarding_class_dscp != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.dscp)
                if err != nil {
                        return nil, err
                }
                msg["dscp"] = &value
        }

        if obj.modified & qos_forwarding_class_trusted != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.trusted)
                if err != nil {
                        return nil, err
                }
                msg["trusted"] = &value
        }

        if obj.modified & qos_forwarding_class_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & qos_forwarding_class_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.qos_queue_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.qos_queue_refs)
                if err != nil {
                        return nil, err
                }
                msg["qos_queue_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *QosForwardingClass) UnmarshalJSON(body []byte) error {
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
                case "dscp":
                        err = json.Unmarshal(value, &obj.dscp)
                        if err == nil {
                                obj.valid |= qos_forwarding_class_dscp
                        }
                        break
                case "trusted":
                        err = json.Unmarshal(value, &obj.trusted)
                        if err == nil {
                                obj.valid |= qos_forwarding_class_trusted
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= qos_forwarding_class_id_perms
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= qos_forwarding_class_display_name
                        }
                        break
                case "qos_queue_refs":
                        err = json.Unmarshal(value, &obj.qos_queue_refs)
                        if err == nil {
                                obj.valid |= qos_forwarding_class_qos_queue_refs
                        }
                        break
                case "virtual_network_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_network_back_refs)
                        if err == nil {
                                obj.valid |= qos_forwarding_class_virtual_network_back_refs
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= qos_forwarding_class_virtual_machine_interface_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosForwardingClass) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & qos_forwarding_class_dscp != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.dscp)
                if err != nil {
                        return nil, err
                }
                msg["dscp"] = &value
        }

        if obj.modified & qos_forwarding_class_trusted != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.trusted)
                if err != nil {
                        return nil, err
                }
                msg["trusted"] = &value
        }

        if obj.modified & qos_forwarding_class_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & qos_forwarding_class_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & qos_forwarding_class_qos_queue_refs != 0 {
                if len(obj.qos_queue_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["qos_queue_refs"] = &value
                } else if !obj.hasReferenceBase("qos-queue") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.qos_queue_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["qos_queue_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *QosForwardingClass) UpdateReferences() error {

        if (obj.modified & qos_forwarding_class_qos_queue_refs != 0) &&
           len(obj.qos_queue_refs) > 0 &&
           obj.hasReferenceBase("qos-queue") {
                err := obj.UpdateReference(
                        obj, "qos-queue",
                        obj.qos_queue_refs,
                        obj.baseMap["qos-queue"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func QosForwardingClassByName(c contrail.ApiClient, fqn string) (*QosForwardingClass, error) {
    obj, err := c.FindByName("qos-forwarding-class", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*QosForwardingClass), nil
}

func QosForwardingClassByUuid(c contrail.ApiClient, uuid string) (*QosForwardingClass, error) {
    obj, err := c.FindByUuid("qos-forwarding-class", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*QosForwardingClass), nil
}
