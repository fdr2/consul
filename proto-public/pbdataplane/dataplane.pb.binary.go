// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: proto-public/pbdataplane/dataplane.proto

package pbdataplane

import (
	"github.com/golang/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *SupportedDataplaneFeaturesRequest) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *SupportedDataplaneFeaturesRequest) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *DataplaneFeatureSupport) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *DataplaneFeatureSupport) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *SupportedDataplaneFeaturesResponse) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *SupportedDataplaneFeaturesResponse) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}