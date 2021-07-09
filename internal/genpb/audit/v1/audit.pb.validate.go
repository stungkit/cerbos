// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: audit/v1/audit.proto

package auditv1

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on AccessLogEntry with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AccessLogEntry) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CallId

	if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogEntryValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPeer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessLogEntryValidationError{
				field:  "Peer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetMetadata() {
		_ = val

		// no validation rules for Metadata[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessLogEntryValidationError{
					field:  fmt.Sprintf("Metadata[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Method

	// no validation rules for StatusCode

	return nil
}

// AccessLogEntryValidationError is the validation error returned by
// AccessLogEntry.Validate if the designated constraints aren't met.
type AccessLogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessLogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessLogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessLogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessLogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessLogEntryValidationError) ErrorName() string { return "AccessLogEntryValidationError" }

// Error satisfies the builtin error interface
func (e AccessLogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessLogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessLogEntryValidationError{}

// Validate checks the field values on DecisionLogEntry with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DecisionLogEntry) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CallId

	if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecisionLogEntryValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPeer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecisionLogEntryValidationError{
				field:  "Peer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetInputs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecisionLogEntryValidationError{
					field:  fmt.Sprintf("Inputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetOutputs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DecisionLogEntryValidationError{
					field:  fmt.Sprintf("Outputs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Error

	return nil
}

// DecisionLogEntryValidationError is the validation error returned by
// DecisionLogEntry.Validate if the designated constraints aren't met.
type DecisionLogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecisionLogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecisionLogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecisionLogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecisionLogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecisionLogEntryValidationError) ErrorName() string { return "DecisionLogEntryValidationError" }

// Error satisfies the builtin error interface
func (e DecisionLogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecisionLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecisionLogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecisionLogEntryValidationError{}

// Validate checks the field values on MetaValues with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *MetaValues) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MetaValuesValidationError is the validation error returned by
// MetaValues.Validate if the designated constraints aren't met.
type MetaValuesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetaValuesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetaValuesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetaValuesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetaValuesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetaValuesValidationError) ErrorName() string { return "MetaValuesValidationError" }

// Error satisfies the builtin error interface
func (e MetaValuesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetaValues.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetaValuesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetaValuesValidationError{}

// Validate checks the field values on Peer with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Peer) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Address

	// no validation rules for AuthInfo

	// no validation rules for UserAgent

	// no validation rules for ForwardedFor

	return nil
}

// PeerValidationError is the validation error returned by Peer.Validate if the
// designated constraints aren't met.
type PeerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PeerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PeerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PeerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PeerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PeerValidationError) ErrorName() string { return "PeerValidationError" }

// Error satisfies the builtin error interface
func (e PeerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPeer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PeerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PeerValidationError{}