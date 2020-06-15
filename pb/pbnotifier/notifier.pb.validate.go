// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: notifier.proto

package pbnotifier

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
var _notifier_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on UpdateStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateStatusRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for NotificationId

	return nil
}

// UpdateStatusRequestValidationError is the validation error returned by
// UpdateStatusRequest.Validate if the designated constraints aren't met.
type UpdateStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateStatusRequestValidationError) ErrorName() string {
	return "UpdateStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateStatusRequestValidationError{}

// Validate checks the field values on CreateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NotificationToken

	if v, ok := interface{}(m.GetNotification()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRequestValidationError{
				field:  "Notification",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateRequestValidationError is the validation error returned by
// CreateRequest.Validate if the designated constraints aren't met.
type CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestValidationError) ErrorName() string { return "CreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestValidationError{}

// Validate checks the field values on CreateResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NotificationId

	if v, ok := interface{}(m.GetNotification()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateResponseValidationError{
				field:  "Notification",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateResponseValidationError is the validation error returned by
// CreateResponse.Validate if the designated constraints aren't met.
type CreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResponseValidationError) ErrorName() string { return "CreateResponseValidationError" }

// Error satisfies the builtin error interface
func (e CreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResponseValidationError{}

// Validate checks the field values on RetrieveRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RetrieveRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for NotificationId

	return nil
}

// RetrieveRequestValidationError is the validation error returned by
// RetrieveRequest.Validate if the designated constraints aren't met.
type RetrieveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveRequestValidationError) ErrorName() string { return "RetrieveRequestValidationError" }

// Error satisfies the builtin error interface
func (e RetrieveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveRequestValidationError{}

// Validate checks the field values on RetrieveStreamRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveStreamRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for TickerMs

	return nil
}

// RetrieveStreamRequestValidationError is the validation error returned by
// RetrieveStreamRequest.Validate if the designated constraints aren't met.
type RetrieveStreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveStreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveStreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveStreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveStreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveStreamRequestValidationError) ErrorName() string {
	return "RetrieveStreamRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveStreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveStreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveStreamRequestValidationError{}

// Validate checks the field values on RetrieveManyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveManyRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	return nil
}

// RetrieveManyRequestValidationError is the validation error returned by
// RetrieveManyRequest.Validate if the designated constraints aren't met.
type RetrieveManyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveManyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveManyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveManyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveManyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveManyRequestValidationError) ErrorName() string {
	return "RetrieveManyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveManyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveManyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveManyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveManyRequestValidationError{}

// Validate checks the field values on RetrieveManyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveManyResponse) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetNotifications() {
		_ = val

		// no validation rules for Notifications[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RetrieveManyResponseValidationError{
					field:  fmt.Sprintf("Notifications[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RetrieveManyResponseValidationError is the validation error returned by
// RetrieveManyResponse.Validate if the designated constraints aren't met.
type RetrieveManyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveManyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveManyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveManyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveManyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveManyResponseValidationError) ErrorName() string {
	return "RetrieveManyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveManyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveManyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveManyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveManyResponseValidationError{}

// Validate checks the field values on Notification with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Notification) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	if v, ok := interface{}(m.GetEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NotificationValidationError{
				field:  "Entity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NotificationValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := _Notification_Priority_InLookup[m.GetPriority()]; !ok {
		return NotificationValidationError{
			field:  "Priority",
			reason: "value must be in list [low mid high max]",
		}
	}

	return nil
}

// NotificationValidationError is the validation error returned by
// Notification.Validate if the designated constraints aren't met.
type NotificationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationValidationError) ErrorName() string { return "NotificationValidationError" }

// Error satisfies the builtin error interface
func (e NotificationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotification.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationValidationError{}

var _Notification_Priority_InLookup = map[string]struct{}{
	"low":  {},
	"mid":  {},
	"high": {},
	"max":  {},
}

// Validate checks the field values on Status with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Status) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetReceivedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusValidationError{
				field:  "ReceivedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSeenAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatusValidationError{
				field:  "SeenAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// StatusValidationError is the validation error returned by Status.Validate if
// the designated constraints aren't met.
type StatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusValidationError) ErrorName() string { return "StatusValidationError" }

// Error satisfies the builtin error interface
func (e StatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusValidationError{}

// Validate checks the field values on Entity with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Entity) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// EntityValidationError is the validation error returned by Entity.Validate if
// the designated constraints aren't met.
type EntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityValidationError) ErrorName() string { return "EntityValidationError" }

// Error satisfies the builtin error interface
func (e EntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityValidationError{}
