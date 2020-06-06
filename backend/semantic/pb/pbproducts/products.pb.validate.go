// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: products.proto

package pbproducts

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
var _products_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DraftLine with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DraftLine) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DraftLineId

	if v, ok := interface{}(m.GetFirst()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DraftLineValidationError{
				field:  "First",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSecond()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DraftLineValidationError{
				field:  "Second",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DraftLineValidationError is the validation error returned by
// DraftLine.Validate if the designated constraints aren't met.
type DraftLineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DraftLineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DraftLineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DraftLineValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DraftLineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DraftLineValidationError) ErrorName() string { return "DraftLineValidationError" }

// Error satisfies the builtin error interface
func (e DraftLineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDraftLine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DraftLineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DraftLineValidationError{}

// Validate checks the field values on DraftLineFirst with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DraftLineFirst) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Title

	// no validation rules for CategoryId

	return nil
}

// DraftLineFirstValidationError is the validation error returned by
// DraftLineFirst.Validate if the designated constraints aren't met.
type DraftLineFirstValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DraftLineFirstValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DraftLineFirstValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DraftLineFirstValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DraftLineFirstValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DraftLineFirstValidationError) ErrorName() string { return "DraftLineFirstValidationError" }

// Error satisfies the builtin error interface
func (e DraftLineFirstValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDraftLineFirst.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DraftLineFirstValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DraftLineFirstValidationError{}

// Validate checks the field values on DraftLineSecond with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DraftLineSecond) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Price

	// no validation rules for Stock

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DraftLineSecondValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DraftLineSecondValidationError is the validation error returned by
// DraftLineSecond.Validate if the designated constraints aren't met.
type DraftLineSecondValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DraftLineSecondValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DraftLineSecondValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DraftLineSecondValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DraftLineSecondValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DraftLineSecondValidationError) ErrorName() string { return "DraftLineSecondValidationError" }

// Error satisfies the builtin error interface
func (e DraftLineSecondValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDraftLineSecond.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DraftLineSecondValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DraftLineSecondValidationError{}

// Validate checks the field values on AttributeData with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AttributeData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

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

// Validate checks the field values on SecondRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SecondRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for DraftLineId

	if v, ok := interface{}(m.GetDraftLine()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SecondRequestValidationError{
				field:  "DraftLine",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SecondRequestValidationError is the validation error returned by
// SecondRequest.Validate if the designated constraints aren't met.
type SecondRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecondRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecondRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecondRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecondRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecondRequestValidationError) ErrorName() string { return "SecondRequestValidationError" }

// Error satisfies the builtin error interface
func (e SecondRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecondRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecondRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecondRequestValidationError{}

// Validate checks the field values on Page with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Page) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Cursor

	// no validation rules for Size

	return nil
}

// PageValidationError is the validation error returned by Page.Validate if the
// designated constraints aren't met.
type PageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageValidationError) ErrorName() string { return "PageValidationError" }

// Error satisfies the builtin error interface
func (e PageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageValidationError{}

// Validate checks the field values on FirstRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FirstRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	if v, ok := interface{}(m.GetDraftLine()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FirstRequestValidationError{
				field:  "DraftLine",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FirstRequestValidationError is the validation error returned by
// FirstRequest.Validate if the designated constraints aren't met.
type FirstRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FirstRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FirstRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FirstRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FirstRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FirstRequestValidationError) ErrorName() string { return "FirstRequestValidationError" }

// Error satisfies the builtin error interface
func (e FirstRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFirstRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FirstRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FirstRequestValidationError{}

// Validate checks the field values on FirstResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FirstResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DraftId

	return nil
}

// FirstResponseValidationError is the validation error returned by
// FirstResponse.Validate if the designated constraints aren't met.
type FirstResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FirstResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FirstResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FirstResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FirstResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FirstResponseValidationError) ErrorName() string { return "FirstResponseValidationError" }

// Error satisfies the builtin error interface
func (e FirstResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFirstResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FirstResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FirstResponseValidationError{}

// Validate checks the field values on CloneDraftLineRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CloneDraftLineRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for DraftLineId

	return nil
}

// CloneDraftLineRequestValidationError is the validation error returned by
// CloneDraftLineRequest.Validate if the designated constraints aren't met.
type CloneDraftLineRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CloneDraftLineRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CloneDraftLineRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CloneDraftLineRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CloneDraftLineRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CloneDraftLineRequestValidationError) ErrorName() string {
	return "CloneDraftLineRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CloneDraftLineRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCloneDraftLineRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CloneDraftLineRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CloneDraftLineRequestValidationError{}

// Validate checks the field values on CloneDraftLineResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CloneDraftLineResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDraftLine()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CloneDraftLineResponseValidationError{
				field:  "DraftLine",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CloneDraftLineResponseValidationError is the validation error returned by
// CloneDraftLineResponse.Validate if the designated constraints aren't met.
type CloneDraftLineResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CloneDraftLineResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CloneDraftLineResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CloneDraftLineResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CloneDraftLineResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CloneDraftLineResponseValidationError) ErrorName() string {
	return "CloneDraftLineResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CloneDraftLineResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCloneDraftLineResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CloneDraftLineResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CloneDraftLineResponseValidationError{}

// Validate checks the field values on FetchDraftRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FetchDraftRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	return nil
}

// FetchDraftRequestValidationError is the validation error returned by
// FetchDraftRequest.Validate if the designated constraints aren't met.
type FetchDraftRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchDraftRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchDraftRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchDraftRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchDraftRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchDraftRequestValidationError) ErrorName() string {
	return "FetchDraftRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FetchDraftRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchDraftRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchDraftRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchDraftRequestValidationError{}

// Validate checks the field values on FetchDraftResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FetchDraftResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DraftId

	// no validation rules for Stage

	for idx, item := range m.GetLines() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FetchDraftResponseValidationError{
					field:  fmt.Sprintf("Lines[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FetchDraftResponseValidationError is the validation error returned by
// FetchDraftResponse.Validate if the designated constraints aren't met.
type FetchDraftResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchDraftResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchDraftResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchDraftResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchDraftResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchDraftResponseValidationError) ErrorName() string {
	return "FetchDraftResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FetchDraftResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchDraftResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchDraftResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchDraftResponseValidationError{}
