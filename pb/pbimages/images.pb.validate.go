// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: images.proto

package pbimages

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _images_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CloneImagesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CloneImagesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetImages() {
		_ = val

		// no validation rules for Images[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CloneImagesResponseValidationError{
					field:  fmt.Sprintf("Images[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CloneImagesResponseValidationError is the validation error returned by
// CloneImagesResponse.Validate if the designated constraints aren't met.
type CloneImagesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CloneImagesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CloneImagesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CloneImagesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CloneImagesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CloneImagesResponseValidationError) ErrorName() string {
	return "CloneImagesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CloneImagesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCloneImagesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CloneImagesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CloneImagesResponseValidationError{}

// Validate checks the field values on CloneImagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CloneImagesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for FromEntityId

	// no validation rules for DestEntityId

	// no validation rules for EntityTable

	return nil
}

// CloneImagesRequestValidationError is the validation error returned by
// CloneImagesRequest.Validate if the designated constraints aren't met.
type CloneImagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CloneImagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CloneImagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CloneImagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CloneImagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CloneImagesRequestValidationError) ErrorName() string {
	return "CloneImagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CloneImagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCloneImagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CloneImagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CloneImagesRequestValidationError{}

// Validate checks the field values on ChangeEntityImagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ChangeEntityImagesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for FromEntityId

	// no validation rules for FromEntityTable

	// no validation rules for DestEntityId

	// no validation rules for DestEntityTable

	return nil
}

// ChangeEntityImagesRequestValidationError is the validation error returned by
// ChangeEntityImagesRequest.Validate if the designated constraints aren't met.
type ChangeEntityImagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChangeEntityImagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChangeEntityImagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChangeEntityImagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChangeEntityImagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChangeEntityImagesRequestValidationError) ErrorName() string {
	return "ChangeEntityImagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChangeEntityImagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChangeEntityImagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChangeEntityImagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChangeEntityImagesRequestValidationError{}

// Validate checks the field values on CreateImageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateImageRequest) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Corpus.(type) {

	case *CreateImageRequest_Metadata_:

		if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateImageRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CreateImageRequest_Chunk:
		// no validation rules for Chunk

	}

	return nil
}

// CreateImageRequestValidationError is the validation error returned by
// CreateImageRequest.Validate if the designated constraints aren't met.
type CreateImageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateImageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateImageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateImageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateImageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateImageRequestValidationError) ErrorName() string {
	return "CreateImageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateImageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateImageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateImageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateImageRequestValidationError{}

// Validate checks the field values on Image with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Image) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Size

	// no validation rules for Uri

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// ImageValidationError is the validation error returned by Image.Validate if
// the designated constraints aren't met.
type ImageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageValidationError) ErrorName() string { return "ImageValidationError" }

// Error satisfies the builtin error interface
func (e ImageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageValidationError{}

// Validate checks the field values on CreateImageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateImageResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ImageId

	if v, ok := interface{}(m.GetImage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateImageResponseValidationError{
				field:  "Image",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateImageResponseValidationError is the validation error returned by
// CreateImageResponse.Validate if the designated constraints aren't met.
type CreateImageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateImageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateImageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateImageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateImageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateImageResponseValidationError) ErrorName() string {
	return "CreateImageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateImageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateImageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateImageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateImageResponseValidationError{}

// Validate checks the field values on RetrieveImagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveImagesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// RetrieveImagesRequestValidationError is the validation error returned by
// RetrieveImagesRequest.Validate if the designated constraints aren't met.
type RetrieveImagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveImagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveImagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveImagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveImagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveImagesRequestValidationError) ErrorName() string {
	return "RetrieveImagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveImagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveImagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveImagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveImagesRequestValidationError{}

// Validate checks the field values on RetrieveImagesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveImagesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetImages() {
		_ = val

		// no validation rules for Images[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RetrieveImagesResponseValidationError{
					field:  fmt.Sprintf("Images[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RetrieveImagesResponseValidationError is the validation error returned by
// RetrieveImagesResponse.Validate if the designated constraints aren't met.
type RetrieveImagesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveImagesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveImagesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveImagesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveImagesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveImagesResponseValidationError) ErrorName() string {
	return "RetrieveImagesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveImagesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveImagesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveImagesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveImagesResponseValidationError{}

// Validate checks the field values on DeleteImagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteImagesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// DeleteImagesRequestValidationError is the validation error returned by
// DeleteImagesRequest.Validate if the designated constraints aren't met.
type DeleteImagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteImagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteImagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteImagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteImagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteImagesRequestValidationError) ErrorName() string {
	return "DeleteImagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteImagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteImagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteImagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteImagesRequestValidationError{}

// Validate checks the field values on CreateImageRequest_Metadata with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateImageRequest_Metadata) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _CreateImageRequest_Metadata_Ext_InLookup[m.GetExt()]; !ok {
		return CreateImageRequest_MetadataValidationError{
			field:  "Ext",
			reason: "value must be in list [svg jpg jpeg png]",
		}
	}

	// no validation rules for AccessToken

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// CreateImageRequest_MetadataValidationError is the validation error returned
// by CreateImageRequest_Metadata.Validate if the designated constraints
// aren't met.
type CreateImageRequest_MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateImageRequest_MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateImageRequest_MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateImageRequest_MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateImageRequest_MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateImageRequest_MetadataValidationError) ErrorName() string {
	return "CreateImageRequest_MetadataValidationError"
}

// Error satisfies the builtin error interface
func (e CreateImageRequest_MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateImageRequest_Metadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateImageRequest_MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateImageRequest_MetadataValidationError{}

var _CreateImageRequest_Metadata_Ext_InLookup = map[string]struct{}{
	"svg":  {},
	"jpg":  {},
	"jpeg": {},
	"png":  {},
}
