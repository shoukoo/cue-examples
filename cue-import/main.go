package main

import (
	"fmt"
	"log"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
)

func main() {
	c := cuecontext.New()

	b, err := os.ReadFile("schema.cue")

	if err != nil {
		log.Fatalf("err %e", err)
	}

	// get schema
	schemaVal := c.CompileBytes(b, cue.Filename("schema.cue"))

	if err := schemaVal.Err(); err != nil {
		log.Fatalf("err %e", err)
	}

	// get yaml file
	f, err := yaml.Extract("gpg.yaml", nil)

	if err != nil {
		log.Fatalf("err %e", err)
	}

	v := c.BuildFile(f, cue.Scope(schemaVal), cue.ImportPath("gpg:config"))

	err = v.Validate()

	if err != nil {
		log.Fatalf("err %e", err)
	}

	fmt.Printf("%v \n", v)
}
