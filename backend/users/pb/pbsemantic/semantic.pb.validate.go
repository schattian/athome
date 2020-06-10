// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: semantic.proto

package pbsemantic

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
var _semantic_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RetrieveCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveCategoryRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CategoryId

	return nil
}

// RetrieveCategoryRequestValidationError is the validation error returned by
// RetrieveCategoryRequest.Validate if the designated constraints aren't met.
type RetrieveCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveCategoryRequestValidationError) ErrorName() string {
	return "RetrieveCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveCategoryRequestValidationError{}

// Validate checks the field values on Category with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Category) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for ParentId

	for key, val := range m.GetChilds() {
		_ = val

		// no validation rules for Childs[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CategoryValidationError{
					field:  fmt.Sprintf("Childs[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for IdentificationTemplate

	return nil
}

// CategoryValidationError is the validation error returned by
// Category.Validate if the designated constraints aren't met.
type CategoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryValidationError) ErrorName() string { return "CategoryValidationError" }

// Error satisfies the builtin error interface
func (e CategoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryValidationError{}

// Validate checks the field values on AttributeSchema with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AttributeSchema) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CategoryId

	// no validation rules for Name

	// no validation rules for ValueType

	return nil
}

// AttributeSchemaValidationError is the validation error returned by
// AttributeSchema.Validate if the designated constraints aren't met.
type AttributeSchemaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeSchemaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeSchemaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeSchemaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeSchemaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeSchemaValidationError) ErrorName() string { return "AttributeSchemaValidationError" }

// Error satisfies the builtin error interface
func (e AttributeSchemaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeSchema.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeSchemaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeSchemaValidationError{}

// Validate checks the field values on RetrieveCategoriesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveCategoriesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetCategories() {
		_ = val

		// no validation rules for Categories[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RetrieveCategoriesResponseValidationError{
					field:  fmt.Sprintf("Categories[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RetrieveCategoriesResponseValidationError is the validation error returned
// by RetrieveCategoriesResponse.Validate if the designated constraints aren't met.
type RetrieveCategoriesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveCategoriesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveCategoriesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveCategoriesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveCategoriesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveCategoriesResponseValidationError) ErrorName() string {
	return "RetrieveCategoriesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveCategoriesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveCategoriesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveCategoriesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveCategoriesResponseValidationError{}

// Validate checks the field values on RetrieveAttributesSchemaRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveAttributesSchemaRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CategoryId

	return nil
}

// RetrieveAttributesSchemaRequestValidationError is the validation error
// returned by RetrieveAttributesSchemaRequest.Validate if the designated
// constraints aren't met.
type RetrieveAttributesSchemaRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveAttributesSchemaRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveAttributesSchemaRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveAttributesSchemaRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveAttributesSchemaRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveAttributesSchemaRequestValidationError) ErrorName() string {
	return "RetrieveAttributesSchemaRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveAttributesSchemaRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveAttributesSchemaRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveAttributesSchemaRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveAttributesSchemaRequestValidationError{}

// Validate checks the field values on RetrieveAttributesDataRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveAttributesDataRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// RetrieveAttributesDataRequestValidationError is the validation error
// returned by RetrieveAttributesDataRequest.Validate if the designated
// constraints aren't met.
type RetrieveAttributesDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveAttributesDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveAttributesDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveAttributesDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveAttributesDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveAttributesDataRequestValidationError) ErrorName() string {
	return "RetrieveAttributesDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveAttributesDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveAttributesDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveAttributesDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveAttributesDataRequestValidationError{}

// Validate checks the field values on DeleteAttributesDataRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteAttributesDataRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// DeleteAttributesDataRequestValidationError is the validation error returned
// by DeleteAttributesDataRequest.Validate if the designated constraints
// aren't met.
type DeleteAttributesDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAttributesDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAttributesDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAttributesDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAttributesDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAttributesDataRequestValidationError) ErrorName() string {
	return "DeleteAttributesDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAttributesDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAttributesDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAttributesDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAttributesDataRequestValidationError{}

// Validate checks the field values on RetrieveAttributesSchemaResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *RetrieveAttributesSchemaResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RetrieveAttributesSchemaResponseValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RetrieveAttributesSchemaResponseValidationError is the validation error
// returned by RetrieveAttributesSchemaResponse.Validate if the designated
// constraints aren't met.
type RetrieveAttributesSchemaResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveAttributesSchemaResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveAttributesSchemaResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveAttributesSchemaResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveAttributesSchemaResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveAttributesSchemaResponseValidationError) ErrorName() string {
	return "RetrieveAttributesSchemaResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveAttributesSchemaResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveAttributesSchemaResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveAttributesSchemaResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveAttributesSchemaResponseValidationError{}

// Validate checks the field values on RetrieveAttributesDataResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveAttributesDataResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RetrieveAttributesDataResponseValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RetrieveAttributesDataResponseValidationError is the validation error
// returned by RetrieveAttributesDataResponse.Validate if the designated
// constraints aren't met.
type RetrieveAttributesDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveAttributesDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveAttributesDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveAttributesDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveAttributesDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveAttributesDataResponseValidationError) ErrorName() string {
	return "RetrieveAttributesDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveAttributesDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveAttributesDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveAttributesDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveAttributesDataResponseValidationError{}

// Validate checks the field values on SetAttributesDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SetAttributesDataResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AttributeDataId

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetAttributesDataResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SetAttributesDataResponseValidationError is the validation error returned by
// SetAttributesDataResponse.Validate if the designated constraints aren't met.
type SetAttributesDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetAttributesDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetAttributesDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetAttributesDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetAttributesDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetAttributesDataResponseValidationError) ErrorName() string {
	return "SetAttributesDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SetAttributesDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetAttributesDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetAttributesDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetAttributesDataResponseValidationError{}

// Validate checks the field values on CloneAttributesDataResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CloneAttributesDataResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CloneAttributesDataResponseValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CloneAttributesDataResponseValidationError is the validation error returned
// by CloneAttributesDataResponse.Validate if the designated constraints
// aren't met.
type CloneAttributesDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CloneAttributesDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CloneAttributesDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CloneAttributesDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CloneAttributesDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CloneAttributesDataResponseValidationError) ErrorName() string {
	return "CloneAttributesDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CloneAttributesDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCloneAttributesDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CloneAttributesDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CloneAttributesDataResponseValidationError{}

// Validate checks the field values on SetAttributesDataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SetAttributesDataRequest) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Corpus.(type) {

	case *SetAttributesDataRequest_Authorization_:

		if v, ok := interface{}(m.GetAuthorization()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SetAttributesDataRequestValidationError{
					field:  "Authorization",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *SetAttributesDataRequest_Data:

		if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SetAttributesDataRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// SetAttributesDataRequestValidationError is the validation error returned by
// SetAttributesDataRequest.Validate if the designated constraints aren't met.
type SetAttributesDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetAttributesDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetAttributesDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetAttributesDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetAttributesDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetAttributesDataRequestValidationError) ErrorName() string {
	return "SetAttributesDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetAttributesDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetAttributesDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetAttributesDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetAttributesDataRequestValidationError{}

// Validate checks the field values on ChangeEntityAttributesDataRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ChangeEntityAttributesDataRequest) Validate() error {
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

// ChangeEntityAttributesDataRequestValidationError is the validation error
// returned by ChangeEntityAttributesDataRequest.Validate if the designated
// constraints aren't met.
type ChangeEntityAttributesDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChangeEntityAttributesDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChangeEntityAttributesDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChangeEntityAttributesDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChangeEntityAttributesDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChangeEntityAttributesDataRequestValidationError) ErrorName() string {
	return "ChangeEntityAttributesDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChangeEntityAttributesDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChangeEntityAttributesDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChangeEntityAttributesDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChangeEntityAttributesDataRequestValidationError{}

// Validate checks the field values on CloneAttributesDataRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CloneAttributesDataRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for FromEntityId

	// no validation rules for DestEntityId

	// no validation rules for EntityTable

	return nil
}

// CloneAttributesDataRequestValidationError is the validation error returned
// by CloneAttributesDataRequest.Validate if the designated constraints aren't met.
type CloneAttributesDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CloneAttributesDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CloneAttributesDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CloneAttributesDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CloneAttributesDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CloneAttributesDataRequestValidationError) ErrorName() string {
	return "CloneAttributesDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CloneAttributesDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCloneAttributesDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CloneAttributesDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CloneAttributesDataRequestValidationError{}

// Validate checks the field values on AttributeData with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AttributeData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SchemaId

	return nil
}

// AttributeDataValidationError is the validation error returned by
// AttributeData.Validate if the designated constraints aren't met.
type AttributeDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeDataValidationError) ErrorName() string { return "AttributeDataValidationError" }

// Error satisfies the builtin error interface
func (e AttributeDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeDataValidationError{}

// Validate checks the field values on PredictCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PredictCategoryRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Title

	return nil
}

// PredictCategoryRequestValidationError is the validation error returned by
// PredictCategoryRequest.Validate if the designated constraints aren't met.
type PredictCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PredictCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PredictCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PredictCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PredictCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PredictCategoryRequestValidationError) ErrorName() string {
	return "PredictCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PredictCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPredictCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PredictCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PredictCategoryRequestValidationError{}

// Validate checks the field values on PredictCategoryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PredictCategoryResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CategoryId

	if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PredictCategoryResponseValidationError{
				field:  "Category",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Score

	return nil
}

// PredictCategoryResponseValidationError is the validation error returned by
// PredictCategoryResponse.Validate if the designated constraints aren't met.
type PredictCategoryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PredictCategoryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PredictCategoryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PredictCategoryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PredictCategoryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PredictCategoryResponseValidationError) ErrorName() string {
	return "PredictCategoryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PredictCategoryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPredictCategoryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PredictCategoryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PredictCategoryResponseValidationError{}

// Validate checks the field values on RetrieveCategoriesByRoleRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveCategoriesByRoleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _RetrieveCategoriesByRoleRequest_Role_InLookup[m.GetRole()]; !ok {
		return RetrieveCategoriesByRoleRequestValidationError{
			field:  "Role",
			reason: "value must be in list [service-provider consumer merchant]",
		}
	}

	return nil
}

// RetrieveCategoriesByRoleRequestValidationError is the validation error
// returned by RetrieveCategoriesByRoleRequest.Validate if the designated
// constraints aren't met.
type RetrieveCategoriesByRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveCategoriesByRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveCategoriesByRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveCategoriesByRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveCategoriesByRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveCategoriesByRoleRequestValidationError) ErrorName() string {
	return "RetrieveCategoriesByRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveCategoriesByRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveCategoriesByRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveCategoriesByRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveCategoriesByRoleRequestValidationError{}

var _RetrieveCategoriesByRoleRequest_Role_InLookup = map[string]struct{}{
	"service-provider": {},
	"consumer":         {},
	"merchant":         {},
}

// Validate checks the field values on SetAttributesDataRequest_Authorization
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *SetAttributesDataRequest_Authorization) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// SetAttributesDataRequest_AuthorizationValidationError is the validation
// error returned by SetAttributesDataRequest_Authorization.Validate if the
// designated constraints aren't met.
type SetAttributesDataRequest_AuthorizationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetAttributesDataRequest_AuthorizationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetAttributesDataRequest_AuthorizationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetAttributesDataRequest_AuthorizationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetAttributesDataRequest_AuthorizationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetAttributesDataRequest_AuthorizationValidationError) ErrorName() string {
	return "SetAttributesDataRequest_AuthorizationValidationError"
}

// Error satisfies the builtin error interface
func (e SetAttributesDataRequest_AuthorizationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetAttributesDataRequest_Authorization.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetAttributesDataRequest_AuthorizationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetAttributesDataRequest_AuthorizationValidationError{}
