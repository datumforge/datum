package rout

import "fmt"

// Wrap wraps an error with the supplied strings.
func Wrap(err error, wrap ...string) error {
	if err == nil {
		return nil
	}
	for i := len(wrap) - 1; i >= 0; i-- {
		err = wrapOne(err, wrap[i])
	}
	return err
}

func wrapOne(err error, wrapPrefix string) error {
	if err == nil || wrapPrefix == "" {
		return err
	}
	return fmt.Errorf("%s: [%w]", wrapPrefix, err)
}

// Wrapf will wrap the error, first performing a `fmt.Sprintf()` on the supplied params.
func Wrapf(origErr error, wrapFormat string, wrapVars ...any) error {
	if origErr == nil {
		return origErr
	}
	return wrapOne(origErr, fmt.Sprintf(wrapFormat, wrapVars...))
}
