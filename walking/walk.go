package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/datumforge/datum/walking/openapi"
)

func main() {
	s := &openapi.Swagger{
		Definitions: openapi.Definitions{},
	}

	// Get All Structs
	// this current does way more than that but it's a start
	r := getInfo()
	names := []string{}

	// get all the struct names, we can filter later if we wanted
	for _, h := range r {
		for _, s := range h.Structs {
			if !exists(s.Name, names) {
				names = append(names, s.Name)
			}
		}
	}

	// Get the openapi definitions for all the structs
	err := openapi.Define(s.Definitions, "github.com/datumforge/datum/internal/httpserve/handlers", names...)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print them so we can see what we got
	for _, k := range s.Definitions {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println(k.Ref)
		fmt.Println(k.Description)
		for _, p := range k.Properties {
			fmt.Println("Name:", p.Name, "Required", p.Schema.Type, "Required:", p.Schema.Required)
		}
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	}
}

type Struct struct {
	Name string
}

type Handler struct {
	Name      string
	Structs   []Struct
	Responses map[string]string
}

func getInfo() []Handler {
	handlers := []Handler{}
	fset := token.NewFileSet()
	files, err := os.ReadDir("./internal/httpserve/handlers/")
	if err != nil {
		log.Fatal(err)
	}

	responses := map[string]string{}

	for _, file := range files {
		if !strings.Contains(file.Name(), ".go") {
			continue
		}

		if strings.Contains(file.Name(), "_test.go") {
			continue
		}

		h := Handler{}

		node, err := parser.ParseFile(fset, "./internal/httpserve/handlers/"+file.Name(), nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}

		ast.Inspect(node, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if ok {
				if strings.Contains(fn.Name.Name, "Handler") {
					h.Name = fn.Name.Name

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

		for _, node := range node.Decls {
			switch node.(type) {
			case *ast.GenDecl:
				genDecl := node.(*ast.GenDecl)
				for _, spec := range genDecl.Specs {
					switch spec.(type) {
					case *ast.TypeSpec:
						typeSpec := spec.(*ast.TypeSpec)

						switch typeSpec.Type.(type) {
						case *ast.StructType:
							h.Structs = append(h.Structs, Struct{
								Name: typeSpec.Name.Name,
							})
						}

					}
				}
			}
		}

		h.Responses = responses

		handlers = append(handlers, h)
	}

	return handlers
}

func exists(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}

	return false
}
