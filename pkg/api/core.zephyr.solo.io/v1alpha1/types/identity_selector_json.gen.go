// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/core/v1alpha1/identity_selector.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for IdentitySelector
func (this *IdentitySelector) MarshalJSON() ([]byte, error) {
	str, err := IdentitySelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IdentitySelector
func (this *IdentitySelector) UnmarshalJSON(b []byte) error {
	return IdentitySelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IdentitySelector_Matcher
func (this *IdentitySelector_Matcher) MarshalJSON() ([]byte, error) {
	str, err := IdentitySelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IdentitySelector_Matcher
func (this *IdentitySelector_Matcher) UnmarshalJSON(b []byte) error {
	return IdentitySelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IdentitySelector_ServiceAccountRefs
func (this *IdentitySelector_ServiceAccountRefs) MarshalJSON() ([]byte, error) {
	str, err := IdentitySelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IdentitySelector_ServiceAccountRefs
func (this *IdentitySelector_ServiceAccountRefs) UnmarshalJSON(b []byte) error {
	return IdentitySelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	IdentitySelectorMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	IdentitySelectorUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
