// Copyright 2018 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: cvss.proto

package cvss_go_proto

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

// CVSS Version.
type CVSSVersion int32

const (
	CVSSVersion_CVSS_VERSION_UNSPECIFIED CVSSVersion = 0
	CVSSVersion_CVSS_VERSION_2           CVSSVersion = 2
	CVSSVersion_CVSS_VERSION_3           CVSSVersion = 3
)

// Enum value maps for CVSSVersion.
var (
	CVSSVersion_name = map[int32]string{
		0: "CVSS_VERSION_UNSPECIFIED",
		2: "CVSS_VERSION_2",
		3: "CVSS_VERSION_3",
	}
	CVSSVersion_value = map[string]int32{
		"CVSS_VERSION_UNSPECIFIED": 0,
		"CVSS_VERSION_2":           2,
		"CVSS_VERSION_3":           3,
	}
)

func (x CVSSVersion) Enum() *CVSSVersion {
	p := new(CVSSVersion)
	*p = x
	return p
}

func (x CVSSVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSSVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_cvss_proto_enumTypes[0].Descriptor()
}

func (CVSSVersion) Type() protoreflect.EnumType {
	return &file_cvss_proto_enumTypes[0]
}

func (x CVSSVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSSVersion.Descriptor instead.
func (CVSSVersion) EnumDescriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0}
}

type CVSS_AttackVector int32

const (
	CVSS_ATTACK_VECTOR_UNSPECIFIED CVSS_AttackVector = 0 // Defined in CVSS v3, CVSS v2
	CVSS_ATTACK_VECTOR_NETWORK     CVSS_AttackVector = 1 // Defined in CVSS v3, CVSS v2
	CVSS_ATTACK_VECTOR_ADJACENT    CVSS_AttackVector = 2 // Defined in CVSS v3, CVSS v2
	CVSS_ATTACK_VECTOR_LOCAL       CVSS_AttackVector = 3 // Defined in CVSS v3, CVSS v2
	CVSS_ATTACK_VECTOR_PHYSICAL    CVSS_AttackVector = 4 // Defined in CVSS v3
)

// Enum value maps for CVSS_AttackVector.
var (
	CVSS_AttackVector_name = map[int32]string{
		0: "ATTACK_VECTOR_UNSPECIFIED",
		1: "ATTACK_VECTOR_NETWORK",
		2: "ATTACK_VECTOR_ADJACENT",
		3: "ATTACK_VECTOR_LOCAL",
		4: "ATTACK_VECTOR_PHYSICAL",
	}
	CVSS_AttackVector_value = map[string]int32{
		"ATTACK_VECTOR_UNSPECIFIED": 0,
		"ATTACK_VECTOR_NETWORK":     1,
		"ATTACK_VECTOR_ADJACENT":    2,
		"ATTACK_VECTOR_LOCAL":       3,
		"ATTACK_VECTOR_PHYSICAL":    4,
	}
)

func (x CVSS_AttackVector) Enum() *CVSS_AttackVector {
	p := new(CVSS_AttackVector)
	*p = x
	return p
}

func (x CVSS_AttackVector) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSS_AttackVector) Descriptor() protoreflect.EnumDescriptor {
	return file_cvss_proto_enumTypes[1].Descriptor()
}

func (CVSS_AttackVector) Type() protoreflect.EnumType {
	return &file_cvss_proto_enumTypes[1]
}

func (x CVSS_AttackVector) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSS_AttackVector.Descriptor instead.
func (CVSS_AttackVector) EnumDescriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0, 0}
}

type CVSS_AttackComplexity int32

const (
	CVSS_ATTACK_COMPLEXITY_UNSPECIFIED CVSS_AttackComplexity = 0 // Defined in CVSS v3, CVSS v2
	CVSS_ATTACK_COMPLEXITY_LOW         CVSS_AttackComplexity = 1 // Defined in CVSS v3, CVSS v2
	CVSS_ATTACK_COMPLEXITY_HIGH        CVSS_AttackComplexity = 2 // Defined in CVSS v3, CVSS v2
	CVSS_ATTACK_COMPLEXITY_MEDIUM      CVSS_AttackComplexity = 3 // Defined in CVSS v2
)

// Enum value maps for CVSS_AttackComplexity.
var (
	CVSS_AttackComplexity_name = map[int32]string{
		0: "ATTACK_COMPLEXITY_UNSPECIFIED",
		1: "ATTACK_COMPLEXITY_LOW",
		2: "ATTACK_COMPLEXITY_HIGH",
		3: "ATTACK_COMPLEXITY_MEDIUM",
	}
	CVSS_AttackComplexity_value = map[string]int32{
		"ATTACK_COMPLEXITY_UNSPECIFIED": 0,
		"ATTACK_COMPLEXITY_LOW":         1,
		"ATTACK_COMPLEXITY_HIGH":        2,
		"ATTACK_COMPLEXITY_MEDIUM":      3,
	}
)

func (x CVSS_AttackComplexity) Enum() *CVSS_AttackComplexity {
	p := new(CVSS_AttackComplexity)
	*p = x
	return p
}

func (x CVSS_AttackComplexity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSS_AttackComplexity) Descriptor() protoreflect.EnumDescriptor {
	return file_cvss_proto_enumTypes[2].Descriptor()
}

func (CVSS_AttackComplexity) Type() protoreflect.EnumType {
	return &file_cvss_proto_enumTypes[2]
}

func (x CVSS_AttackComplexity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSS_AttackComplexity.Descriptor instead.
func (CVSS_AttackComplexity) EnumDescriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0, 1}
}

type CVSS_Authentication int32

const (
	CVSS_AUTHENTICATION_UNSPECIFIED CVSS_Authentication = 0 // Defined in CVSS v2
	CVSS_AUTHENTICATION_MULTIPLE    CVSS_Authentication = 1 // Defined in CVSS v2
	CVSS_AUTHENTICATION_SINGLE      CVSS_Authentication = 2 // Defined in CVSS v2
	CVSS_AUTHENTICATION_NONE        CVSS_Authentication = 3 // Defined in CVSS v2
)

// Enum value maps for CVSS_Authentication.
var (
	CVSS_Authentication_name = map[int32]string{
		0: "AUTHENTICATION_UNSPECIFIED",
		1: "AUTHENTICATION_MULTIPLE",
		2: "AUTHENTICATION_SINGLE",
		3: "AUTHENTICATION_NONE",
	}
	CVSS_Authentication_value = map[string]int32{
		"AUTHENTICATION_UNSPECIFIED": 0,
		"AUTHENTICATION_MULTIPLE":    1,
		"AUTHENTICATION_SINGLE":      2,
		"AUTHENTICATION_NONE":        3,
	}
)

func (x CVSS_Authentication) Enum() *CVSS_Authentication {
	p := new(CVSS_Authentication)
	*p = x
	return p
}

func (x CVSS_Authentication) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSS_Authentication) Descriptor() protoreflect.EnumDescriptor {
	return file_cvss_proto_enumTypes[3].Descriptor()
}

func (CVSS_Authentication) Type() protoreflect.EnumType {
	return &file_cvss_proto_enumTypes[3]
}

func (x CVSS_Authentication) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSS_Authentication.Descriptor instead.
func (CVSS_Authentication) EnumDescriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0, 2}
}

type CVSS_PrivilegesRequired int32

const (
	CVSS_PRIVILEGES_REQUIRED_UNSPECIFIED CVSS_PrivilegesRequired = 0 // Defined in CVSS v3
	CVSS_PRIVILEGES_REQUIRED_NONE        CVSS_PrivilegesRequired = 1 // Defined in CVSS v3
	CVSS_PRIVILEGES_REQUIRED_LOW         CVSS_PrivilegesRequired = 2 // Defined in CVSS v3
	CVSS_PRIVILEGES_REQUIRED_HIGH        CVSS_PrivilegesRequired = 3 // Defined in CVSS v3
)

// Enum value maps for CVSS_PrivilegesRequired.
var (
	CVSS_PrivilegesRequired_name = map[int32]string{
		0: "PRIVILEGES_REQUIRED_UNSPECIFIED",
		1: "PRIVILEGES_REQUIRED_NONE",
		2: "PRIVILEGES_REQUIRED_LOW",
		3: "PRIVILEGES_REQUIRED_HIGH",
	}
	CVSS_PrivilegesRequired_value = map[string]int32{
		"PRIVILEGES_REQUIRED_UNSPECIFIED": 0,
		"PRIVILEGES_REQUIRED_NONE":        1,
		"PRIVILEGES_REQUIRED_LOW":         2,
		"PRIVILEGES_REQUIRED_HIGH":        3,
	}
)

func (x CVSS_PrivilegesRequired) Enum() *CVSS_PrivilegesRequired {
	p := new(CVSS_PrivilegesRequired)
	*p = x
	return p
}

func (x CVSS_PrivilegesRequired) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSS_PrivilegesRequired) Descriptor() protoreflect.EnumDescriptor {
	return file_cvss_proto_enumTypes[4].Descriptor()
}

func (CVSS_PrivilegesRequired) Type() protoreflect.EnumType {
	return &file_cvss_proto_enumTypes[4]
}

func (x CVSS_PrivilegesRequired) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSS_PrivilegesRequired.Descriptor instead.
func (CVSS_PrivilegesRequired) EnumDescriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0, 3}
}

type CVSS_UserInteraction int32

const (
	CVSS_USER_INTERACTION_UNSPECIFIED CVSS_UserInteraction = 0 // Defined in CVSS v3
	CVSS_USER_INTERACTION_NONE        CVSS_UserInteraction = 1 // Defined in CVSS v3
	CVSS_USER_INTERACTION_REQUIRED    CVSS_UserInteraction = 2 // Defined in CVSS v3
)

// Enum value maps for CVSS_UserInteraction.
var (
	CVSS_UserInteraction_name = map[int32]string{
		0: "USER_INTERACTION_UNSPECIFIED",
		1: "USER_INTERACTION_NONE",
		2: "USER_INTERACTION_REQUIRED",
	}
	CVSS_UserInteraction_value = map[string]int32{
		"USER_INTERACTION_UNSPECIFIED": 0,
		"USER_INTERACTION_NONE":        1,
		"USER_INTERACTION_REQUIRED":    2,
	}
)

func (x CVSS_UserInteraction) Enum() *CVSS_UserInteraction {
	p := new(CVSS_UserInteraction)
	*p = x
	return p
}

func (x CVSS_UserInteraction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSS_UserInteraction) Descriptor() protoreflect.EnumDescriptor {
	return file_cvss_proto_enumTypes[5].Descriptor()
}

func (CVSS_UserInteraction) Type() protoreflect.EnumType {
	return &file_cvss_proto_enumTypes[5]
}

func (x CVSS_UserInteraction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSS_UserInteraction.Descriptor instead.
func (CVSS_UserInteraction) EnumDescriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0, 4}
}

type CVSS_Scope int32

const (
	CVSS_SCOPE_UNSPECIFIED CVSS_Scope = 0 // Defined in CVSS v3
	CVSS_SCOPE_UNCHANGED   CVSS_Scope = 1 // Defined in CVSS v3
	CVSS_SCOPE_CHANGED     CVSS_Scope = 2 // Defined in CVSS v3
)

// Enum value maps for CVSS_Scope.
var (
	CVSS_Scope_name = map[int32]string{
		0: "SCOPE_UNSPECIFIED",
		1: "SCOPE_UNCHANGED",
		2: "SCOPE_CHANGED",
	}
	CVSS_Scope_value = map[string]int32{
		"SCOPE_UNSPECIFIED": 0,
		"SCOPE_UNCHANGED":   1,
		"SCOPE_CHANGED":     2,
	}
)

func (x CVSS_Scope) Enum() *CVSS_Scope {
	p := new(CVSS_Scope)
	*p = x
	return p
}

func (x CVSS_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSS_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_cvss_proto_enumTypes[6].Descriptor()
}

func (CVSS_Scope) Type() protoreflect.EnumType {
	return &file_cvss_proto_enumTypes[6]
}

func (x CVSS_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSS_Scope.Descriptor instead.
func (CVSS_Scope) EnumDescriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0, 5}
}

type CVSS_Impact int32

const (
	CVSS_IMPACT_UNSPECIFIED CVSS_Impact = 0 // Defined in CVSS v3, CVSS v2
	CVSS_IMPACT_HIGH        CVSS_Impact = 1 // Defined in CVSS v3
	CVSS_IMPACT_LOW         CVSS_Impact = 2 // Defined in CVSS v3
	CVSS_IMPACT_NONE        CVSS_Impact = 3 // Defined in CVSS v3, CVSS v2
	CVSS_IMPACT_PARTIAL     CVSS_Impact = 4 // Defined in CVSS v2
	CVSS_IMPACT_COMPLETE    CVSS_Impact = 5 // Defined in CVSS v2
)

// Enum value maps for CVSS_Impact.
var (
	CVSS_Impact_name = map[int32]string{
		0: "IMPACT_UNSPECIFIED",
		1: "IMPACT_HIGH",
		2: "IMPACT_LOW",
		3: "IMPACT_NONE",
		4: "IMPACT_PARTIAL",
		5: "IMPACT_COMPLETE",
	}
	CVSS_Impact_value = map[string]int32{
		"IMPACT_UNSPECIFIED": 0,
		"IMPACT_HIGH":        1,
		"IMPACT_LOW":         2,
		"IMPACT_NONE":        3,
		"IMPACT_PARTIAL":     4,
		"IMPACT_COMPLETE":    5,
	}
)

func (x CVSS_Impact) Enum() *CVSS_Impact {
	p := new(CVSS_Impact)
	*p = x
	return p
}

func (x CVSS_Impact) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSS_Impact) Descriptor() protoreflect.EnumDescriptor {
	return file_cvss_proto_enumTypes[7].Descriptor()
}

func (CVSS_Impact) Type() protoreflect.EnumType {
	return &file_cvss_proto_enumTypes[7]
}

func (x CVSS_Impact) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSS_Impact.Descriptor instead.
func (CVSS_Impact) EnumDescriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0, 6}
}

// Common Vulnerability Scoring System.
// This message is compatible with CVSS v2 and v3.
// For CVSS v2 details, see https://www.first.org/cvss/v2/guide
// CVSS v2 calculator: https://nvd.nist.gov/vuln-metrics/cvss/v2-calculator
// For CVSS v3 details, see https://www.first.org/cvss/specification-document
// CVSS v3 calculator: https://nvd.nist.gov/vuln-metrics/cvss/v3-calculator
type CVSS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base score is a function of the base metric scores.
	BaseScore           float32 `protobuf:"fixed32,1,opt,name=base_score,json=baseScore,proto3" json:"base_score,omitempty"`
	ExploitabilityScore float32 `protobuf:"fixed32,2,opt,name=exploitability_score,json=exploitabilityScore,proto3" json:"exploitability_score,omitempty"`
	ImpactScore         float32 `protobuf:"fixed32,3,opt,name=impact_score,json=impactScore,proto3" json:"impact_score,omitempty"`
	// Base Metrics
	// Represents the intrinsic characteristics of a vulnerability that are
	// constant over time and across user environments.
	AttackVector          CVSS_AttackVector       `protobuf:"varint,4,opt,name=attack_vector,json=attackVector,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_AttackVector" json:"attack_vector,omitempty"`                         // Defined in CVSS v3, CVSS v2
	AttackComplexity      CVSS_AttackComplexity   `protobuf:"varint,5,opt,name=attack_complexity,json=attackComplexity,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_AttackComplexity" json:"attack_complexity,omitempty"`         // Defined in CVSS v3, CVSS v2
	Authentication        CVSS_Authentication     `protobuf:"varint,6,opt,name=authentication,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_Authentication" json:"authentication,omitempty"`                                       // Defined in CVSS v2
	PrivilegesRequired    CVSS_PrivilegesRequired `protobuf:"varint,7,opt,name=privileges_required,json=privilegesRequired,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_PrivilegesRequired" json:"privileges_required,omitempty"` // Defined in CVSS v3
	UserInteraction       CVSS_UserInteraction    `protobuf:"varint,8,opt,name=user_interaction,json=userInteraction,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_UserInteraction" json:"user_interaction,omitempty"`             // Defined in CVSS v3
	Scope                 CVSS_Scope              `protobuf:"varint,9,opt,name=scope,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_Scope" json:"scope,omitempty"`                                                                  // Defined in CVSS v3
	ConfidentialityImpact CVSS_Impact             `protobuf:"varint,10,opt,name=confidentiality_impact,json=confidentialityImpact,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_Impact" json:"confidentiality_impact,omitempty"`   // Defined in CVSS v3, CVSS v2
	IntegrityImpact       CVSS_Impact             `protobuf:"varint,11,opt,name=integrity_impact,json=integrityImpact,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_Impact" json:"integrity_impact,omitempty"`                     // Defined in CVSS v3, CVSS v2
	AvailabilityImpact    CVSS_Impact             `protobuf:"varint,12,opt,name=availability_impact,json=availabilityImpact,proto3,enum=grafeas.v1beta1.vulnerability.CVSS_Impact" json:"availability_impact,omitempty"`            // Defined in CVSS v3, CVSS v2
}

func (x *CVSS) Reset() {
	*x = CVSS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVSS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVSS) ProtoMessage() {}

func (x *CVSS) ProtoReflect() protoreflect.Message {
	mi := &file_cvss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVSS.ProtoReflect.Descriptor instead.
func (*CVSS) Descriptor() ([]byte, []int) {
	return file_cvss_proto_rawDescGZIP(), []int{0}
}

func (x *CVSS) GetBaseScore() float32 {
	if x != nil {
		return x.BaseScore
	}
	return 0
}

func (x *CVSS) GetExploitabilityScore() float32 {
	if x != nil {
		return x.ExploitabilityScore
	}
	return 0
}

func (x *CVSS) GetImpactScore() float32 {
	if x != nil {
		return x.ImpactScore
	}
	return 0
}

func (x *CVSS) GetAttackVector() CVSS_AttackVector {
	if x != nil {
		return x.AttackVector
	}
	return CVSS_ATTACK_VECTOR_UNSPECIFIED
}

func (x *CVSS) GetAttackComplexity() CVSS_AttackComplexity {
	if x != nil {
		return x.AttackComplexity
	}
	return CVSS_ATTACK_COMPLEXITY_UNSPECIFIED
}

func (x *CVSS) GetAuthentication() CVSS_Authentication {
	if x != nil {
		return x.Authentication
	}
	return CVSS_AUTHENTICATION_UNSPECIFIED
}

func (x *CVSS) GetPrivilegesRequired() CVSS_PrivilegesRequired {
	if x != nil {
		return x.PrivilegesRequired
	}
	return CVSS_PRIVILEGES_REQUIRED_UNSPECIFIED
}

func (x *CVSS) GetUserInteraction() CVSS_UserInteraction {
	if x != nil {
		return x.UserInteraction
	}
	return CVSS_USER_INTERACTION_UNSPECIFIED
}

func (x *CVSS) GetScope() CVSS_Scope {
	if x != nil {
		return x.Scope
	}
	return CVSS_SCOPE_UNSPECIFIED
}

func (x *CVSS) GetConfidentialityImpact() CVSS_Impact {
	if x != nil {
		return x.ConfidentialityImpact
	}
	return CVSS_IMPACT_UNSPECIFIED
}

func (x *CVSS) GetIntegrityImpact() CVSS_Impact {
	if x != nil {
		return x.IntegrityImpact
	}
	return CVSS_IMPACT_UNSPECIFIED
}

func (x *CVSS) GetAvailabilityImpact() CVSS_Impact {
	if x != nil {
		return x.AvailabilityImpact
	}
	return CVSS_IMPACT_UNSPECIFIED
}

var File_cvss_proto protoreflect.FileDescriptor

var file_cvss_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x76, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xa8, 0x0e, 0x0a, 0x04,
	0x43, 0x56, 0x53, 0x53, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x69, 0x74, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x13, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x69, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x69, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x55, 0x0a, 0x0d, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x30, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x43, 0x56, 0x53, 0x53, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x61, 0x0a, 0x11, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74,
	0x79, 0x52, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x69, 0x74, 0x79, 0x12, 0x5a, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x67, 0x0a, 0x13, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x73, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53,
	0x53, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x52, 0x12, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x5e, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x2e, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x61, 0x0a, 0x16, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x2e, 0x49,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x55, 0x0a, 0x10,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x2e, 0x49, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x12, 0x5b, 0x0a, 0x13, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2a, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x43, 0x56, 0x53, 0x53, 0x2e, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x12, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x22, 0x99, 0x01, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x56, 0x45, 0x43, 0x54,
	0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x19, 0x0a, 0x15, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x56, 0x45, 0x43, 0x54, 0x4f,
	0x52, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x41,
	0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x56, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x41, 0x44, 0x4a,
	0x41, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x54, 0x54, 0x41, 0x43,
	0x4b, 0x5f, 0x56, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x03,
	0x12, 0x1a, 0x0a, 0x16, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x56, 0x45, 0x43, 0x54, 0x4f,
	0x52, 0x5f, 0x50, 0x48, 0x59, 0x53, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04, 0x22, 0x8a, 0x01, 0x0a,
	0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x69, 0x74,
	0x79, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x58, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12,
	0x1a, 0x0a, 0x16, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x58, 0x49, 0x54, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x41,
	0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x49, 0x54, 0x59,
	0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x03, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x1a,
	0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17,
	0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x55, 0x54,
	0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4e, 0x47,
	0x4c, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x22, 0x92, 0x01,
	0x0a, 0x12, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x52, 0x49, 0x56, 0x49, 0x4c, 0x45, 0x47,
	0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x49,
	0x56, 0x49, 0x4c, 0x45, 0x47, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x49, 0x56, 0x49,
	0x4c, 0x45, 0x47, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x4c,
	0x4f, 0x57, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x49, 0x56, 0x49, 0x4c, 0x45, 0x47,
	0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x48, 0x49, 0x47, 0x48,
	0x10, 0x03, 0x22, 0x6d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x02, 0x22, 0x46, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x43,
	0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x02, 0x22, 0x7b, 0x0a, 0x06, 0x49, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x12, 0x0a,
	0x0e, 0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x10,
	0x04, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x2a, 0x53, 0x0a, 0x0b, 0x43, 0x56, 0x53, 0x53, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x56, 0x53, 0x53, 0x5f, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x56, 0x53, 0x53, 0x5f, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x32, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x56, 0x53, 0x53, 0x5f,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x33, 0x10, 0x03, 0x42, 0x6d, 0x0a, 0x20, 0x69,
	0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42,
	0x09, 0x43, 0x56, 0x53, 0x53, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x76, 0x73, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cvss_proto_rawDescOnce sync.Once
	file_cvss_proto_rawDescData = file_cvss_proto_rawDesc
)

func file_cvss_proto_rawDescGZIP() []byte {
	file_cvss_proto_rawDescOnce.Do(func() {
		file_cvss_proto_rawDescData = protoimpl.X.CompressGZIP(file_cvss_proto_rawDescData)
	})
	return file_cvss_proto_rawDescData
}

var file_cvss_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_cvss_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cvss_proto_goTypes = []interface{}{
	(CVSSVersion)(0),             // 0: grafeas.v1beta1.vulnerability.CVSSVersion
	(CVSS_AttackVector)(0),       // 1: grafeas.v1beta1.vulnerability.CVSS.AttackVector
	(CVSS_AttackComplexity)(0),   // 2: grafeas.v1beta1.vulnerability.CVSS.AttackComplexity
	(CVSS_Authentication)(0),     // 3: grafeas.v1beta1.vulnerability.CVSS.Authentication
	(CVSS_PrivilegesRequired)(0), // 4: grafeas.v1beta1.vulnerability.CVSS.PrivilegesRequired
	(CVSS_UserInteraction)(0),    // 5: grafeas.v1beta1.vulnerability.CVSS.UserInteraction
	(CVSS_Scope)(0),              // 6: grafeas.v1beta1.vulnerability.CVSS.Scope
	(CVSS_Impact)(0),             // 7: grafeas.v1beta1.vulnerability.CVSS.Impact
	(*CVSS)(nil),                 // 8: grafeas.v1beta1.vulnerability.CVSS
}
var file_cvss_proto_depIdxs = []int32{
	1, // 0: grafeas.v1beta1.vulnerability.CVSS.attack_vector:type_name -> grafeas.v1beta1.vulnerability.CVSS.AttackVector
	2, // 1: grafeas.v1beta1.vulnerability.CVSS.attack_complexity:type_name -> grafeas.v1beta1.vulnerability.CVSS.AttackComplexity
	3, // 2: grafeas.v1beta1.vulnerability.CVSS.authentication:type_name -> grafeas.v1beta1.vulnerability.CVSS.Authentication
	4, // 3: grafeas.v1beta1.vulnerability.CVSS.privileges_required:type_name -> grafeas.v1beta1.vulnerability.CVSS.PrivilegesRequired
	5, // 4: grafeas.v1beta1.vulnerability.CVSS.user_interaction:type_name -> grafeas.v1beta1.vulnerability.CVSS.UserInteraction
	6, // 5: grafeas.v1beta1.vulnerability.CVSS.scope:type_name -> grafeas.v1beta1.vulnerability.CVSS.Scope
	7, // 6: grafeas.v1beta1.vulnerability.CVSS.confidentiality_impact:type_name -> grafeas.v1beta1.vulnerability.CVSS.Impact
	7, // 7: grafeas.v1beta1.vulnerability.CVSS.integrity_impact:type_name -> grafeas.v1beta1.vulnerability.CVSS.Impact
	7, // 8: grafeas.v1beta1.vulnerability.CVSS.availability_impact:type_name -> grafeas.v1beta1.vulnerability.CVSS.Impact
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cvss_proto_init() }
func file_cvss_proto_init() {
	if File_cvss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cvss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CVSS); i {
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
			RawDescriptor: file_cvss_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cvss_proto_goTypes,
		DependencyIndexes: file_cvss_proto_depIdxs,
		EnumInfos:         file_cvss_proto_enumTypes,
		MessageInfos:      file_cvss_proto_msgTypes,
	}.Build()
	File_cvss_proto = out.File
	file_cvss_proto_rawDesc = nil
	file_cvss_proto_goTypes = nil
	file_cvss_proto_depIdxs = nil
}
