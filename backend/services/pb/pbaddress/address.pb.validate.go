// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: address.proto

package pbaddress

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
var _address_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Address) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Country

	// no validation rules for Province

	// no validation rules for Zipcode

	// no validation rules for Street

	// no validation rules for Number

	// no validation rules for Floor

	// no validation rules for Department

	// no validation rules for Latitude

	// no validation rules for Longitude

	// no validation rules for Alias

	// no validation rules for UserId

	return nil
}

// AddressValidationError is the validation error returned by Address.Validate
// if the designated constraints aren't met.
type AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressValidationError) ErrorName() string { return "AddressValidationError" }

// Error satisfies the builtin error interface
func (e AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressValidationError{}

// Validate checks the field values on CreateAddressRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAddressRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	if v, ok := interface{}(m.GetBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAddressRequestValidationError{
				field:  "Body",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateAddressRequestValidationError is the validation error returned by
// CreateAddressRequest.Validate if the designated constraints aren't met.
type CreateAddressRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAddressRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAddressRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAddressRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAddressRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAddressRequestValidationError) ErrorName() string {
	return "CreateAddressRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAddressRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAddressRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAddressRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAddressRequestValidationError{}

// Validate checks the field values on CreateAddressResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAddressResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AddressId

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAddressResponseValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateAddressResponseValidationError is the validation error returned by
// CreateAddressResponse.Validate if the designated constraints aren't met.
type CreateAddressResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAddressResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAddressResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAddressResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAddressResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAddressResponseValidationError) ErrorName() string {
	return "CreateAddressResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAddressResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAddressResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAddressResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAddressResponseValidationError{}

// Validate checks the field values on RetrieveAddressRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveAddressRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AddressId

	return nil
}

// RetrieveAddressRequestValidationError is the validation error returned by
// RetrieveAddressRequest.Validate if the designated constraints aren't met.
type RetrieveAddressRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveAddressRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveAddressRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveAddressRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveAddressRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveAddressRequestValidationError) ErrorName() string {
	return "RetrieveAddressRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveAddressRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveAddressRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveAddressRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveAddressRequestValidationError{}

// Validate checks the field values on RetrieveMyAddressesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveMyAddressesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	return nil
}

// RetrieveMyAddressesRequestValidationError is the validation error returned
// by RetrieveMyAddressesRequest.Validate if the designated constraints aren't met.
type RetrieveMyAddressesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveMyAddressesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveMyAddressesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveMyAddressesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveMyAddressesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveMyAddressesRequestValidationError) ErrorName() string {
	return "RetrieveMyAddressesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveMyAddressesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveMyAddressesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveMyAddressesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveMyAddressesRequestValidationError{}

// Validate checks the field values on RetrieveMyAddressesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetrieveMyAddressesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetAddresses() {
		_ = val

		// no validation rules for Addresses[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RetrieveMyAddressesResponseValidationError{
					field:  fmt.Sprintf("Addresses[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RetrieveMyAddressesResponseValidationError is the validation error returned
// by RetrieveMyAddressesResponse.Validate if the designated constraints
// aren't met.
type RetrieveMyAddressesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetrieveMyAddressesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetrieveMyAddressesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetrieveMyAddressesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetrieveMyAddressesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetrieveMyAddressesResponseValidationError) ErrorName() string {
	return "RetrieveMyAddressesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RetrieveMyAddressesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetrieveMyAddressesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetrieveMyAddressesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetrieveMyAddressesResponseValidationError{}
