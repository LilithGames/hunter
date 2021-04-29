// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: testdata/proto/test.proto

package metric

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Kubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Kubernetes) Reset() {
	*x = Kubernetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kubernetes) ProtoMessage() {}

func (x *Kubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kubernetes.ProtoReflect.Descriptor instead.
func (*Kubernetes) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0}
}

type KubernetesNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KubernetesNode) Reset() {
	*x = KubernetesNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesNode) ProtoMessage() {}

func (x *KubernetesNode) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesNode.ProtoReflect.Descriptor instead.
func (*KubernetesNode) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 0}
}

type KubernetesPodNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KubernetesPodNetwork) Reset() {
	*x = KubernetesPodNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodNetwork) ProtoMessage() {}

func (x *KubernetesPodNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodNetwork.ProtoReflect.Descriptor instead.
func (*KubernetesPodNetwork) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 1}
}

type KubernetesPodVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KubernetesPodVolume) Reset() {
	*x = KubernetesPodVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodVolume) ProtoMessage() {}

func (x *KubernetesPodVolume) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodVolume.ProtoReflect.Descriptor instead.
func (*KubernetesPodVolume) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 2}
}

type KubernetesPodContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KubernetesPodContainer) Reset() {
	*x = KubernetesPodContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodContainer) ProtoMessage() {}

func (x *KubernetesPodContainer) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodContainer.ProtoReflect.Descriptor instead.
func (*KubernetesPodContainer) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 3}
}

type KubernetesNodeTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}

func (x *KubernetesNodeTag) Reset() {
	*x = KubernetesNodeTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesNodeTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesNodeTag) ProtoMessage() {}

func (x *KubernetesNodeTag) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesNodeTag.ProtoReflect.Descriptor instead.
func (*KubernetesNodeTag) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *KubernetesNodeTag) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type KubernetesNodeField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuUsageCoreNanoseconds int64 `protobuf:"varint,1,opt,name=cpu_usage_core_nanoseconds,json=cpuUsageCoreNanoseconds,proto3" json:"cpu_usage_core_nanoseconds,omitempty"`
	MemoryAvailableBytes    int64 `protobuf:"varint,2,opt,name=memory_available_bytes,json=memoryAvailableBytes,proto3" json:"memory_available_bytes,omitempty"`
}

func (x *KubernetesNodeField) Reset() {
	*x = KubernetesNodeField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesNodeField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesNodeField) ProtoMessage() {}

func (x *KubernetesNodeField) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesNodeField.ProtoReflect.Descriptor instead.
func (*KubernetesNodeField) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *KubernetesNodeField) GetCpuUsageCoreNanoseconds() int64 {
	if x != nil {
		return x.CpuUsageCoreNanoseconds
	}
	return 0
}

func (x *KubernetesNodeField) GetMemoryAvailableBytes() int64 {
	if x != nil {
		return x.MemoryAvailableBytes
	}
	return 0
}

type KubernetesPodNetworkTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	NodeName  string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PodName   string `protobuf:"bytes,3,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
}

func (x *KubernetesPodNetworkTag) Reset() {
	*x = KubernetesPodNetworkTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodNetworkTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodNetworkTag) ProtoMessage() {}

func (x *KubernetesPodNetworkTag) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodNetworkTag.ProtoReflect.Descriptor instead.
func (*KubernetesPodNetworkTag) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *KubernetesPodNetworkTag) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubernetesPodNetworkTag) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *KubernetesPodNetworkTag) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

type KubernetesPodNetworkField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RxBytes  int64 `protobuf:"varint,1,opt,name=rx_bytes,json=rxBytes,proto3" json:"rx_bytes,omitempty"`
	RxErrors int64 `protobuf:"varint,2,opt,name=rx_errors,json=rxErrors,proto3" json:"rx_errors,omitempty"`
}

func (x *KubernetesPodNetworkField) Reset() {
	*x = KubernetesPodNetworkField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodNetworkField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodNetworkField) ProtoMessage() {}

func (x *KubernetesPodNetworkField) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodNetworkField.ProtoReflect.Descriptor instead.
func (*KubernetesPodNetworkField) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *KubernetesPodNetworkField) GetRxBytes() int64 {
	if x != nil {
		return x.RxBytes
	}
	return 0
}

func (x *KubernetesPodNetworkField) GetRxErrors() int64 {
	if x != nil {
		return x.RxErrors
	}
	return 0
}

type KubernetesPodVolumeTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	NodeName   string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PodName    string `protobuf:"bytes,3,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	VolumeName string `protobuf:"bytes,4,opt,name=volume_name,json=volumeName,proto3" json:"volume_name,omitempty"`
}

func (x *KubernetesPodVolumeTag) Reset() {
	*x = KubernetesPodVolumeTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodVolumeTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodVolumeTag) ProtoMessage() {}

func (x *KubernetesPodVolumeTag) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodVolumeTag.ProtoReflect.Descriptor instead.
func (*KubernetesPodVolumeTag) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *KubernetesPodVolumeTag) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubernetesPodVolumeTag) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *KubernetesPodVolumeTag) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *KubernetesPodVolumeTag) GetVolumeName() string {
	if x != nil {
		return x.VolumeName
	}
	return ""
}

type KubernetesPodVolumeField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableBytes int64 `protobuf:"varint,1,opt,name=available_bytes,json=availableBytes,proto3" json:"available_bytes,omitempty"`
	CapacityBytes  int64 `protobuf:"varint,2,opt,name=capacity_bytes,json=capacityBytes,proto3" json:"capacity_bytes,omitempty"`
	UsedBytes      int64 `protobuf:"varint,3,opt,name=used_bytes,json=usedBytes,proto3" json:"used_bytes,omitempty"`
}

func (x *KubernetesPodVolumeField) Reset() {
	*x = KubernetesPodVolumeField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodVolumeField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodVolumeField) ProtoMessage() {}

func (x *KubernetesPodVolumeField) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodVolumeField.ProtoReflect.Descriptor instead.
func (*KubernetesPodVolumeField) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *KubernetesPodVolumeField) GetAvailableBytes() int64 {
	if x != nil {
		return x.AvailableBytes
	}
	return 0
}

func (x *KubernetesPodVolumeField) GetCapacityBytes() int64 {
	if x != nil {
		return x.CapacityBytes
	}
	return 0
}

func (x *KubernetesPodVolumeField) GetUsedBytes() int64 {
	if x != nil {
		return x.UsedBytes
	}
	return 0
}

type KubernetesPodContainerTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName      string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ContainerName string `protobuf:"bytes,2,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
}

func (x *KubernetesPodContainerTag) Reset() {
	*x = KubernetesPodContainerTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodContainerTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodContainerTag) ProtoMessage() {}

func (x *KubernetesPodContainerTag) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodContainerTag.ProtoReflect.Descriptor instead.
func (*KubernetesPodContainerTag) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 3, 0}
}

func (x *KubernetesPodContainerTag) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *KubernetesPodContainerTag) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

type KubernetesPodContainerField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuUsageCoreNanoseconds int64 `protobuf:"varint,1,opt,name=cpu_usage_core_nanoseconds,json=cpuUsageCoreNanoseconds,proto3" json:"cpu_usage_core_nanoseconds,omitempty"`
	MemoryUsageBytes        int64 `protobuf:"varint,2,opt,name=memory_usage_bytes,json=memoryUsageBytes,proto3" json:"memory_usage_bytes,omitempty"`
}

func (x *KubernetesPodContainerField) Reset() {
	*x = KubernetesPodContainerField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_test_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesPodContainerField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesPodContainerField) ProtoMessage() {}

func (x *KubernetesPodContainerField) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_test_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesPodContainerField.ProtoReflect.Descriptor instead.
func (*KubernetesPodContainerField) Descriptor() ([]byte, []int) {
	return file_testdata_proto_test_proto_rawDescGZIP(), []int{0, 3, 1}
}

func (x *KubernetesPodContainerField) GetCpuUsageCoreNanoseconds() int64 {
	if x != nil {
		return x.CpuUsageCoreNanoseconds
	}
	return 0
}

func (x *KubernetesPodContainerField) GetMemoryUsageBytes() int64 {
	if x != nil {
		return x.MemoryUsageBytes
	}
	return 0
}

var File_testdata_proto_test_proto protoreflect.FileDescriptor

var file_testdata_proto_test_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x06, 0x0a, 0x0a,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x1a, 0xa6, 0x01, 0x0a, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x1a, 0x22, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x7a, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x3b, 0x0a, 0x1a, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x72, 0x65, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x34, 0x0a,
	0x16, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x1a, 0xab, 0x01, 0x0a, 0x0b, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x1a, 0x5b, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x1a, 0x3f, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x78, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x78, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x78, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x78, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x1a, 0x82, 0x02, 0x0a, 0x0a, 0x70, 0x6f, 0x64, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x1a, 0x7c, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x76,
	0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x1a, 0xce, 0x01, 0x0a, 0x0d, 0x70, 0x6f, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x1a, 0x49, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x72, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3b, 0x0a, 0x1a,
	0x63, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x17, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x4e, 0x61,
	0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_testdata_proto_test_proto_rawDescOnce sync.Once
	file_testdata_proto_test_proto_rawDescData = file_testdata_proto_test_proto_rawDesc
)

func file_testdata_proto_test_proto_rawDescGZIP() []byte {
	file_testdata_proto_test_proto_rawDescOnce.Do(func() {
		file_testdata_proto_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_proto_test_proto_rawDescData)
	})
	return file_testdata_proto_test_proto_rawDescData
}

var file_testdata_proto_test_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_testdata_proto_test_proto_goTypes = []interface{}{
	(*Kubernetes)(nil),                  // 0: kubernetes
	(*KubernetesNode)(nil),              // 1: kubernetes.node
	(*KubernetesPodNetwork)(nil),        // 2: kubernetes.pod_network
	(*KubernetesPodVolume)(nil),         // 3: kubernetes.pod_volume
	(*KubernetesPodContainer)(nil),      // 4: kubernetes.pod_container
	(*KubernetesNodeTag)(nil),           // 5: kubernetes.node.tag
	(*KubernetesNodeField)(nil),         // 6: kubernetes.node.field
	(*KubernetesPodNetworkTag)(nil),     // 7: kubernetes.pod_network.tag
	(*KubernetesPodNetworkField)(nil),   // 8: kubernetes.pod_network.field
	(*KubernetesPodVolumeTag)(nil),      // 9: kubernetes.pod_volume.tag
	(*KubernetesPodVolumeField)(nil),    // 10: kubernetes.pod_volume.field
	(*KubernetesPodContainerTag)(nil),   // 11: kubernetes.pod_container.tag
	(*KubernetesPodContainerField)(nil), // 12: kubernetes.pod_container.field
}
var file_testdata_proto_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testdata_proto_test_proto_init() }
func file_testdata_proto_test_proto_init() {
	if File_testdata_proto_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_proto_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kubernetes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodNetwork); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodVolume); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodContainer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesNodeTag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesNodeField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodNetworkTag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodNetworkField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodVolumeTag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodVolumeField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodContainerTag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_testdata_proto_test_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesPodContainerField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testdata_proto_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_proto_test_proto_goTypes,
		DependencyIndexes: file_testdata_proto_test_proto_depIdxs,
		MessageInfos:      file_testdata_proto_test_proto_msgTypes,
	}.Build()
	File_testdata_proto_test_proto = out.File
	file_testdata_proto_test_proto_rawDesc = nil
	file_testdata_proto_test_proto_goTypes = nil
	file_testdata_proto_test_proto_depIdxs = nil
}
