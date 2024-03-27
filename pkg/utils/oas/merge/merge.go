package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/mogo/os/osutil"
	iv "github.com/invopop/yaml"
	"gopkg.in/yaml.v2"
	meow "sigs.k8s.io/yaml"

	"github.com/grokify/gocharts/v2/data/table/tabulator"

	"github.com/datumforge/datum/pkg/rout"
)

func specFromFile(path string) *openapi3.T {
	schema, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		panic(err)
	}

	return schema
}

var yamlPath = "pkg/utils/oas/merge/base.yaml"

func main() {
	spec, err := ReadFile("pkg/utils/oas/merge/openapi.json", false)
	if err != nil {
		panic(err)
	}

	restSpec, err := ReadFile("pkg/utils/oas/merge/doc.json", false)
	if err != nil {
		panic(err)
	}

	opts := MergeOptions{
		ValidateEach:  false,
		ValidateFinal: false,
	}

	meowzer, _ := Merge(spec, restSpec, "MITB", &opts)

	yamlz, _ := iv.Marshal(meowzer)

	bufferJSON, err := json.MarshalIndent(meowzer, "", "  ")
	if err != nil {
		panic(err)
	}

	tmp := make([]byte, len(bufferJSON), len(bufferJSON)+1)
	copy(tmp, bufferJSON)
	tmp = append(tmp, '\n')
	bufferJSON = tmp

	err = os.WriteFile("./outputschema.gen.json", bufferJSON, 0o644) // nolint: gomnd,gosec
	if err != nil {
		panic(err)
	}

	// moreyaml, _ := yaml.Marshal(meowzer)
	//	hotdog, _, := meow.JSONToYAML(bufferJSON)

	err = os.WriteFile("./mitb.gen.yaml", yamlz, 0o644) // nolint: gomnd,gosec
	if err != nil {
		panic(err)
	}
}

func notMain(yamlPath string) {
	// store schema part as buffer
	schemasYaml, err := yaml.Marshal(yamlPath)
	if err != nil {
		panic(err)
	}

	bufferYAML, err := os.ReadFile("./schemagen/base.yaml")
	if err != nil {
		panic(err)
	}

	// append both into single schema
	bufferYAML = append(bufferYAML, schemasYaml...)

	// load full schema
	loadedSchema, err := openapi3.NewLoader().LoadFromData(bufferYAML)
	if err != nil {
		panic(err)
	}

	// update version in the full schema and store it again
	if len(os.Args) >= 2 { // nolint: gomnd
		loadedSchema.Info.Version = os.Args[1]
		bufferYAML, err = yaml.Marshal(&loadedSchema)

		if err != nil {
			panic(err)
		}
	}

	// validate it
	err = loadedSchema.Validate(context.Background())
	if err != nil {
		panic(err)
	}

	// and store the full schema as JSON and YAML
	bufferJSON, err := json.MarshalIndent(loadedSchema, "", "  ")

	if err != nil {
		panic(err)
	}

	tmp := make([]byte, len(bufferJSON), len(bufferJSON)+1)
	copy(tmp, bufferJSON)
	tmp = append(tmp, '\n')
	bufferJSON = tmp

	err = os.WriteFile("./oas/outputschema.gen.json", bufferJSON, 0o644) // nolint: gomnd,gosec
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./oas/outputschema.gen.yaml", bufferYAML, 0o644) // nolint: gomnd,gosec
	if err != nil {
		panic(err)
	}
}

type Spec = openapi3.T

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
		bytes, err = meow.YAMLToJSON(bytes)
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
	spec, err := openapi3.NewLoader().LoadFromData(bytes)
	if err != nil {
		return spec, rout.Wrapf(err, "error `openapi3.NewLoader().LoadFromData(bytes)` file: (%s)", oas3file)
	}
	_, err = ValidateMore(spec)
	return spec, err
}

// Parse will parse a byte array to an `*openapi3.Swagger` struct.
// It will use JSON first. If unsuccessful, it will attempt to
// parse it as YAML.
func Parse(oas3Bytes []byte) (*Spec, error) {
	spec := &Spec{}
	err := spec.UnmarshalJSON(oas3Bytes)
	if err != nil {
		bytes, err2 := meow.YAMLToJSON(oas3Bytes)
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

type CollisionCheckResult int

const (
	CollisionCheckSame CollisionCheckResult = iota
	CollisionCheckOverwrite
	CollisionCheckError
	CollisionCheckSkip
)

type MergeOptions struct {
	FileRx               *regexp.Regexp
	SchemaFunc           func(schemaName string, sch1, sch2 interface{}, hint2 string) CollisionCheckResult
	CollisionCheckResult CollisionCheckResult
	ValidateEach         bool
	ValidateFinal        bool
	TableColumns         *tabulator.ColumnSet
	TableOpFilterFunc    func(path, method string, op *openapi3.Operation) bool
}

func NewMergeOptionsSkip() *MergeOptions {
	return &MergeOptions{
		CollisionCheckResult: CollisionCheckSkip,
		SchemaFunc:           SchemaCheckCollisionSkip}
}

func (mo *MergeOptions) CheckSchemaCollision(schemaName string, sch1, sch2 interface{}, hint2 string) CollisionCheckResult {
	if mo.CollisionCheckResult == CollisionCheckSkip {
		mo.SchemaFunc = SchemaCheckCollisionSkip
	} else if mo.SchemaFunc == nil {
		mo.SchemaFunc = SchemaCheckCollisionDefault
	}
	return mo.SchemaFunc(schemaName, sch1, sch2, hint2)
}

func SchemaCheckCollisionDefault(schemaName string, item1, item2 interface{}, item2Note string) CollisionCheckResult {
	if reflect.DeepEqual(item1, item2) {
		return CollisionCheckSame
	}
	return CollisionCheckError
}

func SchemaCheckCollisionSkip(schemaName string, item1, item2 interface{}, item2Note string) CollisionCheckResult {
	if reflect.DeepEqual(item1, item2) {
		return CollisionCheckSame
	}
	return CollisionCheckSkip
}

var jsonFileRx = regexp.MustCompile(`(?i)\.(json|yaml|yml)\s*$`)

func MergeDirectory(dir string, mergeOpts *MergeOptions) (*Spec, int, error) {
	var filenames []string
	var err error
	if mergeOpts != nil && mergeOpts.FileRx != nil {
		entries, err := osutil.ReadDirMore(dir, mergeOpts.FileRx, false, true, false)
		if err != nil {
			filenames = entries.Names(dir)
		}
	} else {
		entries, err := osutil.ReadDirMore(dir, jsonFileRx, false, true, false)
		if err != nil {
			filenames = entries.Names(dir)
		}
	}

	if err != nil {
		return nil, len(filenames), err
	}

	spec, err := MergeFiles(filenames, mergeOpts)
	return spec, len(filenames), err
}

func MergeFiles(filepaths []string, mergeOpts *MergeOptions) (*Spec, error) {
	sort.Strings(filepaths)
	validateEach := false
	validateFinal := true
	if mergeOpts != nil {
		validateEach = mergeOpts.ValidateEach
		validateFinal = mergeOpts.ValidateFinal
	}
	var specMaster *Spec
	for i, fpath := range filepaths {
		thisSpec, err := ReadFile(fpath, validateEach)
		if err != nil {
			return specMaster, rout.Wrap(err, fmt.Sprintf("ReadSpecError [%v] ValidateEach [%v]", fpath, validateEach))
		}
		if i == 0 {
			specMaster = thisSpec
		} else {
			specMaster, err = Merge(specMaster, thisSpec, fpath, mergeOpts)
			if err != nil {
				return nil, rout.Wrap(err, fmt.Sprintf("Merging [%v]", fpath))
			}
		}
	}

	if validateFinal {
		bytes, err := specMaster.MarshalJSON()
		if err != nil {
			return specMaster, err
		}
		newSpec, err := openapi3.NewLoader().LoadFromData(bytes)
		if err != nil {
			return newSpec, rout.Wrap(err, "Loader.LoadSwaggerFromData (MergeFiles().ValidateFinal)")
		}
		return newSpec, nil
	}
	return specMaster, nil
}

func Merge(specMaster, specExtra *Spec, specExtraNote string, mergeOpts *MergeOptions) (*Spec, error) {
	specMaster = MergeTags(specMaster, specExtra)
	specMaster, err := MergeParameters(specMaster, specExtra, specExtraNote, mergeOpts)
	if err != nil {
		return specMaster, err
	}
	specMaster, err = MergeSchemas(specMaster, specExtra, specExtraNote, mergeOpts)
	if err != nil {
		return specMaster, err
	}
	specMaster, err = MergePaths(specMaster, specExtra)
	if err != nil {
		return specMaster, err
	}
	specMaster, err = MergeResponses(specMaster, specExtra, specExtraNote, mergeOpts)
	if err != nil {
		return specMaster, err
	}
	return MergeRequestBodies(specMaster, specExtra, specExtraNote)
}

func MergeTags(specMaster, specExtra *Spec) *Spec {
	tagsMap := map[string]int{}
	for _, tag := range specMaster.Tags {
		tagsMap[tag.Name] = 1
	}
	for _, tag := range specExtra.Tags {
		tag.Name = strings.TrimSpace(tag.Name)
		if _, ok := tagsMap[tag.Name]; !ok {
			specMaster.Tags = append(specMaster.Tags, tag)
		}
	}
	return specMaster
}

type SpecMore struct {
	Spec *Spec
}

func MergePaths(specMaster, specExtra *Spec) (*Spec, error) {
	// getkin v0.121.0 to v0.122.0 - new version
	if specExtra == nil {
		return specMaster, errors.New("spec extra cannot be nil")
	}
	addPathMap := specExtra.Paths.Map()
	for addPathKey, addPathItem := range addPathMap {
		if addPathItem == nil {
			continue
		}
		srcPathItem := specMaster.Paths.Find(addPathKey)
		if srcPathItem == nil {
			specMaster.Paths.Set(addPathKey, addPathItem)
			continue
		}
		srcPathItemMore := PathItemMore{PathItem: srcPathItem}
		err := srcPathItemMore.AddPathItemOperations(addPathItem, false)
		if err != nil {
			return specMaster, err
		}
		specMaster.Paths.Set(addPathKey, srcPathItemMore.PathItem)
	}

	return specMaster, nil
}

type PathItemMore struct {
	PathItem *openapi3.PathItem
}

func (pm *PathItemMore) AddPathItemOperations(add *openapi3.PathItem, overwriteOpration bool) error {
	if add == nil {
		return nil
	} else if pm.PathItem == nil {
		return errors.New("path item is not set")
	}
	methods := httputilmore.Methods()
	for _, method := range methods {
		opAdd := add.GetOperation(method)
		if opAdd == nil {
			continue
		} else if overwriteOpration {
			pm.PathItem.SetOperation(method, opAdd)
		} else {
			opSrc := pm.PathItem.GetOperation(method)
			if opSrc == nil {
				pm.PathItem.SetOperation(method, opAdd)
			} else if !reflect.DeepEqual(opAdd, opSrc) {
				return fmt.Errorf("operation collision on op id (%s)", opSrc.OperationID)
			}
		}
	}
	return nil
}

type HTTPMethod string

const (
	MethodConnect HTTPMethod = http.MethodConnect
	MethodDelete             = http.MethodDelete
	MethodGet                = http.MethodGet
	MethodHead               = http.MethodHead
	MethodOptions            = http.MethodOptions
	MethodPatch              = http.MethodPatch
	MethodPost               = http.MethodPost
	MethodPut                = http.MethodPut
	MethodTrace              = http.MethodTrace
)

// ParseHTTPMethod returns a HTTPMethod type for a string.
func ParseHTTPMethod(method string) (HTTPMethod, error) {
	methodCanonical := strings.ToUpper(strings.TrimSpace(method))
	switch methodCanonical {
	case http.MethodConnect:
		return MethodConnect, nil
	case http.MethodDelete:
		return MethodDelete, nil
	case http.MethodGet:
		return MethodGet, nil
	case http.MethodHead:
		return MethodHead, nil
	case http.MethodOptions:
		return MethodOptions, nil
	case http.MethodPatch:
		return MethodPatch, nil
	case http.MethodPost:
		return MethodPost, nil
	case http.MethodPut:
		return MethodPut, nil
	case http.MethodTrace:
		return MethodTrace, nil
	}
	return MethodConnect, fmt.Errorf("cannot parse method [%v]", method)
}

// ParseHTTPMethodString returns a HTTPMethod as a string for a string.
func ParseHTTPMethodString(method string) (string, error) {
	methodCanonical, err := ParseHTTPMethod(method)
	return string(methodCanonical), err
}

func MethodsMap() map[string]int {
	return map[string]int{
		http.MethodConnect: 1,
		http.MethodDelete:  1,
		http.MethodGet:     1,
		http.MethodHead:    1,
		http.MethodOptions: 1,
		http.MethodPatch:   1,
		http.MethodPost:    1,
		http.MethodPut:     1,
		http.MethodTrace:   1,
	}
}

func Methods() []string {
	return []string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
		http.MethodTrace}
}

func MergeParameters(specMaster, specExtra *Spec, specExtraNote string, mergeOpts *MergeOptions) (*Spec, error) {
	if specMaster.Components.Parameters == nil {
		specMaster.Components.Parameters = map[string]*openapi3.ParameterRef{}
	}
	for pName, pExtra := range specExtra.Components.Parameters {
		if pExtra == nil {
			continue
		} else if pMaster, ok := specMaster.Components.Parameters[pName]; ok {
			if pMaster == nil {
				specMaster.Components.Parameters[pName] = pExtra
			} else {
				if mergeOpts == nil {
					mergeOpts = &MergeOptions{}
				}
				if mergeOpts.CollisionCheckResult == CollisionCheckSkip {
					continue
				} else if reflect.DeepEqual(pExtra, pMaster) {
					continue
				} else if mergeOpts.CollisionCheckResult == CollisionCheckOverwrite {
					specExtra.Components.Parameters[pName] = pExtra
				} else {
					return nil, fmt.Errorf("E_SCHEMA_COLLISION [%v] EXTRA_COMPONENTS_PARAMETER [%s]", specExtraNote, pName)
				}
			}
		} else {
			specMaster.Components.Parameters[pName] = pExtra
		}
	}
	return specMaster, nil
}

func MergeResponses(specMaster, specExtra *Spec, specExtraNote string, mergeOpts *MergeOptions) (*Spec, error) {
	if specMaster.Components.Responses == nil {
		specMaster.Components.Responses = map[string]*openapi3.ResponseRef{}
	}
	for rName, rExtra := range specExtra.Components.Responses {
		if rExtra == nil {
			continue
		} else if rMaster, ok := specMaster.Components.Responses[rName]; ok {
			if rMaster == nil {
				specMaster.Components.Responses[rName] = rExtra
			} else {
				if mergeOpts == nil {
					mergeOpts = &MergeOptions{}
				}
				if mergeOpts.CollisionCheckResult == CollisionCheckSkip {
					continue
				} else if reflect.DeepEqual(rExtra, rMaster) {
					continue
				} else {
					return nil, fmt.Errorf("E_SCHEMA_COLLISION [%v] EXTRA_COMPONENTS_RESPONSE [%s]", specExtraNote, rName)
				}
			}
		} else {
			specMaster.Components.Responses[rName] = rExtra
		}
	}
	return specMaster, nil
}

func MergeSchemas(specMaster, specExtra *Spec, specExtraNote string, mergeOpts *MergeOptions) (*Spec, error) {
	for schemaName, schemaExtra := range specExtra.Components.Schemas {
		if schemaExtra == nil {
			continue
		} else if schemaMaster, ok := specMaster.Components.Schemas[schemaName]; ok {
			if schemaMaster == nil {
				specMaster.Components.Schemas[schemaName] = schemaExtra
			} else {
				if mergeOpts == nil {
					mergeOpts = &MergeOptions{}
				}
				checkCollisionResult := mergeOpts.CheckSchemaCollision(schemaName, schemaMaster, schemaExtra, specExtraNote)
				if checkCollisionResult != CollisionCheckSame &&
					mergeOpts.CollisionCheckResult != CollisionCheckSkip {
					if mergeOpts.CollisionCheckResult == CollisionCheckOverwrite {
						delete(specMaster.Components.Schemas, schemaName)
						specMaster.Components.Schemas[schemaName] = schemaExtra
					} else if mergeOpts.CollisionCheckResult == CollisionCheckError {
						return nil, fmt.Errorf("E_SCHEMA_COLLISION [%v] EXTRA_SPEC [%s]", schemaName, specExtraNote)
					}
				}

				continue
			}
		} else {
			specMaster.Components.Schemas[schemaName] = schemaExtra
		}
	}
	return specMaster, nil
}

func MergeRequestBodies(specMaster, specExtra *Spec, specExtraNote string) (*Spec, error) {
	for rbName, rbExtra := range specExtra.Components.RequestBodies {
		if rbExtra == nil {
			continue
		} else if rbMaster, ok := specMaster.Components.RequestBodies[rbName]; ok {
			if rbMaster == nil {
				if specMaster.Components.RequestBodies == nil {
					specMaster.Components.RequestBodies = map[string]*openapi3.RequestBodyRef{}
				}
				specMaster.Components.RequestBodies[rbName] = rbExtra
			} else if !reflect.DeepEqual(rbMaster, rbExtra) {
				return nil, fmt.Errorf("E_SCHEMA_COLLISION [%v] EXTRA_SPEC [%s]", rbName, specExtraNote)
			}
		} else {
			if specMaster.Components.RequestBodies == nil {
				specMaster.Components.RequestBodies = map[string]*openapi3.RequestBodyRef{}
			}
			specMaster.Components.RequestBodies[rbName] = rbExtra
		}
	}
	return specMaster, nil
}

func WriteFileDirMerge(outfile, inputDir string, perm os.FileMode, mergeOpts *MergeOptions) (int, error) {
	spec, num, err := MergeDirectory(inputDir, mergeOpts)
	if err != nil {
		return num, rout.Wrap(err, "E_OPENAPI3_MERGE_DIRECTORY_FAILED")
	}

	bytes, err := spec.MarshalJSON()
	if err != nil {
		return num, rout.Wrap(err, "E_SWAGGER2_JSON_ENCODING_FAILED")
	}

	err = os.WriteFile(outfile, bytes, perm)
	if err != nil {
		return num, rout.Wrap(err, "E_SWAGGER2_WRITE_FAILED")
	}
	return num, nil
}
