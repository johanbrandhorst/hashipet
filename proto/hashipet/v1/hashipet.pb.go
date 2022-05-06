// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: hashipet/v1/hashipet.proto

package hashipetv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Species int32

const (
	Species_SPECIES_UNSPECIFIED Species = 0
	Species_SPECIES_PIG         Species = 1
	Species_SPECIES_DOG         Species = 2
	Species_SPECIES_CAT         Species = 3
)

// Enum value maps for Species.
var (
	Species_name = map[int32]string{
		0: "SPECIES_UNSPECIFIED",
		1: "SPECIES_PIG",
		2: "SPECIES_DOG",
		3: "SPECIES_CAT",
	}
	Species_value = map[string]int32{
		"SPECIES_UNSPECIFIED": 0,
		"SPECIES_PIG":         1,
		"SPECIES_DOG":         2,
		"SPECIES_CAT":         3,
	}
)

func (x Species) Enum() *Species {
	p := new(Species)
	*p = x
	return p
}

func (x Species) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Species) Descriptor() protoreflect.EnumDescriptor {
	return file_hashipet_v1_hashipet_proto_enumTypes[0].Descriptor()
}

func (Species) Type() protoreflect.EnumType {
	return &file_hashipet_v1_hashipet_proto_enumTypes[0]
}

func (x Species) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Species.Descriptor instead.
func (Species) EnumDescriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{0}
}

type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner      string  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Species    Species `protobuf:"varint,3,opt,name=species,proto3,enum=hashipet.v1.Species" json:"species,omitempty"`
	PictureUrl string  `protobuf:"bytes,4,opt,name=picture_url,json=pictureUrl,proto3" json:"picture_url,omitempty"`
	TheBest    bool    `protobuf:"varint,5,opt,name=the_best,json=theBest,proto3" json:"the_best,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{0}
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Pet) GetSpecies() Species {
	if x != nil {
		return x.Species
	}
	return Species_SPECIES_UNSPECIFIED
}

func (x *Pet) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

func (x *Pet) GetTheBest() bool {
	if x != nil {
		return x.TheBest
	}
	return false
}

type CreatePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *CreatePetRequest) Reset() {
	*x = CreatePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetRequest) ProtoMessage() {}

func (x *CreatePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetRequest.ProtoReflect.Descriptor instead.
func (*CreatePetRequest) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePetRequest) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type CreatePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *CreatePetResponse) Reset() {
	*x = CreatePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetResponse) ProtoMessage() {}

func (x *CreatePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetResponse.ProtoReflect.Descriptor instead.
func (*CreatePetResponse) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type GetPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPetRequest) Reset() {
	*x = GetPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetRequest) ProtoMessage() {}

func (x *GetPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetRequest.ProtoReflect.Descriptor instead.
func (*GetPetRequest) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{3}
}

func (x *GetPetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *GetPetResponse) Reset() {
	*x = GetPetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetResponse) ProtoMessage() {}

func (x *GetPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetResponse.ProtoReflect.Descriptor instead.
func (*GetPetResponse) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{4}
}

func (x *GetPetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type ListPetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPetsRequest) Reset() {
	*x = ListPetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetsRequest) ProtoMessage() {}

func (x *ListPetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetsRequest.ProtoReflect.Descriptor instead.
func (*ListPetsRequest) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{5}
}

type ListPetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pets []*Pet `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
}

func (x *ListPetsResponse) Reset() {
	*x = ListPetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetsResponse) ProtoMessage() {}

func (x *ListPetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetsResponse.ProtoReflect.Descriptor instead.
func (*ListPetsResponse) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{6}
}

func (x *ListPetsResponse) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

type DeletePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeletePetRequest) Reset() {
	*x = DeletePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetRequest) ProtoMessage() {}

func (x *DeletePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetRequest.ProtoReflect.Descriptor instead.
func (*DeletePetRequest) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdatePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet        *Pet                   `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdatePetRequest) Reset() {
	*x = UpdatePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetRequest) ProtoMessage() {}

func (x *UpdatePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetRequest.ProtoReflect.Descriptor instead.
func (*UpdatePetRequest) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePetRequest) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

func (x *UpdatePetRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdatePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *UpdatePetResponse) Reset() {
	*x = UpdatePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashipet_v1_hashipet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetResponse) ProtoMessage() {}

func (x *UpdatePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hashipet_v1_hashipet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetResponse.ProtoReflect.Descriptor instead.
func (*UpdatePetResponse) Descriptor() ([]byte, []int) {
	return file_hashipet_v1_hashipet_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

var File_hashipet_v1_hashipet_proto protoreflect.FileDescriptor

var file_hashipet_v1_hashipet_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52,
	0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x68, 0x65,
	0x5f, 0x62, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x68, 0x65,
	0x42, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x37, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74,
	0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x03,
	0x70, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x70, 0x65, 0x74,
	0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04, 0x70, 0x65, 0x74, 0x73, 0x22, 0x26, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x03, 0x70, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x37, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x03,
	0x70, 0x65, 0x74, 0x2a, 0x55, 0x0a, 0x07, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x17,
	0x0a, 0x13, 0x53, 0x50, 0x45, 0x43, 0x49, 0x45, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x45, 0x53, 0x5f, 0x50, 0x49, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x45, 0x53, 0x5f, 0x44, 0x4f, 0x47, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x45, 0x53, 0x5f, 0x43, 0x41, 0x54, 0x10, 0x03, 0x32, 0x85, 0x04, 0x0a, 0x0f, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x3a, 0x03, 0x70, 0x65, 0x74, 0x62, 0x03, 0x70, 0x65, 0x74, 0x22, 0x08, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74,
	0x12, 0x1a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x62, 0x03, 0x70, 0x65, 0x74, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x59, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65,
	0x74, 0x73, 0x12, 0x5b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12,
	0x1d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12,
	0x71, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x3a, 0x03, 0x70, 0x65, 0x74, 0x62, 0x03, 0x70, 0x65, 0x74, 0x32, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x65, 0x74, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x42, 0x84, 0x02, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x68, 0x61, 0x6e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x68, 0x6f,
	0x72, 0x73, 0x74, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa,
	0x02, 0x0b, 0x48, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x48, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x92, 0x41, 0x52, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01,
	0x01, 0x72, 0x46, 0x0a, 0x17, 0x48, 0x61, 0x73, 0x68, 0x69, 0x50, 0x65, 0x74, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x6f, 0x68, 0x61, 0x6e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x68, 0x6f, 0x72, 0x73, 0x74,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x70, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_hashipet_v1_hashipet_proto_rawDescOnce sync.Once
	file_hashipet_v1_hashipet_proto_rawDescData = file_hashipet_v1_hashipet_proto_rawDesc
)

func file_hashipet_v1_hashipet_proto_rawDescGZIP() []byte {
	file_hashipet_v1_hashipet_proto_rawDescOnce.Do(func() {
		file_hashipet_v1_hashipet_proto_rawDescData = protoimpl.X.CompressGZIP(file_hashipet_v1_hashipet_proto_rawDescData)
	})
	return file_hashipet_v1_hashipet_proto_rawDescData
}

var file_hashipet_v1_hashipet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hashipet_v1_hashipet_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_hashipet_v1_hashipet_proto_goTypes = []interface{}{
	(Species)(0),                  // 0: hashipet.v1.Species
	(*Pet)(nil),                   // 1: hashipet.v1.Pet
	(*CreatePetRequest)(nil),      // 2: hashipet.v1.CreatePetRequest
	(*CreatePetResponse)(nil),     // 3: hashipet.v1.CreatePetResponse
	(*GetPetRequest)(nil),         // 4: hashipet.v1.GetPetRequest
	(*GetPetResponse)(nil),        // 5: hashipet.v1.GetPetResponse
	(*ListPetsRequest)(nil),       // 6: hashipet.v1.ListPetsRequest
	(*ListPetsResponse)(nil),      // 7: hashipet.v1.ListPetsResponse
	(*DeletePetRequest)(nil),      // 8: hashipet.v1.DeletePetRequest
	(*UpdatePetRequest)(nil),      // 9: hashipet.v1.UpdatePetRequest
	(*UpdatePetResponse)(nil),     // 10: hashipet.v1.UpdatePetResponse
	(*fieldmaskpb.FieldMask)(nil), // 11: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_hashipet_v1_hashipet_proto_depIdxs = []int32{
	0,  // 0: hashipet.v1.Pet.species:type_name -> hashipet.v1.Species
	1,  // 1: hashipet.v1.CreatePetRequest.pet:type_name -> hashipet.v1.Pet
	1,  // 2: hashipet.v1.CreatePetResponse.pet:type_name -> hashipet.v1.Pet
	1,  // 3: hashipet.v1.GetPetResponse.pet:type_name -> hashipet.v1.Pet
	1,  // 4: hashipet.v1.ListPetsResponse.pets:type_name -> hashipet.v1.Pet
	1,  // 5: hashipet.v1.UpdatePetRequest.pet:type_name -> hashipet.v1.Pet
	11, // 6: hashipet.v1.UpdatePetRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 7: hashipet.v1.UpdatePetResponse.pet:type_name -> hashipet.v1.Pet
	2,  // 8: hashipet.v1.HashiPetService.CreatePet:input_type -> hashipet.v1.CreatePetRequest
	4,  // 9: hashipet.v1.HashiPetService.GetPet:input_type -> hashipet.v1.GetPetRequest
	6,  // 10: hashipet.v1.HashiPetService.ListPets:input_type -> hashipet.v1.ListPetsRequest
	8,  // 11: hashipet.v1.HashiPetService.DeletePet:input_type -> hashipet.v1.DeletePetRequest
	9,  // 12: hashipet.v1.HashiPetService.UpdatePet:input_type -> hashipet.v1.UpdatePetRequest
	3,  // 13: hashipet.v1.HashiPetService.CreatePet:output_type -> hashipet.v1.CreatePetResponse
	5,  // 14: hashipet.v1.HashiPetService.GetPet:output_type -> hashipet.v1.GetPetResponse
	7,  // 15: hashipet.v1.HashiPetService.ListPets:output_type -> hashipet.v1.ListPetsResponse
	12, // 16: hashipet.v1.HashiPetService.DeletePet:output_type -> google.protobuf.Empty
	10, // 17: hashipet.v1.HashiPetService.UpdatePet:output_type -> hashipet.v1.UpdatePetResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_hashipet_v1_hashipet_proto_init() }
func file_hashipet_v1_hashipet_proto_init() {
	if File_hashipet_v1_hashipet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hashipet_v1_hashipet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePetRequest); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePetResponse); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetRequest); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetResponse); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPetsRequest); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPetsResponse); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetRequest); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetRequest); i {
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
		file_hashipet_v1_hashipet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetResponse); i {
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
			RawDescriptor: file_hashipet_v1_hashipet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hashipet_v1_hashipet_proto_goTypes,
		DependencyIndexes: file_hashipet_v1_hashipet_proto_depIdxs,
		EnumInfos:         file_hashipet_v1_hashipet_proto_enumTypes,
		MessageInfos:      file_hashipet_v1_hashipet_proto_msgTypes,
	}.Build()
	File_hashipet_v1_hashipet_proto = out.File
	file_hashipet_v1_hashipet_proto_rawDesc = nil
	file_hashipet_v1_hashipet_proto_goTypes = nil
	file_hashipet_v1_hashipet_proto_depIdxs = nil
}
