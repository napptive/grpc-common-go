// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/entities.proto

package grpc_common_go

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
var _entities_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on EmptyRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EmptyRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// EmptyRequestValidationError is the validation error returned by
// EmptyRequest.Validate if the designated constraints aren't met.
type EmptyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyRequestValidationError) ErrorName() string { return "EmptyRequestValidationError" }

// Error satisfies the builtin error interface
func (e EmptyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyRequestValidationError{}

// Validate checks the field values on OpResponse with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OpResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for UserInfo

	return nil
}

// OpResponseValidationError is the validation error returned by
// OpResponse.Validate if the designated constraints aren't met.
type OpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpResponseValidationError) ErrorName() string { return "OpResponseValidationError" }

// Error satisfies the builtin error interface
func (e OpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpResponseValidationError{}

// Validate checks the field values on ErrorDetails with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ErrorDetails) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Detail

	return nil
}

// ErrorDetailsValidationError is the validation error returned by
// ErrorDetails.Validate if the designated constraints aren't met.
type ErrorDetailsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorDetailsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorDetailsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorDetailsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorDetailsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorDetailsValidationError) ErrorName() string { return "ErrorDetailsValidationError" }

// Error satisfies the builtin error interface
func (e ErrorDetailsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorDetails.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorDetailsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorDetailsValidationError{}
