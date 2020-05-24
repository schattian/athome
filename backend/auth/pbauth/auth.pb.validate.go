// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth.proto

package pbauth

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
var _auth_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RetrieveAuthenticationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveAuthenticationRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	return nil
}

// RetrieveAuthenticationRequestValidationError is the validation error
// returned by RetrieveAuthenticationRequest.Validate if the designated
// constraints aren't met.
type RetrieveAuthenticationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveAuthenticationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveAuthenticationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveAuthenticationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveAuthenticationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveAuthenticationRequestValidationError) ErrorName() string {
	return "RetrieveAuthenticationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveAuthenticationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveAuthenticationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveAuthenticationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveAuthenticationRequestValidationError{}

// Validate checks the field values on RetrieveAuthenticationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveAuthenticationResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	return nil
}

// RetrieveAuthenticationResponseValidationError is the validation error
// returned by RetrieveAuthenticationResponse.Validate if the designated
// constraints aren't met.
type RetrieveAuthenticationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveAuthenticationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveAuthenticationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveAuthenticationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveAuthenticationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveAuthenticationResponseValidationError) ErrorName() string {
	return "RetrieveAuthenticationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveAuthenticationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveAuthenticationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveAuthenticationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveAuthenticationResponseValidationError{}

// Validate checks the field values on CreateAuthenticationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAuthenticationRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SignToken

	return nil
}

// CreateAuthenticationRequestValidationError is the validation error returned
// by CreateAuthenticationRequest.Validate if the designated constraints
// aren't met.
type CreateAuthenticationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAuthenticationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAuthenticationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAuthenticationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAuthenticationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAuthenticationRequestValidationError) ErrorName() string {
	return "CreateAuthenticationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAuthenticationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAuthenticationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAuthenticationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAuthenticationRequestValidationError{}

// Validate checks the field values on CreateAuthenticationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAuthenticationResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	// no validation rules for AccessTokenExpNs

	// no validation rules for RefreshTokenExpNs

	return nil
}

// CreateAuthenticationResponseValidationError is the validation error returned
// by CreateAuthenticationResponse.Validate if the designated constraints
// aren't met.
type CreateAuthenticationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAuthenticationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAuthenticationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAuthenticationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAuthenticationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAuthenticationResponseValidationError) ErrorName() string {
	return "CreateAuthenticationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAuthenticationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAuthenticationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAuthenticationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAuthenticationResponseValidationError{}
