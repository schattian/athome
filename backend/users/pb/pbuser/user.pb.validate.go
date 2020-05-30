// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package pbuser

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
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SwitchRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SwitchRoleRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	if _, ok := _SwitchRoleRequest_Role_InLookup[m.GetRole()]; !ok {
		return SwitchRoleRequestValidationError{
			field:  "Role",
			reason: "value must be in list [service-provider consumer merchant]",
		}
	}

	// no validation rules for Password

	return nil
}

// SwitchRoleRequestValidationError is the validation error returned by
// SwitchRoleRequest.Validate if the designated constraints aren't met.
type SwitchRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SwitchRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SwitchRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SwitchRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SwitchRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SwitchRoleRequestValidationError) ErrorName() string {
	return "SwitchRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SwitchRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSwitchRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SwitchRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SwitchRoleRequestValidationError{}

var _SwitchRoleRequest_Role_InLookup = map[string]struct{}{
	"service-provider": {},
	"consumer":         {},
	"merchant":         {},
}

// Validate checks the field values on SignOutRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SignOutRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	return nil
}

// SignOutRequestValidationError is the validation error returned by
// SignOutRequest.Validate if the designated constraints aren't met.
type SignOutRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignOutRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignOutRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignOutRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignOutRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignOutRequestValidationError) ErrorName() string { return "SignOutRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignOutRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignOutRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignOutRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignOutRequestValidationError{}

// Validate checks the field values on ResetPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResetPasswordRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	if l := utf8.RuneCountInString(m.GetPassword()); l < 6 || l > 25 {
		return ResetPasswordRequestValidationError{
			field:  "Password",
			reason: "value length must be between 6 and 25 runes, inclusive",
		}
	}

	return nil
}

// ResetPasswordRequestValidationError is the validation error returned by
// ResetPasswordRequest.Validate if the designated constraints aren't met.
type ResetPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordRequestValidationError) ErrorName() string {
	return "ResetPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ResetPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordRequestValidationError{}

// Validate checks the field values on ForgotPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ForgotPasswordRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return ForgotPasswordRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	return nil
}

func (m *ForgotPasswordRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *ForgotPasswordRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// ForgotPasswordRequestValidationError is the validation error returned by
// ForgotPasswordRequest.Validate if the designated constraints aren't met.
type ForgotPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForgotPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForgotPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForgotPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForgotPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForgotPasswordRequestValidationError) ErrorName() string {
	return "ForgotPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ForgotPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForgotPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForgotPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForgotPasswordRequestValidationError{}

// Validate checks the field values on FetchOnboardingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FetchOnboardingRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OnboardingId

	return nil
}

// FetchOnboardingRequestValidationError is the validation error returned by
// FetchOnboardingRequest.Validate if the designated constraints aren't met.
type FetchOnboardingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchOnboardingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchOnboardingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchOnboardingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchOnboardingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchOnboardingRequestValidationError) ErrorName() string {
	return "FetchOnboardingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FetchOnboardingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchOnboardingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchOnboardingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchOnboardingRequestValidationError{}

// Validate checks the field values on FetchOnboardingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FetchOnboardingResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Role

	// no validation rules for Email

	// no validation rules for Name

	// no validation rules for Surname

	// no validation rules for Stage

	return nil
}

// FetchOnboardingResponseValidationError is the validation error returned by
// FetchOnboardingResponse.Validate if the designated constraints aren't met.
type FetchOnboardingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchOnboardingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchOnboardingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchOnboardingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchOnboardingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchOnboardingResponseValidationError) ErrorName() string {
	return "FetchOnboardingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FetchOnboardingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchOnboardingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchOnboardingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchOnboardingResponseValidationError{}

// Validate checks the field values on SignUpStartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SignUpStartRequest) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _SignUpStartRequest_Role_InLookup[m.GetRole()]; !ok {
		return SignUpStartRequestValidationError{
			field:  "Role",
			reason: "value must be in list [service-provider consumer merchant]",
		}
	}

	return nil
}

// SignUpStartRequestValidationError is the validation error returned by
// SignUpStartRequest.Validate if the designated constraints aren't met.
type SignUpStartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpStartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpStartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpStartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpStartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpStartRequestValidationError) ErrorName() string {
	return "SignUpStartRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SignUpStartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpStartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpStartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpStartRequestValidationError{}

var _SignUpStartRequest_Role_InLookup = map[string]struct{}{
	"service-provider": {},
	"consumer":         {},
	"merchant":         {},
}

// Validate checks the field values on SignUpStartResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SignUpStartResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OnboardingId

	return nil
}

// SignUpStartResponseValidationError is the validation error returned by
// SignUpStartResponse.Validate if the designated constraints aren't met.
type SignUpStartResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpStartResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpStartResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpStartResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpStartResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpStartResponseValidationError) ErrorName() string {
	return "SignUpStartResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SignUpStartResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpStartResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpStartResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpStartResponseValidationError{}

// Validate checks the field values on SignUpSharedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SignUpSharedRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OnboardingId

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return SignUpSharedRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 30 {
		return SignUpSharedRequestValidationError{
			field:  "Name",
			reason: "value length must be between 2 and 30 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetSurname()); l < 2 || l > 30 {
		return SignUpSharedRequestValidationError{
			field:  "Surname",
			reason: "value length must be between 2 and 30 runes, inclusive",
		}
	}

	return nil
}

func (m *SignUpSharedRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *SignUpSharedRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// SignUpSharedRequestValidationError is the validation error returned by
// SignUpSharedRequest.Validate if the designated constraints aren't met.
type SignUpSharedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpSharedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpSharedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpSharedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpSharedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpSharedRequestValidationError) ErrorName() string {
	return "SignUpSharedRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SignUpSharedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpSharedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpSharedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpSharedRequestValidationError{}

// Validate checks the field values on SignUpSharedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SignUpSharedResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OnboardingId

	return nil
}

// SignUpSharedResponseValidationError is the validation error returned by
// SignUpSharedResponse.Validate if the designated constraints aren't met.
type SignUpSharedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpSharedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpSharedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpSharedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpSharedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpSharedResponseValidationError) ErrorName() string {
	return "SignUpSharedResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SignUpSharedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpSharedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpSharedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpSharedResponseValidationError{}

// Validate checks the field values on SignUpEndRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SignUpEndRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OnboardingId

	if l := utf8.RuneCountInString(m.GetPassword()); l < 6 || l > 25 {
		return SignUpEndRequestValidationError{
			field:  "Password",
			reason: "value length must be between 6 and 25 runes, inclusive",
		}
	}

	return nil
}

// SignUpEndRequestValidationError is the validation error returned by
// SignUpEndRequest.Validate if the designated constraints aren't met.
type SignUpEndRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpEndRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpEndRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpEndRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpEndRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpEndRequestValidationError) ErrorName() string { return "SignUpEndRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignUpEndRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpEndRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpEndRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpEndRequestValidationError{}

// Validate checks the field values on SignUpEndResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SignUpEndResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SignUpEndResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SignUpEndResponseValidationError is the validation error returned by
// SignUpEndResponse.Validate if the designated constraints aren't met.
type SignUpEndResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpEndResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpEndResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpEndResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpEndResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpEndResponseValidationError) ErrorName() string {
	return "SignUpEndResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SignUpEndResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUpEndResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpEndResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpEndResponseValidationError{}

// Validate checks the field values on SignInRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SignInRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return SignInRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 6 || l > 25 {
		return SignInRequestValidationError{
			field:  "Password",
			reason: "value length must be between 6 and 25 runes, inclusive",
		}
	}

	return nil
}

func (m *SignInRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *SignInRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// SignInRequestValidationError is the validation error returned by
// SignInRequest.Validate if the designated constraints aren't met.
type SignInRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInRequestValidationError) ErrorName() string { return "SignInRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignInRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInRequestValidationError{}

// Validate checks the field values on SignInResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SignInResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SignInResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for SignTokenExpNs

	return nil
}

// SignInResponseValidationError is the validation error returned by
// SignInResponse.Validate if the designated constraints aren't met.
type SignInResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInResponseValidationError) ErrorName() string { return "SignInResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignInResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInResponseValidationError{}

// Validate checks the field values on SignInUser with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SignInUser) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for SignToken

	// no validation rules for Role

	// no validation rules for Name

	// no validation rules for Surname

	// no validation rules for Email

	return nil
}

// SignInUserValidationError is the validation error returned by
// SignInUser.Validate if the designated constraints aren't met.
type SignInUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignInUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignInUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignInUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignInUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignInUserValidationError) ErrorName() string { return "SignInUserValidationError" }

// Error satisfies the builtin error interface
func (e SignInUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignInUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignInUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignInUserValidationError{}

// Validate checks the field values on SignRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SignRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SignToken

	return nil
}

// SignRequestValidationError is the validation error returned by
// SignRequest.Validate if the designated constraints aren't met.
type SignRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignRequestValidationError) ErrorName() string { return "SignRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignRequestValidationError{}

// Validate checks the field values on SignResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SignResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	// no validation rules for AccessTokenExpNs

	// no validation rules for RefreshTokenExpNs

	return nil
}

// SignResponseValidationError is the validation error returned by
// SignResponse.Validate if the designated constraints aren't met.
type SignResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignResponseValidationError) ErrorName() string { return "SignResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignResponseValidationError{}
