package envparse

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// ErrInvalidSpecification indicates that a specification is of the wrong type.
var ErrInvalidSpecification = errors.New("specification must be a struct pointer")

// varInfo maintains information about the configuration variable
type varInfo struct {
	Name string
	Key  string
	Tags reflect.StructTag
}

// GatherEnvInfo gathers information about the specified struct, including defaults and environment variable names.
func GatherEnvInfo(prefix string, spec interface{}) ([]varInfo, error) {
	s := reflect.ValueOf(spec)

	// Ensure the specification is a pointer to a struct
	if s.Kind() != reflect.Ptr {
		return nil, ErrInvalidSpecification
	}

	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, ErrInvalidSpecification
	}

	typeOfSpec := s.Type()

	// Create a slice to hold the information about the configuration variables
	var infos []varInfo

	// Iterate over the struct fields
	for i := range s.NumField() {
		f := s.Field(i)
		ftype := typeOfSpec.Field(i)

		if !f.CanSet() {
			continue
		}

		for f.Kind() == reflect.Ptr {
			if f.IsNil() {
				if f.Type().Elem().Kind() != reflect.Struct {
					// nil pointer to a non-struct: leave it alone
					break
				}

				// nil pointer to struct: create a zero instance
				f.Set(reflect.New(f.Type().Elem()))
			}

			f = f.Elem()
		}

		// Capture information about the config variable
		info := varInfo{
			Name: ftype.Name,
			Tags: ftype.Tag,
		}

		// Default to the field name as the env var name (will be upcased)
		info.Key = info.Name

		if prefix != "" {
			info.Key = fmt.Sprintf("%s_%s", prefix, info.Key)
		}

		info.Key = strings.ToUpper(info.Key)
		infos = append(infos, info)

		if f.Kind() == reflect.Struct {
			innerPrefix := prefix

			if !ftype.Anonymous {
				innerPrefix = info.Key
			}

			embeddedPtr := f.Addr().Interface()

			// Recursively gather information about the embedded struct
			embeddedInfos, err := GatherEnvInfo(innerPrefix, embeddedPtr)
			if err != nil {
				return nil, err
			}

			infos = append(infos[:len(infos)-1], embeddedInfos...)

			continue
		}
	}
	return infos, nil
}
