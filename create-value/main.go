package main

import (
	"fmt"
	"log"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	cueerrors "cuelang.org/go/cue/errors"
	"cuelang.org/go/encoding/yaml"
)

const v1alpha1 = `
#v1alpha1: {
	apiVersion: "example.com/v1alpha1"
	kind: string
	metadata: labels: {[string]: string} 
	env: [...]
}
`

func main() {

	// create a context
	c := cuecontext.New()

	// compile our schema first
	def := c.CompileString(v1alpha1)
	if err := def.Err(); err != nil {
		log.Fatalf("err %+v", cueerrors.Errors(err))
	}

	v := c.CompileString(`
	#v1alpha1 &{
		kind: "Test"
		metadata: labels: {
		"test": "test"
		}

		env: [
			"test"
		]
	}

	`,
		cue.Scope(def))

	if err := v.Err(); err != nil {
		log.Fatalf("err %+v", cueerrors.Errors(err))
	}

	err := v.Validate(
		// not final or concrete
		cue.Concrete(true),
		// check everything
		cue.Definitions(true),
		cue.Hidden(true),
		cue.Optional(true),
	)

	if err != nil {
		log.Fatalf("err %+v", cueerrors.Errors(err))
	}

	b, err := yaml.Encode(v)
	if err != nil {
		log.Fatalf("err %+v", cueerrors.Errors(err))
	}

	fmt.Printf("%v", string(b))
}
