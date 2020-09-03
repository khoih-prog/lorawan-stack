// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _message_services_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on EncodeDownlinkMessageRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *EncodeDownlinkMessageRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = EncodeDownlinkMessageRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.EndDeviceIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return EncodeDownlinkMessageRequestValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "end_device_version_ids":

			if v, ok := interface{}(&m.EndDeviceVersionIDs).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return EncodeDownlinkMessageRequestValidationError{
						field:  "end_device_version_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "message":

			if v, ok := interface{}(&m.Message).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return EncodeDownlinkMessageRequestValidationError{
						field:  "message",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "formatter":
			// no validation rules for Formatter
		case "parameter":
			// no validation rules for Parameter
		default:
			return EncodeDownlinkMessageRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// EncodeDownlinkMessageRequestValidationError is the validation error returned
// by EncodeDownlinkMessageRequest.ValidateFields if the designated
// constraints aren't met.
type EncodeDownlinkMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EncodeDownlinkMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EncodeDownlinkMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EncodeDownlinkMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EncodeDownlinkMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EncodeDownlinkMessageRequestValidationError) ErrorName() string {
	return "EncodeDownlinkMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e EncodeDownlinkMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEncodeDownlinkMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EncodeDownlinkMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EncodeDownlinkMessageRequestValidationError{}

// ValidateFields checks the field values on DecodeUplinkMessageRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *DecodeUplinkMessageRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = DecodeUplinkMessageRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.EndDeviceIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return DecodeUplinkMessageRequestValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "end_device_version_ids":

			if v, ok := interface{}(&m.EndDeviceVersionIDs).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return DecodeUplinkMessageRequestValidationError{
						field:  "end_device_version_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "message":

			if v, ok := interface{}(&m.Message).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return DecodeUplinkMessageRequestValidationError{
						field:  "message",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "formatter":
			// no validation rules for Formatter
		case "parameter":
			// no validation rules for Parameter
		default:
			return DecodeUplinkMessageRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// DecodeUplinkMessageRequestValidationError is the validation error returned
// by DecodeUplinkMessageRequest.ValidateFields if the designated constraints
// aren't met.
type DecodeUplinkMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecodeUplinkMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecodeUplinkMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecodeUplinkMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecodeUplinkMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecodeUplinkMessageRequestValidationError) ErrorName() string {
	return "DecodeUplinkMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DecodeUplinkMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecodeUplinkMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecodeUplinkMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecodeUplinkMessageRequestValidationError{}

// ValidateFields checks the field values on DecodeDownlinkMessageRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *DecodeDownlinkMessageRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = DecodeDownlinkMessageRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.EndDeviceIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return DecodeDownlinkMessageRequestValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "end_device_version_ids":

			if v, ok := interface{}(&m.EndDeviceVersionIDs).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return DecodeDownlinkMessageRequestValidationError{
						field:  "end_device_version_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "message":

			if v, ok := interface{}(&m.Message).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return DecodeDownlinkMessageRequestValidationError{
						field:  "message",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "formatter":
			// no validation rules for Formatter
		case "parameter":
			// no validation rules for Parameter
		default:
			return DecodeDownlinkMessageRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// DecodeDownlinkMessageRequestValidationError is the validation error returned
// by DecodeDownlinkMessageRequest.ValidateFields if the designated
// constraints aren't met.
type DecodeDownlinkMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecodeDownlinkMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecodeDownlinkMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecodeDownlinkMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecodeDownlinkMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecodeDownlinkMessageRequestValidationError) ErrorName() string {
	return "DecodeDownlinkMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DecodeDownlinkMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecodeDownlinkMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecodeDownlinkMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecodeDownlinkMessageRequestValidationError{}
