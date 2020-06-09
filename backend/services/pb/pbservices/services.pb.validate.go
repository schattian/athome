// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: services.proto

package pbservices

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
var _services_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RetrieveRegistryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveRegistryRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	if v, ok := interface{}(m.GetFirst()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RetrieveRegistryRequestValidationError{
				field:  "First",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSecond()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RetrieveRegistryRequestValidationError{
				field:  "Second",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetThird()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RetrieveRegistryRequestValidationError{
				field:  "Third",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RetrieveRegistryRequestValidationError is the validation error returned by
// RetrieveRegistryRequest.Validate if the designated constraints aren't met.
type RetrieveRegistryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveRegistryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveRegistryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveRegistryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveRegistryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveRegistryRequestValidationError) ErrorName() string {
	return "RetrieveRegistryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveRegistryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveRegistryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveRegistryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveRegistryRequestValidationError{}

// Validate checks the field values on RetrieveRegistryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveRegistryResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RegistryId

	// no validation rules for Stage

	return nil
}

// RetrieveRegistryResponseValidationError is the validation error returned by
// RetrieveRegistryResponse.Validate if the designated constraints aren't met.
type RetrieveRegistryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveRegistryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveRegistryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveRegistryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveRegistryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveRegistryResponseValidationError) ErrorName() string {
	return "RetrieveRegistryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveRegistryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveRegistryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveRegistryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveRegistryResponseValidationError{}

// Validate checks the field values on FirstRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FirstRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	if v, ok := interface{}(m.GetBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FirstRequestValidationError{
				field:  "Body",
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

	// no validation rules for RegistryId

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

// Validate checks the field values on SecondRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SecondRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	if v, ok := interface{}(m.GetBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SecondRequestValidationError{
				field:  "Body",
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

// Validate checks the field values on ThirdRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ThirdRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	if v, ok := interface{}(m.GetBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ThirdRequestValidationError{
				field:  "Body",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ThirdRequestValidationError is the validation error returned by
// ThirdRequest.Validate if the designated constraints aren't met.
type ThirdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThirdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThirdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThirdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThirdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThirdRequestValidationError) ErrorName() string { return "ThirdRequestValidationError" }

// Error satisfies the builtin error interface
func (e ThirdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThirdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThirdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThirdRequestValidationError{}

// Validate checks the field values on CreateCalendarRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateCalendarRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for Name

	// no validation rules for GroupId

	for idx, item := range m.GetAvailabilities() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateCalendarRequestValidationError{
					field:  fmt.Sprintf("Availabilities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateCalendarRequestValidationError is the validation error returned by
// CreateCalendarRequest.Validate if the designated constraints aren't met.
type CreateCalendarRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCalendarRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCalendarRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCalendarRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCalendarRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCalendarRequestValidationError) ErrorName() string {
	return "CreateCalendarRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCalendarRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCalendarRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCalendarRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCalendarRequestValidationError{}

// Validate checks the field values on CreateCalendarResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateCalendarResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CalendarId

	// no validation rules for Name

	// no validation rules for GroupId

	// no validation rules for UserId

	for idx, item := range m.GetAvailabilities() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateCalendarResponseValidationError{
					field:  fmt.Sprintf("Availabilities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateCalendarResponseValidationError is the validation error returned by
// CreateCalendarResponse.Validate if the designated constraints aren't met.
type CreateCalendarResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCalendarResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCalendarResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCalendarResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCalendarResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCalendarResponseValidationError) ErrorName() string {
	return "CreateCalendarResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCalendarResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCalendarResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCalendarResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCalendarResponseValidationError{}

// Validate checks the field values on TimeOfDay with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TimeOfDay) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Hour

	// no validation rules for Minute

	return nil
}

// TimeOfDayValidationError is the validation error returned by
// TimeOfDay.Validate if the designated constraints aren't met.
type TimeOfDayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimeOfDayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimeOfDayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimeOfDayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimeOfDayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimeOfDayValidationError) ErrorName() string { return "TimeOfDayValidationError" }

// Error satisfies the builtin error interface
func (e TimeOfDayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimeOfDay.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimeOfDayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimeOfDayValidationError{}

// Validate checks the field values on CreateAvailability with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAvailability) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _CreateAvailability_Dow_InLookup[m.GetDow()]; !ok {
		return CreateAvailabilityValidationError{
			field:  "Dow",
			reason: "value must be in list [monday tuesday wednesday thursday friday saturday sunday]",
		}
	}

	if v, ok := interface{}(m.GetStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAvailabilityValidationError{
				field:  "Start",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAvailabilityValidationError{
				field:  "End",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateAvailabilityValidationError is the validation error returned by
// CreateAvailability.Validate if the designated constraints aren't met.
type CreateAvailabilityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAvailabilityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAvailabilityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAvailabilityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAvailabilityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAvailabilityValidationError) ErrorName() string {
	return "CreateAvailabilityValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAvailabilityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAvailability.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAvailabilityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAvailabilityValidationError{}

var _CreateAvailability_Dow_InLookup = map[string]struct{}{
	"monday":    {},
	"tuesday":   {},
	"wednesday": {},
	"thursday":  {},
	"friday":    {},
	"saturday":  {},
	"sunday":    {},
}

// Validate checks the field values on RetrieveAvailability with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveAvailability) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AvailabilityId

	if _, ok := _RetrieveAvailability_Dow_InLookup[m.GetDow()]; !ok {
		return RetrieveAvailabilityValidationError{
			field:  "Dow",
			reason: "value must be in list [monday tuesday wednesday thursday friday saturday sunday]",
		}
	}

	if v, ok := interface{}(m.GetStart()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RetrieveAvailabilityValidationError{
				field:  "Start",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RetrieveAvailabilityValidationError{
				field:  "End",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RetrieveAvailabilityValidationError is the validation error returned by
// RetrieveAvailability.Validate if the designated constraints aren't met.
type RetrieveAvailabilityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveAvailabilityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveAvailabilityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveAvailabilityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveAvailabilityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveAvailabilityValidationError) ErrorName() string {
	return "RetrieveAvailabilityValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveAvailabilityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveAvailability.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveAvailabilityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveAvailabilityValidationError{}

var _RetrieveAvailability_Dow_InLookup = map[string]struct{}{
	"monday":    {},
	"tuesday":   {},
	"wednesday": {},
	"thursday":  {},
	"friday":    {},
	"saturday":  {},
	"sunday":    {},
}

// Validate checks the field values on Price with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Price) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Min

	// no validation rules for Max

	return nil
}

// PriceValidationError is the validation error returned by Price.Validate if
// the designated constraints aren't met.
type PriceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PriceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PriceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PriceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PriceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PriceValidationError) ErrorName() string { return "PriceValidationError" }

// Error satisfies the builtin error interface
func (e PriceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrice.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PriceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PriceValidationError{}

// Validate checks the field values on FirstRequest_Body with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FirstRequest_Body) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AddressId

	return nil
}

// FirstRequest_BodyValidationError is the validation error returned by
// FirstRequest_Body.Validate if the designated constraints aren't met.
type FirstRequest_BodyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FirstRequest_BodyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FirstRequest_BodyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FirstRequest_BodyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FirstRequest_BodyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FirstRequest_BodyValidationError) ErrorName() string {
	return "FirstRequest_BodyValidationError"
}

// Error satisfies the builtin error interface
func (e FirstRequest_BodyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFirstRequest_Body.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FirstRequest_BodyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FirstRequest_BodyValidationError{}

// Validate checks the field values on SecondRequest_Body with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SecondRequest_Body) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for DurationInMinutes

	if v, ok := interface{}(m.GetPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SecondRequest_BodyValidationError{
				field:  "Price",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SecondRequest_BodyValidationError is the validation error returned by
// SecondRequest_Body.Validate if the designated constraints aren't met.
type SecondRequest_BodyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecondRequest_BodyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecondRequest_BodyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecondRequest_BodyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecondRequest_BodyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecondRequest_BodyValidationError) ErrorName() string {
	return "SecondRequest_BodyValidationError"
}

// Error satisfies the builtin error interface
func (e SecondRequest_BodyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecondRequest_Body.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecondRequest_BodyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecondRequest_BodyValidationError{}

// Validate checks the field values on ThirdRequest_Body with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ThirdRequest_Body) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CalendarId

	return nil
}

// ThirdRequest_BodyValidationError is the validation error returned by
// ThirdRequest_Body.Validate if the designated constraints aren't met.
type ThirdRequest_BodyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThirdRequest_BodyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThirdRequest_BodyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThirdRequest_BodyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThirdRequest_BodyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThirdRequest_BodyValidationError) ErrorName() string {
	return "ThirdRequest_BodyValidationError"
}

// Error satisfies the builtin error interface
func (e ThirdRequest_BodyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThirdRequest_Body.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThirdRequest_BodyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThirdRequest_BodyValidationError{}