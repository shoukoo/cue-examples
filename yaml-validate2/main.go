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

	val := c.CompileBytes(b, cue.Filename("schema.cue"))
	if err := val.Err(); err != nil {
		log.Fatalf("err %+v", cueerrors.Errors(err))
	}

	schema := val.LookupPath(cue.ParsePath("#Deployment"))

	b2, err := os.ReadFile("pod.yaml")
	if err != nil {
		log.Fatalf("err %e", cueerrors.Errors(err))
	}

	err = yaml.Validate(b2, schema)
	if err != nil {
		log.Fatalf("err %+v", cueerrors.Errors(err))
	}

}
