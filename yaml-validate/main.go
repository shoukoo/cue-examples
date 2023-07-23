package main

import (
	"log"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	cueerrors "cuelang.org/go/cue/errors"
	"cuelang.org/go/encoding/yaml"
)

func main() {
	c := cuecontext.New()

	b, err := os.ReadFile("schema.cue")
	if err != nil {
		log.Fatalf("err %e", err)
	}
	schemaVal := c.CompileBytes(b, cue.Filename("schema.cue"))
	if err := schemaVal.Err(); err != nil {
		log.Fatalf("err %e", err)
	}

	b2, err := os.ReadFile("gpg.yaml")
	if err != nil {
		log.Fatalf("err %e", err)
	}
	err = yaml.Validate(b2, schemaVal)
	if err != nil {
		log.Fatalf("err %v", cueerrors.Errors(err))
	}

	// schemaVal.Validate(cue.Schema())
	//
	// f, err := yaml.Extract("gpg.yaml", nil)
	//
	//	if err != nil {
	//		log.Fatalf("err %e", err)
	//	}
	//
	// v := c.BuildFile(f, cue.Scope(schemaVal), cue.ImportPath("'gpg'"))
	// v2 := schemaVal.Unify(v)
	//
	// err = v2.Validate()
	//
	//	if err != nil {
	//		log.Fatalf("err %e", err)
	//	}
	//
	// fmt.Printf("%v \n", v)
}
