// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: cards.proto

package card_db

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_cards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_cards_proto_rawDescGZIP(), []int{0}
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Search    string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_cards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_cards_proto_rawDescGZIP(), []int{1}
}

func (x *Query) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Query) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName   string   `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
	Season       string   `protobuf:"bytes,2,opt,name=season,proto3" json:"season,omitempty"`
	Manufacturer string   `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Set          string   `protobuf:"bytes,4,opt,name=set,proto3" json:"set,omitempty"`
	Insert       string   `protobuf:"bytes,5,opt,name=insert,proto3" json:"insert,omitempty"`
	Parallel     string   `protobuf:"bytes,6,opt,name=parallel,proto3" json:"parallel,omitempty"`
	CardNumber   string   `protobuf:"bytes,7,opt,name=cardNumber,proto3" json:"cardNumber,omitempty"`
	Notes        []string `protobuf:"bytes,8,rep,name=notes,proto3" json:"notes,omitempty"`
	ImageLink    string   `protobuf:"bytes,9,opt,name=imageLink,proto3" json:"imageLink,omitempty"`
	Team         string   `protobuf:"bytes,10,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_cards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_cards_proto_rawDescGZIP(), []int{2}
}

func (x *Card) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *Card) GetSeason() string {
	if x != nil {
		return x.Season
	}
	return ""
}

func (x *Card) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Card) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *Card) GetInsert() string {
	if x != nil {
		return x.Insert
	}
	return ""
}

func (x *Card) GetParallel() string {
	if x != nil {
		return x.Parallel
	}
	return ""
}

func (x *Card) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *Card) GetNotes() []string {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *Card) GetImageLink() string {
	if x != nil {
		return x.ImageLink
	}
	return ""
}

func (x *Card) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName     string   `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string   `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	SeasonsPlayed string   `protobuf:"bytes,3,opt,name=seasonsPlayed,proto3" json:"seasonsPlayed,omitempty"`
	Seasons       []string `protobuf:"bytes,4,rep,name=seasons,proto3" json:"seasons,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cards_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_cards_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_cards_proto_rawDescGZIP(), []int{3}
}

func (x *Player) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Player) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Player) GetSeasonsPlayed() string {
	if x != nil {
		return x.SeasonsPlayed
	}
	return ""
}

func (x *Player) GetSeasons() []string {
	if x != nil {
		return x.Seasons
	}
	return nil
}

type CardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerName   string   `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
	Season       string   `protobuf:"bytes,2,opt,name=season,proto3" json:"season,omitempty"`
	Manufacturer string   `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Set          string   `protobuf:"bytes,4,opt,name=set,proto3" json:"set,omitempty"`
	Insert       string   `protobuf:"bytes,5,opt,name=insert,proto3" json:"insert,omitempty"`
	Parallel     string   `protobuf:"bytes,6,opt,name=parallel,proto3" json:"parallel,omitempty"`
	CardNumber   string   `protobuf:"bytes,7,opt,name=cardNumber,proto3" json:"cardNumber,omitempty"`
	Notes        []string `protobuf:"bytes,8,rep,name=notes,proto3" json:"notes,omitempty"`
	ImageLink    string   `protobuf:"bytes,9,opt,name=imageLink,proto3" json:"imageLink,omitempty"`
	TableName    string   `protobuf:"bytes,10,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Team         string   `protobuf:"bytes,11,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *CardRequest) Reset() {
	*x = CardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cards_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardRequest) ProtoMessage() {}

func (x *CardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cards_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardRequest.ProtoReflect.Descriptor instead.
func (*CardRequest) Descriptor() ([]byte, []int) {
	return file_cards_proto_rawDescGZIP(), []int{4}
}

func (x *CardRequest) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *CardRequest) GetSeason() string {
	if x != nil {
		return x.Season
	}
	return ""
}

func (x *CardRequest) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *CardRequest) GetSet() string {
	if x != nil {
		return x.Set
	}
	return ""
}

func (x *CardRequest) GetInsert() string {
	if x != nil {
		return x.Insert
	}
	return ""
}

func (x *CardRequest) GetParallel() string {
	if x != nil {
		return x.Parallel
	}
	return ""
}

func (x *CardRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *CardRequest) GetNotes() []string {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *CardRequest) GetImageLink() string {
	if x != nil {
		return x.ImageLink
	}
	return ""
}

func (x *CardRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *CardRequest) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

type PlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName     string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	SeasonsPlayed string `protobuf:"bytes,3,opt,name=seasonsPlayed,proto3" json:"seasonsPlayed,omitempty"`
}

func (x *PlayerRequest) Reset() {
	*x = PlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cards_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRequest) ProtoMessage() {}

func (x *PlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cards_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRequest.ProtoReflect.Descriptor instead.
func (*PlayerRequest) Descriptor() ([]byte, []int) {
	return file_cards_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PlayerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PlayerRequest) GetSeasonsPlayed() string {
	if x != nil {
		return x.SeasonsPlayed
	}
	return ""
}

type SeasonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Season    string `protobuf:"bytes,1,opt,name=season,proto3" json:"season,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *SeasonRequest) Reset() {
	*x = SeasonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cards_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeasonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeasonRequest) ProtoMessage() {}

func (x *SeasonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cards_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeasonRequest.ProtoReflect.Descriptor instead.
func (*SeasonRequest) Descriptor() ([]byte, []int) {
	return file_cards_proto_rawDescGZIP(), []int{6}
}

func (x *SeasonRequest) GetSeason() string {
	if x != nil {
		return x.Season
	}
	return ""
}

func (x *SeasonRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SeasonRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cards_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_cards_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_cards_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_cards_proto protoreflect.FileDescriptor

var file_cards_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3d, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x90, 0x02, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x82,
	0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x73, 0x22, 0xb5, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6d,
	0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0x6f, 0x0a, 0x0d, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x22, 0x61, 0x0a, 0x0d,
	0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x26, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfb, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x2b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x0b,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x32, 0x0a,
	0x09, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x13,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x72, 0x64, 0x61, 0x6e, 0x78, 0x38, 0x2f, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cards_proto_rawDescOnce sync.Once
	file_cards_proto_rawDescData = file_cards_proto_rawDesc
)

func file_cards_proto_rawDescGZIP() []byte {
	file_cards_proto_rawDescOnce.Do(func() {
		file_cards_proto_rawDescData = protoimpl.X.CompressGZIP(file_cards_proto_rawDescData)
	})
	return file_cards_proto_rawDescData
}

var file_cards_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cards_proto_goTypes = []interface{}{
	(*Empty)(nil),         // 0: main.Empty
	(*Query)(nil),         // 1: main.Query
	(*Card)(nil),          // 2: main.Card
	(*Player)(nil),        // 3: main.Player
	(*CardRequest)(nil),   // 4: main.CardRequest
	(*PlayerRequest)(nil), // 5: main.PlayerRequest
	(*SeasonRequest)(nil), // 6: main.SeasonRequest
	(*Response)(nil),      // 7: main.Response
}
var file_cards_proto_depIdxs = []int32{
	1, // 0: main.CardService.GetCards:input_type -> main.Query
	0, // 1: main.CardService.GetPlayers:input_type -> main.Empty
	5, // 2: main.CardService.AddPlayer:input_type -> main.PlayerRequest
	6, // 3: main.CardService.AddSeason:input_type -> main.SeasonRequest
	4, // 4: main.CardService.AddCard:input_type -> main.CardRequest
	2, // 5: main.CardService.GetCards:output_type -> main.Card
	3, // 6: main.CardService.GetPlayers:output_type -> main.Player
	7, // 7: main.CardService.AddPlayer:output_type -> main.Response
	7, // 8: main.CardService.AddSeason:output_type -> main.Response
	7, // 9: main.CardService.AddCard:output_type -> main.Response
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cards_proto_init() }
func file_cards_proto_init() {
	if File_cards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_cards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_cards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_cards_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_cards_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardRequest); i {
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
		file_cards_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRequest); i {
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
		file_cards_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeasonRequest); i {
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
		file_cards_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_cards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cards_proto_goTypes,
		DependencyIndexes: file_cards_proto_depIdxs,
		MessageInfos:      file_cards_proto_msgTypes,
	}.Build()
	File_cards_proto = out.File
	file_cards_proto_rawDesc = nil
	file_cards_proto_goTypes = nil
	file_cards_proto_depIdxs = nil
}
