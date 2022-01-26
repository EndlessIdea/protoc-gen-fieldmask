// Code generated by protoc-gen-fieldmask. DO NOT EDIT.
// source: user.proto

package pb

import (
	pbfieldmask "github.com/yeqown/protoc-gen-fieldmask/proto/fieldmask"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (x *UserInfoRequest) FieldMaskWithMode(mode pbfieldmask.MaskMode) *UserInfoRequest_FieldMask {
	fm := &UserInfoRequest_FieldMask{
		maskMode: mode,
		mask:     pbfieldmask.New(x.FieldMask),
	}

	return fm
}

// FieldMask_Prune generates *UserInfoRequest_FieldMask with filter mode, so that
// only the fields in the UserInfoRequest.FieldMask will be
// appended into the UserInfoRequest.
func (x *UserInfoRequest) FieldMask_Filter() *UserInfoRequest_FieldMask {
	return x.FieldMaskWithMode(pbfieldmask.MaskMode_Filter)
}

// FieldMask_Prune generates *UserInfoRequest_FieldMask with prune mode, so that
// only the fields NOT in the UserInfoRequest.FieldMask will be
// appended into the UserInfoRequest.
func (x *UserInfoRequest) FieldMask_Prune() *UserInfoRequest_FieldMask {
	return x.FieldMaskWithMode(pbfieldmask.MaskMode_Prune)
}

// UserInfoRequest_FieldMask provide provide helper functions to deal with FieldMask.
type UserInfoRequest_FieldMask struct {
	maskMode pbfieldmask.MaskMode
	mask     pbfieldmask.NestedFieldMask
}

// _fm_UserInfoRequest is built in variable for UserInfoRequest to call FieldMask.Append
var _fm_UserInfoRequest = new(UserInfoRequest)

// MaskIn_UserId append UserInfoRequest.UserId into
// UserInfoRequest.FieldMask.
func (x *UserInfoRequest) MaskIn_UserId() *UserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UserInfoRequest, "user_id")

	return x
}

// MaskedIn_UserId indicates the field UserInfoRequest.UserId
// exists in the UserInfoRequest.FieldMask or not.
func (x *UserInfoRequest_FieldMask) MaskedIn_UserId() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("user_id")
}

// _fm_UserInfoResponse is built in variable for UserInfoResponse to call FieldMask.Append
var _fm_UserInfoResponse = new(UserInfoResponse)

// MaskOut_UserId append UserInfoResponse.UserId into
// UserInfoRequest.FieldMask.
func (x *UserInfoRequest) MaskOut_UserId() *UserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UserInfoResponse, "user_id")

	return x
}

// MaskedOut_UserId indicates the field UserInfoRequest.UserId
// exists in the UserInfoRequest.FieldMask or not.
func (x *UserInfoRequest_FieldMask) MaskedOut_UserId() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("user_id")
}

// MaskOut_Name append UserInfoResponse.Name into
// UserInfoRequest.FieldMask.
func (x *UserInfoRequest) MaskOut_Name() *UserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UserInfoResponse, "name")

	return x
}

// MaskedOut_Name indicates the field UserInfoRequest.Name
// exists in the UserInfoRequest.FieldMask or not.
func (x *UserInfoRequest_FieldMask) MaskedOut_Name() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("name")
}

// MaskOut_Email append UserInfoResponse.Email into
// UserInfoRequest.FieldMask.
func (x *UserInfoRequest) MaskOut_Email() *UserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UserInfoResponse, "email")

	return x
}

// MaskedOut_Email indicates the field UserInfoRequest.Email
// exists in the UserInfoRequest.FieldMask or not.
func (x *UserInfoRequest_FieldMask) MaskedOut_Email() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("email")
}

// MaskOut_Address append UserInfoResponse.Address into
// UserInfoRequest.FieldMask.
func (x *UserInfoRequest) MaskOut_Address() *UserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UserInfoResponse, "address")

	return x
}

// MaskedOut_Address indicates the field UserInfoRequest.Address
// exists in the UserInfoRequest.FieldMask or not.
func (x *UserInfoRequest_FieldMask) MaskedOut_Address() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("address")
}

// MaskOut_Address_Country append UserInfoResponse.Country into
// UserInfoRequest.FieldMask.
func (x *UserInfoRequest) MaskOut_Address_Country() *UserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UserInfoResponse, "address.country")

	return x
}

// MaskedOut_Address_Country indicates the field UserInfoRequest.Country
// exists in the UserInfoRequest.FieldMask or not.
func (x *UserInfoRequest_FieldMask) MaskedOut_Address_Country() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("address.country")
}

// MaskOut_Address_Province append UserInfoResponse.Province into
// UserInfoRequest.FieldMask.
func (x *UserInfoRequest) MaskOut_Address_Province() *UserInfoRequest {
	if x.FieldMask == nil {
		x.FieldMask = new(fieldmaskpb.FieldMask)
	}
	x.FieldMask.Append(_fm_UserInfoResponse, "address.province")

	return x
}

// MaskedOut_Address_Province indicates the field UserInfoRequest.Province
// exists in the UserInfoRequest.FieldMask or not.
func (x *UserInfoRequest_FieldMask) MaskedOut_Address_Province() bool {
	if x.mask == nil {
		return false
	}

	return x.mask.Masked("address.province")
}

// Mask only affects the fields in the UserInfoRequest.
func (x *UserInfoRequest_FieldMask) Mask(m *UserInfoResponse) *UserInfoResponse {
	switch x.maskMode {
	case pbfieldmask.MaskMode_Filter:
		x.mask.Filter(m)
	case pbfieldmask.MaskMode_Prune:
		x.mask.Prune(m)
	}

	return m
}
