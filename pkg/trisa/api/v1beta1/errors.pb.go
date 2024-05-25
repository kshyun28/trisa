// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: trisa/api/v1beta1/errors.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Error_Code int32

const (
	// Errors 0-49 are reserved for service-specific errors
	// Default error - something very bad happened.
	Error_UNHANDLED Error_Code = 0
	// Generic error - could not handle request, retry later.
	Error_UNAVAILABLE       Error_Code = 1
	Error_SERVICE_DOWN_TIME Error_Code = 1 // Alias: Sygna BVRC Rejected Type
	Error_BVRC002           Error_Code = 1 // Alias: Sygna BVRC Rejected Code
	// The service is currently in maintenance mode and cannot respond.
	Error_MAINTENANCE Error_Code = 2
	// The RPC is not currently implemented by the TRISA node.
	Error_UNIMPLEMENTED Error_Code = 3
	// Request could not be processed by recipient.
	Error_INTERNAL_ERROR Error_Code = 49
	Error_BVRC999        Error_Code = 49 // Alias: Sygna BVRC Rejected Code
	// Errors 50-99 are reserved for transaction rejections.
	// Default rejection - no specified reason for rejecting the transaction.
	Error_REJECTED Error_Code = 50
	// VASP does not control the specified wallet address.
	Error_UNKNOWN_WALLET_ADDRESS Error_Code = 51
	// VASP does not have KYC information for the specified wallet address.
	Error_UNKNOWN_IDENTITY Error_Code = 52
	Error_UNKOWN_IDENTITY  Error_Code = 52 // Typo left for backwards compatibility.
	// Specifically, the Originator account cannot be identified.
	Error_UNKNOWN_ORIGINATOR Error_Code = 53
	// Specifically, the Beneficiary account cannot be identified.
	Error_UNKOWN_BENEFICIARY         Error_Code = 54
	Error_BENEFICIARY_NAME_UNMATCHED Error_Code = 54 // Alias: Sygna BVRC Rejected Type
	Error_BVRC007                    Error_Code = 54 // Alias: Sygna BVRC Rejected Code
	// VASP cannot support the fiat currency or coin described in the transaction.
	Error_UNSUPPORTED_CURRENCY Error_Code = 60
	Error_BVRC001              Error_Code = 60 // Alias: Sygna BVRC Rejected Code
	// No longer able to receive more transaction inflows
	Error_EXCEEDED_TRADING_VOLUME Error_Code = 61
	Error_BVRC003                 Error_Code = 61 // Alias: Sygna BVRC Rejected Code
	// An internal compliance check has failed or black listing has occurred
	Error_COMPLIANCE_CHECK_FAIL Error_Code = 90
	Error_BVRC004               Error_Code = 90 // Alias: Sygna BVRC Rejected Code
	// VASP not able to implement travel rule compliance.
	Error_NO_COMPLIANCE Error_Code = 91
	// VASP unwilling to conduct the transaction because of a risk assessment.
	Error_HIGH_RISK Error_Code = 92
	// Wallet address or transaction is not available on this network.
	Error_OUT_OF_NETWORK Error_Code = 99
	// Errors 100-149 are reserved for authentication or cryptography failures.
	// Default access denied - no specified reason for forbidding the transaction.
	Error_FORBIDDEN Error_Code = 100
	// Could not sign transaction because no signing key is available.
	Error_NO_SIGNING_KEY Error_Code = 101
	// Could not sign transaction because keys have been revoked.
	Error_CERTIFICATE_REVOKED Error_Code = 102
	// Could not verify certificates with any certificate authority.
	Error_UNVERIFIED Error_Code = 103
	// A trusted connection could not be established.
	Error_UNTRUSTED Error_Code = 104
	// An HMAC signature could not be verified
	Error_INVALID_SIGNATURE Error_Code = 105
	// The transaction bundle cannot be decrypted with the specified key
	Error_INVALID_KEY Error_Code = 106
	// Could not decode or decrypt private transaction data
	Error_ENVELOPE_DECODE_FAIL     Error_Code = 107
	Error_PRIVATE_INFO_DECODE_FAIL Error_Code = 107 // Alias: Sygna BVRC Rejected Type
	Error_BVRC005                  Error_Code = 107 // Alias: Sygna BVRC Rejected Code
	// The algorithm specified by the encryption or signature is not implemented
	Error_UNHANDLED_ALGORITHM Error_Code = 108
	// Errors 150-199 are reserved for repairing exchange information.
	// Generic bad request - usually implies retry when reuqest is fixed.
	Error_BAD_REQUEST Error_Code = 150
	// Could not parse the identity record; specify the type in extra
	Error_UNPARSEABLE_IDENTITY      Error_Code = 151
	Error_PRIVATE_INFO_WRONG_FORMAT Error_Code = 151 // Alias: Sygna BVRC Rejected Type
	Error_BVRC006                   Error_Code = 151 // Alias: Sygna BVRC Rejected Code
	// Could not parse the transaction data record; specify the type in extra
	Error_UNPARSEABLE_TRANSACTION Error_Code = 152
	// There are missing required fields in the transaction data, a list of these
	// fields is specified in extra.
	Error_MISSING_FIELDS Error_Code = 153
	// The identity record is not complete enough for compliance purposes of the
	// receiving VASPs. Required fields or format specified in extra.
	Error_INCOMPLETE_IDENTITY Error_Code = 154
	// There was an error validating a field in the transaction data (specified in extra)
	Error_VALIDATION_ERROR Error_Code = 155
)

// Enum value maps for Error_Code.
var (
	Error_Code_name = map[int32]string{
		0: "UNHANDLED",
		1: "UNAVAILABLE",
		// Duplicate value: 1: "SERVICE_DOWN_TIME",
		// Duplicate value: 1: "BVRC002",
		2:  "MAINTENANCE",
		3:  "UNIMPLEMENTED",
		49: "INTERNAL_ERROR",
		// Duplicate value: 49: "BVRC999",
		50: "REJECTED",
		51: "UNKNOWN_WALLET_ADDRESS",
		52: "UNKNOWN_IDENTITY",
		// Duplicate value: 52: "UNKOWN_IDENTITY",
		53: "UNKNOWN_ORIGINATOR",
		54: "UNKOWN_BENEFICIARY",
		// Duplicate value: 54: "BENEFICIARY_NAME_UNMATCHED",
		// Duplicate value: 54: "BVRC007",
		60: "UNSUPPORTED_CURRENCY",
		// Duplicate value: 60: "BVRC001",
		61: "EXCEEDED_TRADING_VOLUME",
		// Duplicate value: 61: "BVRC003",
		90: "COMPLIANCE_CHECK_FAIL",
		// Duplicate value: 90: "BVRC004",
		91:  "NO_COMPLIANCE",
		92:  "HIGH_RISK",
		99:  "OUT_OF_NETWORK",
		100: "FORBIDDEN",
		101: "NO_SIGNING_KEY",
		102: "CERTIFICATE_REVOKED",
		103: "UNVERIFIED",
		104: "UNTRUSTED",
		105: "INVALID_SIGNATURE",
		106: "INVALID_KEY",
		107: "ENVELOPE_DECODE_FAIL",
		// Duplicate value: 107: "PRIVATE_INFO_DECODE_FAIL",
		// Duplicate value: 107: "BVRC005",
		108: "UNHANDLED_ALGORITHM",
		150: "BAD_REQUEST",
		151: "UNPARSEABLE_IDENTITY",
		// Duplicate value: 151: "PRIVATE_INFO_WRONG_FORMAT",
		// Duplicate value: 151: "BVRC006",
		152: "UNPARSEABLE_TRANSACTION",
		153: "MISSING_FIELDS",
		154: "INCOMPLETE_IDENTITY",
		155: "VALIDATION_ERROR",
	}
	Error_Code_value = map[string]int32{
		"UNHANDLED":                  0,
		"UNAVAILABLE":                1,
		"SERVICE_DOWN_TIME":          1,
		"BVRC002":                    1,
		"MAINTENANCE":                2,
		"UNIMPLEMENTED":              3,
		"INTERNAL_ERROR":             49,
		"BVRC999":                    49,
		"REJECTED":                   50,
		"UNKNOWN_WALLET_ADDRESS":     51,
		"UNKNOWN_IDENTITY":           52,
		"UNKOWN_IDENTITY":            52,
		"UNKNOWN_ORIGINATOR":         53,
		"UNKOWN_BENEFICIARY":         54,
		"BENEFICIARY_NAME_UNMATCHED": 54,
		"BVRC007":                    54,
		"UNSUPPORTED_CURRENCY":       60,
		"BVRC001":                    60,
		"EXCEEDED_TRADING_VOLUME":    61,
		"BVRC003":                    61,
		"COMPLIANCE_CHECK_FAIL":      90,
		"BVRC004":                    90,
		"NO_COMPLIANCE":              91,
		"HIGH_RISK":                  92,
		"OUT_OF_NETWORK":             99,
		"FORBIDDEN":                  100,
		"NO_SIGNING_KEY":             101,
		"CERTIFICATE_REVOKED":        102,
		"UNVERIFIED":                 103,
		"UNTRUSTED":                  104,
		"INVALID_SIGNATURE":          105,
		"INVALID_KEY":                106,
		"ENVELOPE_DECODE_FAIL":       107,
		"PRIVATE_INFO_DECODE_FAIL":   107,
		"BVRC005":                    107,
		"UNHANDLED_ALGORITHM":        108,
		"BAD_REQUEST":                150,
		"UNPARSEABLE_IDENTITY":       151,
		"PRIVATE_INFO_WRONG_FORMAT":  151,
		"BVRC006":                    151,
		"UNPARSEABLE_TRANSACTION":    152,
		"MISSING_FIELDS":             153,
		"INCOMPLETE_IDENTITY":        154,
		"VALIDATION_ERROR":           155,
	}
)

func (x Error_Code) Enum() *Error_Code {
	p := new(Error_Code)
	*p = x
	return p
}

func (x Error_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_trisa_api_v1beta1_errors_proto_enumTypes[0].Descriptor()
}

func (Error_Code) Type() protoreflect.EnumType {
	return &file_trisa_api_v1beta1_errors_proto_enumTypes[0]
}

func (x Error_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Code.Descriptor instead.
func (Error_Code) EnumDescriptor() ([]byte, []int) {
	return file_trisa_api_v1beta1_errors_proto_rawDescGZIP(), []int{0, 0}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error codes are standardized in the TRISA network to prevent confusion and to
	// allow easy identification of rejections or other problems so that the repair of
	// the connection or information exchange can be expedited.
	Code Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=trisa.api.v1beta1.Error_Code" json:"code,omitempty"`
	// Human readable message stating the reason for the error, should be loggable and
	// actionable. Both standardized and unique/detail messages are acceptable.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// If the message that caused the error should be retried with a fix; otherwise the
	// rejection is permenant and the request should not be retried.
	Retry bool `protobuf:"varint,3,opt,name=retry,proto3" json:"retry,omitempty"`
	// Any additional data or reasons for the rejection, e.g. a parent error, a diff,
	// a location for redirect, etc. The payload of the additional data should be
	// described by the error code.
	Details *anypb.Any `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trisa_api_v1beta1_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_trisa_api_v1beta1_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_trisa_api_v1beta1_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() Error_Code {
	if x != nil {
		return x.Code
	}
	return Error_UNHANDLED
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetRetry() bool {
	if x != nil {
		return x.Retry
	}
	return false
}

func (x *Error) GetDetails() *anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_trisa_api_v1beta1_errors_proto protoreflect.FileDescriptor

var file_trisa_api_v1beta1_errors_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x72, 0x69, 0x73, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x74, 0x72, 0x69, 0x73, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1,
	0x08, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x72, 0x69, 0x73, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x94, 0x07, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x44, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42,
	0x56, 0x52, 0x43, 0x30, 0x30, 0x32, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x41, 0x49, 0x4e,
	0x54, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x49,
	0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x31,
	0x12, 0x0b, 0x0a, 0x07, 0x42, 0x56, 0x52, 0x43, 0x39, 0x39, 0x39, 0x10, 0x31, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x32, 0x12, 0x1a, 0x0a, 0x16, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x57, 0x41, 0x4c, 0x4c, 0x45, 0x54, 0x5f, 0x41, 0x44,
	0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x33, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x34, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x4e, 0x4b, 0x4f, 0x57, 0x4e, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59,
	0x10, 0x34, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4f, 0x52,
	0x49, 0x47, 0x49, 0x4e, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x35, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e,
	0x4b, 0x4f, 0x57, 0x4e, 0x5f, 0x42, 0x45, 0x4e, 0x45, 0x46, 0x49, 0x43, 0x49, 0x41, 0x52, 0x59,
	0x10, 0x36, 0x12, 0x1e, 0x0a, 0x1a, 0x42, 0x45, 0x4e, 0x45, 0x46, 0x49, 0x43, 0x49, 0x41, 0x52,
	0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44,
	0x10, 0x36, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x56, 0x52, 0x43, 0x30, 0x30, 0x37, 0x10, 0x36, 0x12,
	0x18, 0x0a, 0x14, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x43,
	0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x3c, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x56, 0x52,
	0x43, 0x30, 0x30, 0x31, 0x10, 0x3c, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44,
	0x45, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x4f, 0x4c, 0x55, 0x4d,
	0x45, 0x10, 0x3d, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x56, 0x52, 0x43, 0x30, 0x30, 0x33, 0x10, 0x3d,
	0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x43,
	0x48, 0x45, 0x43, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x5a, 0x12, 0x0b, 0x0a, 0x07, 0x42,
	0x56, 0x52, 0x43, 0x30, 0x30, 0x34, 0x10, 0x5a, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x5b, 0x12, 0x0d, 0x0a, 0x09, 0x48,
	0x49, 0x47, 0x48, 0x5f, 0x52, 0x49, 0x53, 0x4b, 0x10, 0x5c, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x55,
	0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x63, 0x12, 0x0d,
	0x0a, 0x09, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x64, 0x12, 0x12, 0x0a,
	0x0e, 0x4e, 0x4f, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x59, 0x10,
	0x65, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x5f, 0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x66, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e,
	0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x67, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e,
	0x54, 0x52, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x68, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x69,
	0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4b, 0x45, 0x59, 0x10,
	0x6a, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x5f, 0x44, 0x45,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6b, 0x12, 0x1c, 0x0a, 0x18, 0x50,
	0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x44, 0x45, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x6b, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x56, 0x52,
	0x43, 0x30, 0x30, 0x35, 0x10, 0x6b, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x48, 0x41, 0x4e, 0x44,
	0x4c, 0x45, 0x44, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x10, 0x6c, 0x12,
	0x10, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x96,
	0x01, 0x12, 0x19, 0x0a, 0x14, 0x55, 0x4e, 0x50, 0x41, 0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x97, 0x01, 0x12, 0x1e, 0x0a, 0x19,
	0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x57, 0x52, 0x4f,
	0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x97, 0x01, 0x12, 0x0c, 0x0a, 0x07,
	0x42, 0x56, 0x52, 0x43, 0x30, 0x30, 0x36, 0x10, 0x97, 0x01, 0x12, 0x1c, 0x0a, 0x17, 0x55, 0x4e,
	0x50, 0x41, 0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x98, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x10, 0x99, 0x01, 0x12, 0x18, 0x0a,
	0x13, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x10, 0x9a, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x9b, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x72, 0x69, 0x73, 0x61, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x69,
	0x73, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x69, 0x73, 0x61, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trisa_api_v1beta1_errors_proto_rawDescOnce sync.Once
	file_trisa_api_v1beta1_errors_proto_rawDescData = file_trisa_api_v1beta1_errors_proto_rawDesc
)

func file_trisa_api_v1beta1_errors_proto_rawDescGZIP() []byte {
	file_trisa_api_v1beta1_errors_proto_rawDescOnce.Do(func() {
		file_trisa_api_v1beta1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_trisa_api_v1beta1_errors_proto_rawDescData)
	})
	return file_trisa_api_v1beta1_errors_proto_rawDescData
}

var file_trisa_api_v1beta1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trisa_api_v1beta1_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_trisa_api_v1beta1_errors_proto_goTypes = []interface{}{
	(Error_Code)(0),   // 0: trisa.api.v1beta1.Error.Code
	(*Error)(nil),     // 1: trisa.api.v1beta1.Error
	(*anypb.Any)(nil), // 2: google.protobuf.Any
}
var file_trisa_api_v1beta1_errors_proto_depIdxs = []int32{
	0, // 0: trisa.api.v1beta1.Error.code:type_name -> trisa.api.v1beta1.Error.Code
	2, // 1: trisa.api.v1beta1.Error.details:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trisa_api_v1beta1_errors_proto_init() }
func file_trisa_api_v1beta1_errors_proto_init() {
	if File_trisa_api_v1beta1_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trisa_api_v1beta1_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_trisa_api_v1beta1_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trisa_api_v1beta1_errors_proto_goTypes,
		DependencyIndexes: file_trisa_api_v1beta1_errors_proto_depIdxs,
		EnumInfos:         file_trisa_api_v1beta1_errors_proto_enumTypes,
		MessageInfos:      file_trisa_api_v1beta1_errors_proto_msgTypes,
	}.Build()
	File_trisa_api_v1beta1_errors_proto = out.File
	file_trisa_api_v1beta1_errors_proto_rawDesc = nil
	file_trisa_api_v1beta1_errors_proto_goTypes = nil
	file_trisa_api_v1beta1_errors_proto_depIdxs = nil
}
