// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: contrib/envoy/extensions/filters/network/kafka_mesh/v3alpha/kafka_mesh.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type KafkaMesh_ConsumerProxyMode int32

const (
	// Records received are going to be distributed amongst downstream consumer connections.
	// In this mode Envoy uses librdkafka consumers pointing at upstream Kafka clusters, what means that these
	// consumers' position is meaningful and affects what records are received from upstream.
	// Users might want to take a look into these consumers' custom configuration to manage their auto-committing
	// capabilities, as it will impact Envoy's behaviour in case of restarts.
	KafkaMesh_StatefulConsumerProxy KafkaMesh_ConsumerProxyMode = 0
)

// Enum value maps for KafkaMesh_ConsumerProxyMode.
var (
	KafkaMesh_ConsumerProxyMode_name = map[int32]string{
		0: "StatefulConsumerProxy",
	}
	KafkaMesh_ConsumerProxyMode_value = map[string]int32{
		"StatefulConsumerProxy": 0,
	}
)

func (x KafkaMesh_ConsumerProxyMode) Enum() *KafkaMesh_ConsumerProxyMode {
	p := new(KafkaMesh_ConsumerProxyMode)
	*p = x
	return p
}

func (x KafkaMesh_ConsumerProxyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KafkaMesh_ConsumerProxyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_enumTypes[0].Descriptor()
}

func (KafkaMesh_ConsumerProxyMode) Type() protoreflect.EnumType {
	return &file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_enumTypes[0]
}

func (x KafkaMesh_ConsumerProxyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KafkaMesh_ConsumerProxyMode.Descriptor instead.
func (KafkaMesh_ConsumerProxyMode) EnumDescriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescGZIP(), []int{0, 0}
}

// [#next-free-field: 6]
type KafkaMesh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Envoy's host that's advertised to clients.
	// Has the same meaning as corresponding Kafka broker properties.
	// Usually equal to filter chain's listener config, but needs to be reachable by clients
	// (so 0.0.0.0 will not work).
	AdvertisedHost string `protobuf:"bytes,1,opt,name=advertised_host,json=advertisedHost,proto3" json:"advertised_host,omitempty"`
	// Envoy's port that's advertised to clients.
	AdvertisedPort int32 `protobuf:"varint,2,opt,name=advertised_port,json=advertisedPort,proto3" json:"advertised_port,omitempty"`
	// Upstream clusters this filter will connect to.
	UpstreamClusters []*KafkaClusterDefinition `protobuf:"bytes,3,rep,name=upstream_clusters,json=upstreamClusters,proto3" json:"upstream_clusters,omitempty"`
	// Rules that will decide which cluster gets which request.
	ForwardingRules []*ForwardingRule `protobuf:"bytes,4,rep,name=forwarding_rules,json=forwardingRules,proto3" json:"forwarding_rules,omitempty"`
	// How the consumer proxying should behave - this relates mostly to Fetch request handling.
	ConsumerProxyMode KafkaMesh_ConsumerProxyMode `protobuf:"varint,5,opt,name=consumer_proxy_mode,json=consumerProxyMode,proto3,enum=envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaMesh_ConsumerProxyMode" json:"consumer_proxy_mode,omitempty"`
}

func (x *KafkaMesh) Reset() {
	*x = KafkaMesh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaMesh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaMesh) ProtoMessage() {}

func (x *KafkaMesh) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaMesh.ProtoReflect.Descriptor instead.
func (*KafkaMesh) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescGZIP(), []int{0}
}

func (x *KafkaMesh) GetAdvertisedHost() string {
	if x != nil {
		return x.AdvertisedHost
	}
	return ""
}

func (x *KafkaMesh) GetAdvertisedPort() int32 {
	if x != nil {
		return x.AdvertisedPort
	}
	return 0
}

func (x *KafkaMesh) GetUpstreamClusters() []*KafkaClusterDefinition {
	if x != nil {
		return x.UpstreamClusters
	}
	return nil
}

func (x *KafkaMesh) GetForwardingRules() []*ForwardingRule {
	if x != nil {
		return x.ForwardingRules
	}
	return nil
}

func (x *KafkaMesh) GetConsumerProxyMode() KafkaMesh_ConsumerProxyMode {
	if x != nil {
		return x.ConsumerProxyMode
	}
	return KafkaMesh_StatefulConsumerProxy
}

// [#next-free-field: 6]
type KafkaClusterDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster name.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Kafka cluster address.
	BootstrapServers string `protobuf:"bytes,2,opt,name=bootstrap_servers,json=bootstrapServers,proto3" json:"bootstrap_servers,omitempty"`
	// Default number of partitions present in this cluster.
	// This is especially important for clients that do not specify partition in their payloads and depend on this value for hashing.
	// The same number of partitions is going to be used by upstream-pointing Kafka consumers for consumer proxying scenarios.
	PartitionCount int32 `protobuf:"varint,3,opt,name=partition_count,json=partitionCount,proto3" json:"partition_count,omitempty"`
	// Custom configuration passed to Kafka producer.
	ProducerConfig map[string]string `protobuf:"bytes,4,rep,name=producer_config,json=producerConfig,proto3" json:"producer_config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Custom configuration passed to Kafka consumer.
	ConsumerConfig map[string]string `protobuf:"bytes,5,rep,name=consumer_config,json=consumerConfig,proto3" json:"consumer_config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *KafkaClusterDefinition) Reset() {
	*x = KafkaClusterDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaClusterDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaClusterDefinition) ProtoMessage() {}

func (x *KafkaClusterDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaClusterDefinition.ProtoReflect.Descriptor instead.
func (*KafkaClusterDefinition) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescGZIP(), []int{1}
}

func (x *KafkaClusterDefinition) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *KafkaClusterDefinition) GetBootstrapServers() string {
	if x != nil {
		return x.BootstrapServers
	}
	return ""
}

func (x *KafkaClusterDefinition) GetPartitionCount() int32 {
	if x != nil {
		return x.PartitionCount
	}
	return 0
}

func (x *KafkaClusterDefinition) GetProducerConfig() map[string]string {
	if x != nil {
		return x.ProducerConfig
	}
	return nil
}

func (x *KafkaClusterDefinition) GetConsumerConfig() map[string]string {
	if x != nil {
		return x.ConsumerConfig
	}
	return nil
}

type ForwardingRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster name.
	TargetCluster string `protobuf:"bytes,1,opt,name=target_cluster,json=targetCluster,proto3" json:"target_cluster,omitempty"`
	// Types that are assignable to Trigger:
	//
	//	*ForwardingRule_TopicPrefix
	Trigger isForwardingRule_Trigger `protobuf_oneof:"trigger"`
}

func (x *ForwardingRule) Reset() {
	*x = ForwardingRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardingRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardingRule) ProtoMessage() {}

func (x *ForwardingRule) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardingRule.ProtoReflect.Descriptor instead.
func (*ForwardingRule) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescGZIP(), []int{2}
}

func (x *ForwardingRule) GetTargetCluster() string {
	if x != nil {
		return x.TargetCluster
	}
	return ""
}

func (m *ForwardingRule) GetTrigger() isForwardingRule_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (x *ForwardingRule) GetTopicPrefix() string {
	if x, ok := x.GetTrigger().(*ForwardingRule_TopicPrefix); ok {
		return x.TopicPrefix
	}
	return ""
}

type isForwardingRule_Trigger interface {
	isForwardingRule_Trigger()
}

type ForwardingRule_TopicPrefix struct {
	// Intended place for future types of forwarding rules.
	TopicPrefix string `protobuf:"bytes,2,opt,name=topic_prefix,json=topicPrefix,proto3,oneof"`
}

func (*ForwardingRule_TopicPrefix) isForwardingRule_Trigger() {}

var File_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x5f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x04, 0x0a,
	0x09, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x68, 0x12, 0x30, 0x0a, 0x0f, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0f,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x0e,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x78,
	0x0a, 0x11, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x6e, 0x0a, 0x10, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x50, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x11, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x10, 0x00, 0x22, 0xc8, 0x04, 0x0a, 0x16,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x88, 0x01, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x5f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x88, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x5f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x33,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x41, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x67, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0c, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x42, 0x09, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x42,
	0xc1, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02,
	0x08, 0x01, 0x0a, 0x41, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x33,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x68,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescData = file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDesc
)

func file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDescData
}

var file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_goTypes = []interface{}{
	(KafkaMesh_ConsumerProxyMode)(0), // 0: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaMesh.ConsumerProxyMode
	(*KafkaMesh)(nil),                // 1: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaMesh
	(*KafkaClusterDefinition)(nil),   // 2: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaClusterDefinition
	(*ForwardingRule)(nil),           // 3: envoy.extensions.filters.network.kafka_mesh.v3alpha.ForwardingRule
	nil,                              // 4: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaClusterDefinition.ProducerConfigEntry
	nil,                              // 5: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaClusterDefinition.ConsumerConfigEntry
}
var file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaMesh.upstream_clusters:type_name -> envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaClusterDefinition
	3, // 1: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaMesh.forwarding_rules:type_name -> envoy.extensions.filters.network.kafka_mesh.v3alpha.ForwardingRule
	0, // 2: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaMesh.consumer_proxy_mode:type_name -> envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaMesh.ConsumerProxyMode
	4, // 3: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaClusterDefinition.producer_config:type_name -> envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaClusterDefinition.ProducerConfigEntry
	5, // 4: envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaClusterDefinition.consumer_config:type_name -> envoy.extensions.filters.network.kafka_mesh.v3alpha.KafkaClusterDefinition.ConsumerConfigEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_init() }
func file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_init() {
	if File_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaMesh); i {
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
		file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaClusterDefinition); i {
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
		file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardingRule); i {
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
	file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ForwardingRule_TopicPrefix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_depIdxs,
		EnumInfos:         file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_enumTypes,
		MessageInfos:      file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto = out.File
	file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_rawDesc = nil
	file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_goTypes = nil
	file_contrib_envoy_extensions_filters_network_kafka_mesh_v3alpha_kafka_mesh_proto_depIdxs = nil
}
