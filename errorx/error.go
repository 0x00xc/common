package errorx

import "errors"

func New(s string) error {
	return errors.New(s)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func IErrors(f ...IError) error {
	for _, v := range f {
		if err := v.Error(); err != nil {
			return err
		}
	}
	return nil
}

func ErrorFuncs(f ...func() error) error {
	for _, v := range f {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}

type IError interface {
	Error() error
}

type ErrorFunc func() error

func (e ErrorFunc) Error() error {
	return e()
}
