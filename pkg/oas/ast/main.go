package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"strings"
)

func main() {
	getReturns()
}

func getReturns() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "./internal/httpserve/handlers/login.go", nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}

	responses := map[string]string{}

	ast.Inspect(node, func(n ast.Node) bool {
		t, ok := n.(*ast.TypeSpec)
		if ok {
			fmt.Println("Type:", t.Name.Name)
		}

		s, ok := n.(*ast.StructType)
		if ok {
			buf := new(bytes.Buffer)
			printer.Fprint(buf, fset, s)
			fmt.Println(buf)
		}

		fn, ok := n.(*ast.FuncDecl)
		if ok {
			if strings.Contains(fn.Name.Name, "Handler") {
				fmt.Println("Handler found")
				fmt.Println(fn.Name.Name)

				for _, param := range fn.Type.Params.List {
					fmt.Println("Param:", param.Names[0].Name, "Type:", param.Type)
				}

				// Find Return Statements
				ast.Inspect(fn, func(n ast.Node) bool {
					ret, ok := n.(*ast.ReturnStmt)
					if ok {
						// get return string
						buf := new(bytes.Buffer)
						printer.Fprint(buf, fset, ret)

						// parse response
						if strings.Contains(buf.String(), "ctx.JSON") {
							returnValue := strings.Split(buf.String(), "(")
							returnValue = strings.Split(returnValue[1], ",")
							responses[returnValue[0]] = returnValue[1]
						}

						return true
					}
					return true
				})
			}

			return true
		}
		return true
	})

	fmt.Println("responses:", responses)
}
