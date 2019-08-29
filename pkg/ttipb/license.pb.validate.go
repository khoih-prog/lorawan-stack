// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttipb

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
var _license_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on License with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *License) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = LicenseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id":

			if v, ok := interface{}(&m.LicenseIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "id",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "license_issuer_ids":

			if v, ok := interface{}(&m.LicenseIssuerIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "license_issuer_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "created_at":

			if v, ok := interface{}(&m.CreatedAt).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "created_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "valid_from":

			if v, ok := interface{}(&m.ValidFrom).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "valid_from",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "valid_until":

			if v, ok := interface{}(&m.ValidUntil).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "valid_until",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "warn_for":

			if v, ok := interface{}(&m.WarnFor).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "warn_for",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "limit_for":

			if v, ok := interface{}(&m.LimitFor).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "limit_for",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "components":

		case "component_address_regexps":

		case "dev_addr_prefixes":

		case "join_eui_prefixes":

		case "multi_tenancy":
			// no validation rules for MultiTenancy
		case "max_applications":

			if v, ok := interface{}(m.GetMaxApplications()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "max_applications",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "max_clients":

			if v, ok := interface{}(m.GetMaxClients()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "max_clients",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "max_end_devices":

			if v, ok := interface{}(m.GetMaxEndDevices()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "max_end_devices",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "max_gateways":

			if v, ok := interface{}(m.GetMaxGateways()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "max_gateways",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "max_organizations":

			if v, ok := interface{}(m.GetMaxOrganizations()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "max_organizations",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "max_users":

			if v, ok := interface{}(m.GetMaxUsers()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "max_users",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "metering":

			if v, ok := interface{}(m.GetMetering()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseValidationError{
						field:  "metering",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return LicenseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// LicenseValidationError is the validation error returned by
// License.ValidateFields if the designated constraints aren't met.
type LicenseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LicenseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LicenseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LicenseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LicenseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LicenseValidationError) ErrorName() string { return "LicenseValidationError" }

// Error satisfies the builtin error interface
func (e LicenseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLicense.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LicenseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LicenseValidationError{}

// ValidateFields checks the field values on LicenseUpdate with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LicenseUpdate) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = LicenseUpdateFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "extend_valid_until":

			if v, ok := interface{}(m.GetExtendValidUntil()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return LicenseUpdateValidationError{
						field:  "extend_valid_until",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return LicenseUpdateValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// LicenseUpdateValidationError is the validation error returned by
// LicenseUpdate.ValidateFields if the designated constraints aren't met.
type LicenseUpdateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LicenseUpdateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LicenseUpdateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LicenseUpdateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LicenseUpdateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LicenseUpdateValidationError) ErrorName() string { return "LicenseUpdateValidationError" }

// Error satisfies the builtin error interface
func (e LicenseUpdateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLicenseUpdate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LicenseUpdateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LicenseUpdateValidationError{}

// ValidateFields checks the field values on MeteringConfiguration with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MeteringConfiguration) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = MeteringConfigurationFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "interval":

			if v, ok := interface{}(&m.Interval).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return MeteringConfigurationValidationError{
						field:  "interval",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "on_success":

			if v, ok := interface{}(m.GetOnSuccess()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return MeteringConfigurationValidationError{
						field:  "on_success",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "metering":
			if len(subs) == 0 {
				subs = []string{
					"aws",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "aws":

					if v, ok := interface{}(m.GetAWS()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return MeteringConfigurationValidationError{
								field:  "aws",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				}
			}
		default:
			return MeteringConfigurationValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// MeteringConfigurationValidationError is the validation error returned by
// MeteringConfiguration.ValidateFields if the designated constraints aren't met.
type MeteringConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MeteringConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MeteringConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MeteringConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MeteringConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MeteringConfigurationValidationError) ErrorName() string {
	return "MeteringConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e MeteringConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMeteringConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MeteringConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MeteringConfigurationValidationError{}

// ValidateFields checks the field values on LicenseKey with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LicenseKey) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = LicenseKeyFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "license":
			// no validation rules for License
		case "signatures":

			for idx, item := range m.GetSignatures() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return LicenseKeyValidationError{
							field:  fmt.Sprintf("signatures[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return LicenseKeyValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// LicenseKeyValidationError is the validation error returned by
// LicenseKey.ValidateFields if the designated constraints aren't met.
type LicenseKeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LicenseKeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LicenseKeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LicenseKeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LicenseKeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LicenseKeyValidationError) ErrorName() string { return "LicenseKeyValidationError" }

// Error satisfies the builtin error interface
func (e LicenseKeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLicenseKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LicenseKeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LicenseKeyValidationError{}

// ValidateFields checks the field values on MeteringData with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MeteringData) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = MeteringDataFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "tenants":

			for idx, item := range m.GetTenants() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return MeteringDataValidationError{
							field:  fmt.Sprintf("tenants[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return MeteringDataValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// MeteringDataValidationError is the validation error returned by
// MeteringData.ValidateFields if the designated constraints aren't met.
type MeteringDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MeteringDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MeteringDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MeteringDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MeteringDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MeteringDataValidationError) ErrorName() string { return "MeteringDataValidationError" }

// Error satisfies the builtin error interface
func (e MeteringDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMeteringData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MeteringDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MeteringDataValidationError{}

// ValidateFields checks the field values on MeteringConfiguration_AWS with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MeteringConfiguration_AWS) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = MeteringConfiguration_AWSFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "sku":
			// no validation rules for SKU
		default:
			return MeteringConfiguration_AWSValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// MeteringConfiguration_AWSValidationError is the validation error returned by
// MeteringConfiguration_AWS.ValidateFields if the designated constraints
// aren't met.
type MeteringConfiguration_AWSValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MeteringConfiguration_AWSValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MeteringConfiguration_AWSValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MeteringConfiguration_AWSValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MeteringConfiguration_AWSValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MeteringConfiguration_AWSValidationError) ErrorName() string {
	return "MeteringConfiguration_AWSValidationError"
}

// Error satisfies the builtin error interface
func (e MeteringConfiguration_AWSValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMeteringConfiguration_AWS.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MeteringConfiguration_AWSValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MeteringConfiguration_AWSValidationError{}

// ValidateFields checks the field values on LicenseKey_Signature with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LicenseKey_Signature) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = LicenseKey_SignatureFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "key_id":
			// no validation rules for KeyID
		case "signature":
			// no validation rules for Signature
		default:
			return LicenseKey_SignatureValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// LicenseKey_SignatureValidationError is the validation error returned by
// LicenseKey_Signature.ValidateFields if the designated constraints aren't met.
type LicenseKey_SignatureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LicenseKey_SignatureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LicenseKey_SignatureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LicenseKey_SignatureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LicenseKey_SignatureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LicenseKey_SignatureValidationError) ErrorName() string {
	return "LicenseKey_SignatureValidationError"
}

// Error satisfies the builtin error interface
func (e LicenseKey_SignatureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLicenseKey_Signature.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LicenseKey_SignatureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LicenseKey_SignatureValidationError{}

// ValidateFields checks the field values on MeteringData_TenantMeteringData
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *MeteringData_TenantMeteringData) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = MeteringData_TenantMeteringDataFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "tenant_id":

			if v, ok := interface{}(&m.TenantIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return MeteringData_TenantMeteringDataValidationError{
						field:  "tenant_id",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "totals":

			if v, ok := interface{}(m.GetTotals()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return MeteringData_TenantMeteringDataValidationError{
						field:  "totals",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return MeteringData_TenantMeteringDataValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// MeteringData_TenantMeteringDataValidationError is the validation error
// returned by MeteringData_TenantMeteringData.ValidateFields if the
// designated constraints aren't met.
type MeteringData_TenantMeteringDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MeteringData_TenantMeteringDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MeteringData_TenantMeteringDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MeteringData_TenantMeteringDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MeteringData_TenantMeteringDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MeteringData_TenantMeteringDataValidationError) ErrorName() string {
	return "MeteringData_TenantMeteringDataValidationError"
}

// Error satisfies the builtin error interface
func (e MeteringData_TenantMeteringDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMeteringData_TenantMeteringData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MeteringData_TenantMeteringDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MeteringData_TenantMeteringDataValidationError{}
