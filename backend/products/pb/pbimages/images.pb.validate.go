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

// Validate checks the field values on CreateImageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateImageRequest) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Corpus.(type) {

	case *CreateImageRequest_Metadata:

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

// Validate checks the field values on CreateImageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateImageResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Size

	// no validation rules for Uri

	// no validation rules for UserId

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

	for idx, item := range m.GetImages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RetrieveImagesResponseValidationError{
					field:  fmt.Sprintf("Images[%v]", idx),
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

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Metadata) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _Metadata_Ext_InLookup[m.GetExt()]; !ok {
		return MetadataValidationError{
			field:  "Ext",
			reason: "value must be in list [svg jpg jpeg png]",
		}
	}

	// no validation rules for AccessToken

	return nil
}

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

var _Metadata_Ext_InLookup = map[string]struct{}{
	"svg":  {},
	"jpg":  {},
	"jpeg": {},
	"png":  {},
}

// Validate checks the field values on RetrieveImagesResponse_Data with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveImagesResponse_Data) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Uri

	// no validation rules for UserId

	return nil
}

// RetrieveImagesResponse_DataValidationError is the validation error returned
// by RetrieveImagesResponse_Data.Validate if the designated constraints
// aren't met.
type RetrieveImagesResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveImagesResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveImagesResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveImagesResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveImagesResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveImagesResponse_DataValidationError) ErrorName() string {
	return "RetrieveImagesResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveImagesResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveImagesResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveImagesResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveImagesResponse_DataValidationError{}