package entx

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/vektah/gqlparser/v2/ast"
)

//nolint:goerr113
var (
	addJSONScalar = func(g *gen.Graph, s *ast.Schema) error {
		s.Types["JSON"] = &ast.Definition{
			Kind:        ast.Scalar,
			Description: "A valid JSON string.",
			Name:        "JSON",
		}
		return nil
	}
)

// import string mutations from entc
var (
	_ entc.Extension = (*Extension)(nil)
)
