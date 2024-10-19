// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: response_volume_profile_minute_bars.proto

package rti

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

type ResponseVolumeProfileMinuteBars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId                *int32    `protobuf:"varint,154467,req,name=template_id,json=templateId" json:"template_id,omitempty"`
	RequestKey                *string   `protobuf:"bytes,132758,opt,name=request_key,json=requestKey" json:"request_key,omitempty"`
	UserMsg                   []string  `protobuf:"bytes,132760,rep,name=user_msg,json=userMsg" json:"user_msg,omitempty"`
	RqHandlerRpCode           []string  `protobuf:"bytes,132764,rep,name=rq_handler_rp_code,json=rqHandlerRpCode" json:"rq_handler_rp_code,omitempty"`
	RpCode                    []string  `protobuf:"bytes,132766,rep,name=rp_code,json=rpCode" json:"rp_code,omitempty"`
	Symbol                    *string   `protobuf:"bytes,110100,opt,name=symbol" json:"symbol,omitempty"`
	Exchange                  *string   `protobuf:"bytes,110101,opt,name=exchange" json:"exchange,omitempty"`
	Period                    *string   `protobuf:"bytes,119215,opt,name=period" json:"period,omitempty"`
	Marker                    *int32    `protobuf:"varint,119100,opt,name=marker" json:"marker,omitempty"`
	NumTrades                 *uint64   `protobuf:"varint,119204,opt,name=num_trades,json=numTrades" json:"num_trades,omitempty"`
	Volume                    *uint64   `protobuf:"varint,119205,opt,name=volume" json:"volume,omitempty"`
	BidVolume                 *uint64   `protobuf:"varint,119213,opt,name=bid_volume,json=bidVolume" json:"bid_volume,omitempty"`
	AskVolume                 *uint64   `protobuf:"varint,119214,opt,name=ask_volume,json=askVolume" json:"ask_volume,omitempty"`
	OpenPrice                 *float64  `protobuf:"fixed64,100019,opt,name=open_price,json=openPrice" json:"open_price,omitempty"`
	ClosePrice                *float64  `protobuf:"fixed64,100021,opt,name=close_price,json=closePrice" json:"close_price,omitempty"`
	HighPrice                 *float64  `protobuf:"fixed64,100012,opt,name=high_price,json=highPrice" json:"high_price,omitempty"`
	LowPrice                  *float64  `protobuf:"fixed64,100013,opt,name=low_price,json=lowPrice" json:"low_price,omitempty"`
	ProfilePrice              []float64 `protobuf:"fixed64,119216,rep,name=profile_price,json=profilePrice" json:"profile_price,omitempty"`
	ProfileNoAggressorVolume  []int32   `protobuf:"varint,119217,rep,name=profile_no_aggressor_volume,json=profileNoAggressorVolume" json:"profile_no_aggressor_volume,omitempty"`
	ProfileBidVolume          []int32   `protobuf:"varint,119218,rep,name=profile_bid_volume,json=profileBidVolume" json:"profile_bid_volume,omitempty"`
	ProfileAskVolume          []int32   `protobuf:"varint,119219,rep,name=profile_ask_volume,json=profileAskVolume" json:"profile_ask_volume,omitempty"`
	ProfileNoAggressorTrades  []int32   `protobuf:"varint,119220,rep,name=profile_no_aggressor_trades,json=profileNoAggressorTrades" json:"profile_no_aggressor_trades,omitempty"`
	ProfileBidAggressorTrades []int32   `protobuf:"varint,119221,rep,name=profile_bid_aggressor_trades,json=profileBidAggressorTrades" json:"profile_bid_aggressor_trades,omitempty"`
	ProfileAskAggressorTrades []int32   `protobuf:"varint,119222,rep,name=profile_ask_aggressor_trades,json=profileAskAggressorTrades" json:"profile_ask_aggressor_trades,omitempty"`
}

func (x *ResponseVolumeProfileMinuteBars) Reset() {
	*x = ResponseVolumeProfileMinuteBars{}
	mi := &file_response_volume_profile_minute_bars_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseVolumeProfileMinuteBars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseVolumeProfileMinuteBars) ProtoMessage() {}

func (x *ResponseVolumeProfileMinuteBars) ProtoReflect() protoreflect.Message {
	mi := &file_response_volume_profile_minute_bars_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseVolumeProfileMinuteBars.ProtoReflect.Descriptor instead.
func (*ResponseVolumeProfileMinuteBars) Descriptor() ([]byte, []int) {
	return file_response_volume_profile_minute_bars_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseVolumeProfileMinuteBars) GetTemplateId() int32 {
	if x != nil && x.TemplateId != nil {
		return *x.TemplateId
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetRequestKey() string {
	if x != nil && x.RequestKey != nil {
		return *x.RequestKey
	}
	return ""
}

func (x *ResponseVolumeProfileMinuteBars) GetUserMsg() []string {
	if x != nil {
		return x.UserMsg
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetRqHandlerRpCode() []string {
	if x != nil {
		return x.RqHandlerRpCode
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetRpCode() []string {
	if x != nil {
		return x.RpCode
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetSymbol() string {
	if x != nil && x.Symbol != nil {
		return *x.Symbol
	}
	return ""
}

func (x *ResponseVolumeProfileMinuteBars) GetExchange() string {
	if x != nil && x.Exchange != nil {
		return *x.Exchange
	}
	return ""
}

func (x *ResponseVolumeProfileMinuteBars) GetPeriod() string {
	if x != nil && x.Period != nil {
		return *x.Period
	}
	return ""
}

func (x *ResponseVolumeProfileMinuteBars) GetMarker() int32 {
	if x != nil && x.Marker != nil {
		return *x.Marker
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetNumTrades() uint64 {
	if x != nil && x.NumTrades != nil {
		return *x.NumTrades
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetVolume() uint64 {
	if x != nil && x.Volume != nil {
		return *x.Volume
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetBidVolume() uint64 {
	if x != nil && x.BidVolume != nil {
		return *x.BidVolume
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetAskVolume() uint64 {
	if x != nil && x.AskVolume != nil {
		return *x.AskVolume
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetOpenPrice() float64 {
	if x != nil && x.OpenPrice != nil {
		return *x.OpenPrice
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetClosePrice() float64 {
	if x != nil && x.ClosePrice != nil {
		return *x.ClosePrice
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetHighPrice() float64 {
	if x != nil && x.HighPrice != nil {
		return *x.HighPrice
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetLowPrice() float64 {
	if x != nil && x.LowPrice != nil {
		return *x.LowPrice
	}
	return 0
}

func (x *ResponseVolumeProfileMinuteBars) GetProfilePrice() []float64 {
	if x != nil {
		return x.ProfilePrice
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetProfileNoAggressorVolume() []int32 {
	if x != nil {
		return x.ProfileNoAggressorVolume
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetProfileBidVolume() []int32 {
	if x != nil {
		return x.ProfileBidVolume
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetProfileAskVolume() []int32 {
	if x != nil {
		return x.ProfileAskVolume
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetProfileNoAggressorTrades() []int32 {
	if x != nil {
		return x.ProfileNoAggressorTrades
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetProfileBidAggressorTrades() []int32 {
	if x != nil {
		return x.ProfileBidAggressorTrades
	}
	return nil
}

func (x *ResponseVolumeProfileMinuteBars) GetProfileAskAggressorTrades() []int32 {
	if x != nil {
		return x.ProfileAskAggressorTrades
	}
	return nil
}

var File_response_volume_profile_minute_bars_proto protoreflect.FileDescriptor

var file_response_volume_profile_minute_bars_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x5f, 0x62, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x74, 0x69,
	0x22, 0xca, 0x07, 0x0a, 0x1f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x42, 0x61, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0xe3, 0xb6, 0x09, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x96, 0x8d, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x98, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x71, 0x5f, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x9c, 0x8d,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x71, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x52, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x72, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x9e, 0x8d, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x94, 0xdc, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x95, 0xdc, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0xaf, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x18, 0xbc, 0xa2,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0xa4, 0xa3, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0xa5, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0xad, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x62, 0x69, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x61, 0x73, 0x6b,
	0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0xae, 0xa3, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x61, 0x73, 0x6b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xb3, 0x8d, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xb5, 0x8d, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f,
	0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xac, 0x8d, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x09, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xad, 0x8d, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0xb0, 0xa3, 0x07, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6e, 0x6f, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0xb1, 0xa3, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x18, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x41, 0x67, 0x67, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0xb2, 0xa3, 0x07,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x69, 0x64,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0xb3, 0xa3, 0x07,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x73, 0x6b,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x6f, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0xb4, 0xa3, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x18, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x6f, 0x41, 0x67, 0x67, 0x72, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x62, 0x69, 0x64, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0xb5, 0xa3, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x19, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x69, 0x64, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x1c, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0xb6, 0xa3, 0x07, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x19, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x73, 0x6b, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x72, 0x74, 0x69,
}

var (
	file_response_volume_profile_minute_bars_proto_rawDescOnce sync.Once
	file_response_volume_profile_minute_bars_proto_rawDescData = file_response_volume_profile_minute_bars_proto_rawDesc
)

func file_response_volume_profile_minute_bars_proto_rawDescGZIP() []byte {
	file_response_volume_profile_minute_bars_proto_rawDescOnce.Do(func() {
		file_response_volume_profile_minute_bars_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_volume_profile_minute_bars_proto_rawDescData)
	})
	return file_response_volume_profile_minute_bars_proto_rawDescData
}

var file_response_volume_profile_minute_bars_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_response_volume_profile_minute_bars_proto_goTypes = []any{
	(*ResponseVolumeProfileMinuteBars)(nil), // 0: rti.ResponseVolumeProfileMinuteBars
}
var file_response_volume_profile_minute_bars_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_response_volume_profile_minute_bars_proto_init() }
func file_response_volume_profile_minute_bars_proto_init() {
	if File_response_volume_profile_minute_bars_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_volume_profile_minute_bars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_volume_profile_minute_bars_proto_goTypes,
		DependencyIndexes: file_response_volume_profile_minute_bars_proto_depIdxs,
		MessageInfos:      file_response_volume_profile_minute_bars_proto_msgTypes,
	}.Build()
	File_response_volume_profile_minute_bars_proto = out.File
	file_response_volume_profile_minute_bars_proto_rawDesc = nil
	file_response_volume_profile_minute_bars_proto_goTypes = nil
	file_response_volume_profile_minute_bars_proto_depIdxs = nil
}