package openapi

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

type typeWalker struct {
	Pkg   *packages.Package
	Enums map[types.Type][]interface{}
}

func NewTypeWalker(pkgname string) (*typeWalker, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo}, pkgname)
	if err != nil {
		return nil, err
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("found %d packages, expected 1", len(pkgs))
	}

	tw := &typeWalker{
		Pkg:   pkgs[0],
		Enums: map[types.Type][]interface{}{},
	}

	// populate enums: walk all types declared at package scope
	for _, n := range pkgs[0].Types.Scope().Names() {
		o := pkgs[0].Types.Scope().Lookup(n)
		if !o.Exported() {
			continue
		}

		if o, ok := o.(*types.Const); ok {
			switch o.Val().Kind() {
			case constant.Int:
				i, _ := constant.Int64Val(o.Val())
				tw.Enums[o.Type()] = append(tw.Enums[o.Type()], i)
			case constant.String:
				tw.Enums[o.Type()] = append(tw.Enums[o.Type()], constant.StringVal(o.Val()))
			default:
				panic(o.Val())
			}
		}
	}

	return tw, nil
}

// GetNodes returns the hierarchy of ast Nodes for a given token.Pos
func (tw *typeWalker) GetNodes(pos token.Pos) (path []ast.Node, exact bool) {
	// find the matching file
	for _, f := range tw.Pkg.Syntax {
		if tw.Pkg.Fset.File(f.Pos()) == tw.Pkg.Fset.File(pos) {
			// find the matching token in the file
			return astutil.PathEnclosingInterval(f, pos, pos)
		}
	}

	return
}

// SchemaFromType returns a Schema object populated from the given Type.  If
// references are made to dependent types, these are added to deps
func (tw *typeWalker) SchemaFromType(t types.Type, deps map[*types.Named]struct{}) (s *Schema) {
	s = &Schema{}

	switch t := t.(type) {
	case *types.Basic:
		switch t.Kind() {
		case types.Bool:
			s.Type = "boolean"
		case types.Int, types.Int64:
			s.Type = "integer"
		case types.String:
			s.Type = "string"
		default:
			panic(t)
		}

	case *types.Map:
		s.Type = "object"
		s.AdditionalProperties = tw.SchemaFromType(t.Elem(), deps)

	case *types.Named:
		s.Ref = "#/definitions/" + t.Obj().Name()
		deps[t] = struct{}{}

	case *types.Pointer:
		s = tw.SchemaFromType(t.Elem(), deps)

	case *types.Slice:
		if e, ok := t.Elem().(*types.Basic); ok {
			// handle []byte as a string (it'll be base64 encoded by json.Marshal)
			if e.Kind() == types.Uint8 {
				s.Type = "string"
			}
		} else {
			s.Type = "array"
			s.Items = tw.SchemaFromType(t.Elem(), deps)
		}

	case *types.Struct:
		for i := 0; i < t.NumFields(); i++ {
			field := t.Field(i)

			nodes, _ := tw.GetNodes(field.Pos())
			if len(nodes) < 2 {
				continue
			}

			node, ok := nodes[1].(*ast.Field)
			// TOOD: look at this
			if !ok {
				continue
			}

			// We can skip this
			if node.Tag == nil {
				continue
			}

			tag, _ := strconv.Unquote(node.Tag.Value)

			name := strings.SplitN(reflect.StructTag(tag).Get("json"), ",", 2)[0]
			if name == "-" {
				continue
			}

			properties := tw.SchemaFromType(field.Type(), deps)
			properties.Description = strings.Trim(node.Doc.Text(), "\n")

			s.Properties = append(s.Properties, NameSchema{
				Name:   name,
				Schema: properties,
			})
		}

	default:
		return
	}

	return
}

// _define adds a Definition for the given type and recurses on any dependencies
func (tw *typeWalker) _define(definitions Definitions, t *types.Named) {
	deps := map[*types.Named]struct{}{}

	s := tw.SchemaFromType(t.Underlying(), deps)

	path, _ := tw.GetNodes(t.Obj().Pos())
	if len(path) < 2 {
		return
	}

	s.Description = strings.Trim(path[len(path)-2].(*ast.GenDecl).Doc.Text(), "\n")
	s.Enum = tw.Enums[t]

	definitions[t.Obj().Name()] = s

	for dep := range deps {
		if _, found := definitions[dep.Obj().Name()]; !found {
			tw._define(definitions, dep)
		}
	}
}

// Define adds a Definition for the named type
func (tw *typeWalker) Define(definitions Definitions, name string) {
	o := tw.Pkg.Types.Scope().Lookup(name)

	t, ok := o.(*types.TypeName)
	if !ok {
		return
	}
	tw._define(definitions, t.Type().(*types.Named))
}

// Define adds a Definition for the named types in the given package
func Define(definitions Definitions, pkgname string, names ...string) error {
	th, err := NewTypeWalker(pkgname)
	if err != nil {
		return err
	}

	for _, name := range names {
		th.Define(definitions, name)
	}

	return nil
}
