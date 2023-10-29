// nolint
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"entgo.io/ent/entc/gen"
)

func main() {
	var (
		entpath = flag.String("path", "ent", "path to ent package")
		skipgen = flag.Bool("skipgen", false, "skip re-generate assets")
	)
	flag.Parse()
	if _, err := os.Stat(*entpath); os.IsNotExist(err) {
		log.Fatalf("entclean: package %q was not found", *entpath)
	}
	for _, name := range flag.Args() {
		typ := &gen.Type{Name: name}
		for _, t := range gen.Templates {
			err := os.Remove(filepath.Join(*entpath, t.Format(typ)))
			if err != nil && !os.IsNotExist(err) {
				log.Fatalf("entclean: %s", err)
			}
		}
		err := os.Remove(filepath.Join(*entpath, typ.Package()))
		if err != nil && !os.IsNotExist(err) {
			log.Fatalf("entclean: %s", err)
		}
	}
	if !*skipgen {
		cmd := exec.Command("go", "generate", filepath.Join(*entpath, "generate.go"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("entclean: %s", err)
		}
	}
}
