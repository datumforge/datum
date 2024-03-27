package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	oas3 "github.com/getkin/kin-openapi/openapi3"
	"sigs.k8s.io/yaml"

	"github.com/datumforge/datum/pkg/rout"
)

type Spec = oas3.T

var rxYamlExtension = regexp.MustCompile(`(?i)\.ya?ml\s*$`)

func ReadURL(oas3url string) (*Spec, error) {
	resp, err := http.Get(oas3url) // #nosec G107
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return Parse(bytes)
}

// ReadFile does optional validation which is useful when
// merging incomplete spec files.
func ReadFile(oas3file string, validate bool) (*Spec, error) {
	if validate {
		return readAndValidateFile(oas3file)
	}
	bytes, err := os.ReadFile(oas3file)
	if err != nil {
		return nil, rout.Wrapf(err, "ReadFile.ReadFile.Error.Filename file: (%v)", oas3file)
	}
	if rxYamlExtension.MatchString(oas3file) {
		bytes, err = yaml.YAMLToJSON(bytes)
		if err != nil {
			return nil, err
		}
	}
	spec := &Spec{}
	err = spec.UnmarshalJSON(bytes)
	if err != nil {
		return nil, rout.Wrapf(err, "error ReadFile.UnmarshalJSON.Error.Filename file: (%s) ", oas3file)
	}
	return spec, nil
}

func readAndValidateFile(oas3file string) (*Spec, error) {
	bytes, err := os.ReadFile(oas3file)
	if err != nil {
		return nil, rout.Wrap(err, "E_READ_FILE_ERROR")
	}
	spec, err := oas3.NewLoader().LoadFromData(bytes)
	if err != nil {
		return spec, rout.Wrapf(err, "error `oas3.NewLoader().LoadFromData(bytes)` file: (%s)", oas3file)
	}
	_, err = ValidateMore(spec)
	return spec, err
}

// Parse will parse a byte array to an `*oas3.Swagger` struct.
// It will use JSON first. If unsuccessful, it will attempt to
// parse it as YAML.
func Parse(oas3Bytes []byte) (*Spec, error) {
	spec := &Spec{}
	err := spec.UnmarshalJSON(oas3Bytes)
	if err != nil {
		bytes, err2 := yaml.YAMLToJSON(oas3Bytes)
		if err2 != nil {
			return spec, err
		}
		spec = &Spec{}
		err3 := spec.UnmarshalJSON(bytes)
		return spec, err3
	}
	return spec, err
}

type ValidationStatus struct {
	Status  bool
	Message string
	Context string
	OpenAPI string
}

/*
status: false
message: |-
  expected Object {
    title: 'Medium API',
    description: 'Articles that matter on social publishing platform'
  } to have key version
  	missing keys: version
context: "#/info"
openapi: 3.0.0
*/

func ValidateMore(spec *Spec) (ValidationStatus, error) {
	vs := ValidationStatus{OpenAPI: "3.0.0"}
	version := strings.TrimSpace(spec.Info.Version)
	if len(version) == 0 {
		jdata, err := MarshalSimple(spec.Info, "", "  ")
		if err != nil {
			return vs, err
		}
		vs := ValidationStatus{
			Context: "#/info",
			Message: fmt.Sprintf("expect Object %s to have key version\nmissing keys:version", string(jdata)),
			OpenAPI: "3.0.0"}
		return vs, fmt.Errorf("E_OPENAPI3_MISSING_KEY [%s]", "info/version")
	}
	vs.Status = true
	return vs, nil
}

func MarshalSimple(v any, prefix, indent string) ([]byte, error) {
	if prefix == "" && indent == "" {
		return json.Marshal(v)
	} else {
		return json.MarshalIndent(v, prefix, indent)
	}
}

const FileExt = ".json"

var (
	MarshalPrefix = ""
	MarshalIndent = "    "
)
