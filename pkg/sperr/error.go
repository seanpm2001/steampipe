package sperr

import (
	"fmt"
	"io"
	"strings"
)

type Error struct {
	stack         *stack
	cause         error
	detail        string
	message       string
	isRootMessage bool
}

// RootCause will retrieves the underlying root error in the error stack
// RootCause will recursively retrieve
// the topmost error that does not have a cause, which is assumed to be
// the original cause.
func (e Error) RootCause() error {
	type hasCause interface {
		Cause() error
	}
	if e.cause == nil {
		// return self if we don't have a cause
		// I was created with New
		return e
	}
	if cause, ok := e.cause.(hasCause); ok {
		return cause.Cause()
	}
	return e.cause
}

// Cause returns the underlying cause of this error. Maybe <nil> if this was created with New
func (e Error) Cause() error {
	return e.cause
}

// Stack retrieves the stack trace of the absolute underlying sperr.Error
func (e Error) Stack() StackTrace {
	type hasStack interface {
		Stack() StackTrace
	}
	if cause, ok := e.cause.(hasStack); ok {
		return cause.Stack()
	}
	if e.stack == nil {
		panic("sperr: stack cannot be nil")
	}
	return e.stack.StackTrace()
}

// Unwrap returns the immediately underlying error
func (e Error) Unwrap() error { return e.cause }

func (e Error) Error() (str string) {
	res := []string{}
	if len(e.message) > 0 {
		res = append(res, e.message)
	}
	if e.isRootMessage || e.cause == nil {
		return e.message
	}
	if e.cause != nil && len(e.cause.Error()) > 0 {
		res = append(res, e.cause.Error())
	}
	return strings.Join(res, " : ")
}

func (e Error) Detail() string {
	type hasDetail interface {
		Detail() string
	}
	res := []string{}
	if len(e.detail) > 0 {
		// if this is available - the underlying error will always be a sperr
		res = append(res, fmt.Sprintf("%s :: %s", e.message, e.detail))
	}
	if e.cause != nil && len(e.cause.Error()) > 0 {
		if asD, ok := e.cause.(hasDetail); ok {
			res = append(res, asD.Detail())
		} else {
			if len(e.Error()) > 0 {
				res = append(res, e.Error())
			}
		}
	}
	return strings.Join(res, "\n|-- ")
}

// All error values returned from this package implement fmt.Formatter and can
// be formatted by the fmt package. The following verbs are supported:
//
//			%s    print the error. If the error has a Cause it will be
//			      printed recursively.
//			%v    see %s
//			%+v   detailed format - includes messages and detail.
//	    %#v   Each Frame of the error's StackTrace will be printed in detail.
//		  %q		a double-quoted string safely escaped with Go syntax
//
// TODO: add Details for +
func (e Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		io.WriteString(s, e.Error())
		io.WriteString(s, "\n")

		printStack := s.Flag('#')
		printDetail := printStack || s.Flag('+')

		if printDetail {
			io.WriteString(s, "\nDetails:\n")
			io.WriteString(s, e.Detail())
			io.WriteString(s, "\n")
		}

		if printStack {
			io.WriteString(s, "\nStack:")
			e.Stack().Format(s, verb)
			io.WriteString(s, "\n")
		}
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

// WithMessage wraps Error and sets the provided message in the new error
func (e *Error) WithMessage(format string, args ...interface{}) *Error {
	if e == nil {
		return nil
	}
	res := e
	// if there's a message, wrap this error and set the message on the new error
	if len(e.message) > 0 {
		res = &Error{
			cause: e,
		}
	}
	res.message = fmt.Sprintf(format, args...)
	return res
}

// WithDetail wraps Error and sets the provided detail message in the new error
func (e *Error) WithDetail(format string, args ...interface{}) *Error {
	if e == nil {
		return nil
	}
	res := e
	// if there's a detail, wrap this error and set the detail the new error
	if len(e.detail) > 0 {
		res = &Error{
			cause: e,
		}
	}
	res.detail = fmt.Sprintf(format, args...)
	return res
}

// AsRootMessage sets this error as the root error in the error stack.
// When an Error is set as root, all child errors are hidden from display
func (e *Error) AsRootMessage() *Error {
	if e == nil {
		return nil
	}
	e.isRootMessage = true
	return e
}
