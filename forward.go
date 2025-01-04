package errors

import "errors"

// Forward official go errors package

func New(text string) error         { return errors.New(text) }
func Join(errs ...error) error      { return errors.Join(errs...) }
func Unwrap(err error) error        { return errors.Unwrap(err) }
func Is(err, target error) bool     { return errors.Is(err, target) }
func As(err error, target any) bool { return errors.As(err, target) }
