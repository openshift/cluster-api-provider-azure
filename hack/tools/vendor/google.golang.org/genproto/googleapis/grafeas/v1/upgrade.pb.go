// Copyright 2019 The Grafeas Authors. All rights reserved.
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: grafeas/v1/upgrade.proto

package grafeas

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An Upgrade Note represents a potential upgrade of a package to a given
// version. For each package version combination (i.e. bash 4.0, bash 4.1,
// bash 4.1.2), there will be an Upgrade Note. For Windows, windows_update field
// represents the information related to the update.
type UpgradeNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required for non-Windows OS. The package this Upgrade is for.
	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	// Required for non-Windows OS. The version of the package in machine + human
	// readable form.
	Version *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Metadata about the upgrade for each specific operating system.
	Distributions []*UpgradeDistribution `protobuf:"bytes,3,rep,name=distributions,proto3" json:"distributions,omitempty"`
	// Required for Windows OS. Represents the metadata about the Windows update.
	WindowsUpdate *WindowsUpdate `protobuf:"bytes,4,opt,name=windows_update,json=windowsUpdate,proto3" json:"windows_update,omitempty"`
}

func (x *UpgradeNote) Reset() {
	*x = UpgradeNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grafeas_v1_upgrade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeNote) ProtoMessage() {}

func (x *UpgradeNote) ProtoReflect() protoreflect.Message {
	mi := &file_grafeas_v1_upgrade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeNote.ProtoReflect.Descriptor instead.
func (*UpgradeNote) Descriptor() ([]byte, []int) {
	return file_grafeas_v1_upgrade_proto_rawDescGZIP(), []int{0}
}

func (x *UpgradeNote) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *UpgradeNote) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *UpgradeNote) GetDistributions() []*UpgradeDistribution {
	if x != nil {
		return x.Distributions
	}
	return nil
}

func (x *UpgradeNote) GetWindowsUpdate() *WindowsUpdate {
	if x != nil {
		return x.WindowsUpdate
	}
	return nil
}

// The Upgrade Distribution represents metadata about the Upgrade for each
// operating system (CPE). Some distributions have additional metadata around
// updates, classifying them into various categories and severities.
type UpgradeDistribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required - The specific operating system this metadata applies to. See
	// https://cpe.mitre.org/specification/.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The operating system classification of this Upgrade, as specified by the
	// upstream operating system upgrade feed. For Windows the classification is
	// one of the category_ids listed at
	// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/ff357803(v=vs.85)
	Classification string `protobuf:"bytes,2,opt,name=classification,proto3" json:"classification,omitempty"`
	// The severity as specified by the upstream operating system.
	Severity string `protobuf:"bytes,3,opt,name=severity,proto3" json:"severity,omitempty"`
	// The cve tied to this Upgrade.
	Cve []string `protobuf:"bytes,4,rep,name=cve,proto3" json:"cve,omitempty"`
}

func (x *UpgradeDistribution) Reset() {
	*x = UpgradeDistribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grafeas_v1_upgrade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeDistribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeDistribution) ProtoMessage() {}

func (x *UpgradeDistribution) ProtoReflect() protoreflect.Message {
	mi := &file_grafeas_v1_upgrade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeDistribution.ProtoReflect.Descriptor instead.
func (*UpgradeDistribution) Descriptor() ([]byte, []int) {
	return file_grafeas_v1_upgrade_proto_rawDescGZIP(), []int{1}
}

func (x *UpgradeDistribution) GetCpeUri() string {
	if x != nil {
		return x.CpeUri
	}
	return ""
}

func (x *UpgradeDistribution) GetClassification() string {
	if x != nil {
		return x.Classification
	}
	return ""
}

func (x *UpgradeDistribution) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *UpgradeDistribution) GetCve() []string {
	if x != nil {
		return x.Cve
	}
	return nil
}

// Windows Update represents the metadata about the update for the Windows
// operating system. The fields in this message come from the Windows Update API
// documented at
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdate.
type WindowsUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required - The unique identifier for the update.
	Identity *WindowsUpdate_Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// The localized title of the update.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// The localized description of the update.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The list of categories to which the update belongs.
	Categories []*WindowsUpdate_Category `protobuf:"bytes,4,rep,name=categories,proto3" json:"categories,omitempty"`
	// The Microsoft Knowledge Base article IDs that are associated with the
	// update.
	KbArticleIds []string `protobuf:"bytes,5,rep,name=kb_article_ids,json=kbArticleIds,proto3" json:"kb_article_ids,omitempty"`
	// The hyperlink to the support information for the update.
	SupportUrl string `protobuf:"bytes,6,opt,name=support_url,json=supportUrl,proto3" json:"support_url,omitempty"`
	// The last published timestamp of the update.
	LastPublishedTimestamp *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_published_timestamp,json=lastPublishedTimestamp,proto3" json:"last_published_timestamp,omitempty"`
}

func (x *WindowsUpdate) Reset() {
	*x = WindowsUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grafeas_v1_upgrade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindowsUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindowsUpdate) ProtoMessage() {}

func (x *WindowsUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_grafeas_v1_upgrade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindowsUpdate.ProtoReflect.Descriptor instead.
func (*WindowsUpdate) Descriptor() ([]byte, []int) {
	return file_grafeas_v1_upgrade_proto_rawDescGZIP(), []int{2}
}

func (x *WindowsUpdate) GetIdentity() *WindowsUpdate_Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *WindowsUpdate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WindowsUpdate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WindowsUpdate) GetCategories() []*WindowsUpdate_Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *WindowsUpdate) GetKbArticleIds() []string {
	if x != nil {
		return x.KbArticleIds
	}
	return nil
}

func (x *WindowsUpdate) GetSupportUrl() string {
	if x != nil {
		return x.SupportUrl
	}
	return ""
}

func (x *WindowsUpdate) GetLastPublishedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastPublishedTimestamp
	}
	return nil
}

// An Upgrade Occurrence represents that a specific resource_url could install a
// specific upgrade. This presence is supplied via local sources (i.e. it is
// present in the mirror and the running system has noticed its availability).
// For Windows, both distribution and windows_update contain information for the
// Windows update.
type UpgradeOccurrence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required for non-Windows OS. The package this Upgrade is for.
	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	// Required for non-Windows OS. The version of the package in a machine +
	// human readable form.
	ParsedVersion *Version `protobuf:"bytes,3,opt,name=parsed_version,json=parsedVersion,proto3" json:"parsed_version,omitempty"`
	// Metadata about the upgrade for available for the specific operating system
	// for the resource_url. This allows efficient filtering, as well as
	// making it easier to use the occurrence.
	Distribution *UpgradeDistribution `protobuf:"bytes,4,opt,name=distribution,proto3" json:"distribution,omitempty"`
	// Required for Windows OS. Represents the metadata about the Windows update.
	WindowsUpdate *WindowsUpdate `protobuf:"bytes,5,opt,name=windows_update,json=windowsUpdate,proto3" json:"windows_update,omitempty"`
}

func (x *UpgradeOccurrence) Reset() {
	*x = UpgradeOccurrence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grafeas_v1_upgrade_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeOccurrence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeOccurrence) ProtoMessage() {}

func (x *UpgradeOccurrence) ProtoReflect() protoreflect.Message {
	mi := &file_grafeas_v1_upgrade_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeOccurrence.ProtoReflect.Descriptor instead.
func (*UpgradeOccurrence) Descriptor() ([]byte, []int) {
	return file_grafeas_v1_upgrade_proto_rawDescGZIP(), []int{3}
}

func (x *UpgradeOccurrence) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *UpgradeOccurrence) GetParsedVersion() *Version {
	if x != nil {
		return x.ParsedVersion
	}
	return nil
}

func (x *UpgradeOccurrence) GetDistribution() *UpgradeDistribution {
	if x != nil {
		return x.Distribution
	}
	return nil
}

func (x *UpgradeOccurrence) GetWindowsUpdate() *WindowsUpdate {
	if x != nil {
		return x.WindowsUpdate
	}
	return nil
}

// The unique identifier of the update.
type WindowsUpdate_Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The revision independent identifier of the update.
	UpdateId string `protobuf:"bytes,1,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
	// The revision number of the update.
	Revision int32 `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *WindowsUpdate_Identity) Reset() {
	*x = WindowsUpdate_Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grafeas_v1_upgrade_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindowsUpdate_Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindowsUpdate_Identity) ProtoMessage() {}

func (x *WindowsUpdate_Identity) ProtoReflect() protoreflect.Message {
	mi := &file_grafeas_v1_upgrade_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindowsUpdate_Identity.ProtoReflect.Descriptor instead.
func (*WindowsUpdate_Identity) Descriptor() ([]byte, []int) {
	return file_grafeas_v1_upgrade_proto_rawDescGZIP(), []int{2, 0}
}

func (x *WindowsUpdate_Identity) GetUpdateId() string {
	if x != nil {
		return x.UpdateId
	}
	return ""
}

func (x *WindowsUpdate_Identity) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

// The category to which the update belongs.
type WindowsUpdate_Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the category.
	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	// The localized name of the category.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WindowsUpdate_Category) Reset() {
	*x = WindowsUpdate_Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grafeas_v1_upgrade_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindowsUpdate_Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindowsUpdate_Category) ProtoMessage() {}

func (x *WindowsUpdate_Category) ProtoReflect() protoreflect.Message {
	mi := &file_grafeas_v1_upgrade_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindowsUpdate_Category.ProtoReflect.Descriptor instead.
func (*WindowsUpdate_Category) Descriptor() ([]byte, []int) {
	return file_grafeas_v1_upgrade_proto_rawDescGZIP(), []int{2, 1}
}

func (x *WindowsUpdate_Category) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *WindowsUpdate_Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_grafeas_v1_upgrade_proto protoreflect.FileDescriptor

var file_grafeas_v1_upgrade_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0d, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x40, 0x0a, 0x0e, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x70, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x70,
	0x65, 0x55, 0x72, 0x69, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x65, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x65, 0x22, 0xee, 0x03, 0x0a, 0x0d, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6b, 0x62, 0x5f, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x6b, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x54, 0x0a, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x6c,
	0x61, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x43, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x3f, 0x0a, 0x08, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf0, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x0d, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x51,
	0x0a, 0x0d, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x52,
	0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grafeas_v1_upgrade_proto_rawDescOnce sync.Once
	file_grafeas_v1_upgrade_proto_rawDescData = file_grafeas_v1_upgrade_proto_rawDesc
)

func file_grafeas_v1_upgrade_proto_rawDescGZIP() []byte {
	file_grafeas_v1_upgrade_proto_rawDescOnce.Do(func() {
		file_grafeas_v1_upgrade_proto_rawDescData = protoimpl.X.CompressGZIP(file_grafeas_v1_upgrade_proto_rawDescData)
	})
	return file_grafeas_v1_upgrade_proto_rawDescData
}

var file_grafeas_v1_upgrade_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grafeas_v1_upgrade_proto_goTypes = []interface{}{
	(*UpgradeNote)(nil),            // 0: grafeas.v1.UpgradeNote
	(*UpgradeDistribution)(nil),    // 1: grafeas.v1.UpgradeDistribution
	(*WindowsUpdate)(nil),          // 2: grafeas.v1.WindowsUpdate
	(*UpgradeOccurrence)(nil),      // 3: grafeas.v1.UpgradeOccurrence
	(*WindowsUpdate_Identity)(nil), // 4: grafeas.v1.WindowsUpdate.Identity
	(*WindowsUpdate_Category)(nil), // 5: grafeas.v1.WindowsUpdate.Category
	(*Version)(nil),                // 6: grafeas.v1.Version
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
}
var file_grafeas_v1_upgrade_proto_depIdxs = []int32{
	6, // 0: grafeas.v1.UpgradeNote.version:type_name -> grafeas.v1.Version
	1, // 1: grafeas.v1.UpgradeNote.distributions:type_name -> grafeas.v1.UpgradeDistribution
	2, // 2: grafeas.v1.UpgradeNote.windows_update:type_name -> grafeas.v1.WindowsUpdate
	4, // 3: grafeas.v1.WindowsUpdate.identity:type_name -> grafeas.v1.WindowsUpdate.Identity
	5, // 4: grafeas.v1.WindowsUpdate.categories:type_name -> grafeas.v1.WindowsUpdate.Category
	7, // 5: grafeas.v1.WindowsUpdate.last_published_timestamp:type_name -> google.protobuf.Timestamp
	6, // 6: grafeas.v1.UpgradeOccurrence.parsed_version:type_name -> grafeas.v1.Version
	1, // 7: grafeas.v1.UpgradeOccurrence.distribution:type_name -> grafeas.v1.UpgradeDistribution
	2, // 8: grafeas.v1.UpgradeOccurrence.windows_update:type_name -> grafeas.v1.WindowsUpdate
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_grafeas_v1_upgrade_proto_init() }
func file_grafeas_v1_upgrade_proto_init() {
	if File_grafeas_v1_upgrade_proto != nil {
		return
	}
	file_grafeas_v1_package_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grafeas_v1_upgrade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeNote); i {
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
		file_grafeas_v1_upgrade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeDistribution); i {
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
		file_grafeas_v1_upgrade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindowsUpdate); i {
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
		file_grafeas_v1_upgrade_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeOccurrence); i {
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
		file_grafeas_v1_upgrade_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindowsUpdate_Identity); i {
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
		file_grafeas_v1_upgrade_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindowsUpdate_Category); i {
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
			RawDescriptor: file_grafeas_v1_upgrade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grafeas_v1_upgrade_proto_goTypes,
		DependencyIndexes: file_grafeas_v1_upgrade_proto_depIdxs,
		MessageInfos:      file_grafeas_v1_upgrade_proto_msgTypes,
	}.Build()
	File_grafeas_v1_upgrade_proto = out.File
	file_grafeas_v1_upgrade_proto_rawDesc = nil
	file_grafeas_v1_upgrade_proto_goTypes = nil
	file_grafeas_v1_upgrade_proto_depIdxs = nil
}
