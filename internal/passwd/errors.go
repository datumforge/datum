package passwd

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// Error constants
var (
	ErrCannotCreateDK        = errors.New("cannot create derived key for empty password")
	ErrCouldNotGenerate      = fmt.Errorf("could not generate %d length", dkSLen)
	ErrUnableToVerify        = errors.New("cannot verify empty derived key or password")
	ErrCannotParseDK         = errors.New("cannot parse encoded derived key, does not match regular expression")
	ErrCannotParseEncodedEK  = errors.New("cannot parse encoded derived key, matched expression does not contain enough subgroups")
	ErrDKProtocol            = fmt.Errorf("current code only works with the the dk protcol %q", dkAlg)
	ErrExpectedDKVersion     = fmt.Errorf("expected %s version %d", dkAlg, argon2.Version)
	ErrDKCouldNotParseMemory = fmt.Errorf("could not parse memory")
	ErrDKParseTime           = fmt.Errorf("could not parse time")
	ErrDKParseThreads        = fmt.Errorf("could not parse threads")
	ErrDKParseSalt           = fmt.Errorf("could not parse salt")
	ErrDKParseDK             = fmt.Errorf("could not parse derived key")
)
