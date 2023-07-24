package main

import (
	"log"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	cueerrors "cuelang.org/go/cue/errors"
	"cuelang.org/go/encoding/yaml"
)

// Location contains information about where an error has occurred during cue
// validation.
type Location struct {
	File   string `json:"file,omitempty"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

// Error is a collection of fields that represent positions in files where the user
// has made some kind of error.
type Error struct {
	Message  string   `json:"message"`
	Location Location `json:"location"`
}

// Result is a collection of errors that occurred during validation.
type Result struct {
	Errors []Error `json:"errors"`
}

func main() {
	var result Result
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

	for _, e := range cueerrors.Errors(err) {
		pos := cueerrors.Positions(e)
		if len(pos) < 1 {
			continue
		}

		p := pos[len(pos)-1]

		result.Errors = append(result.Errors, Error{
			Message: e.Error(),
			Location: Location{
				File:   "gpg.yaml",
				Line:   p.Line(),
				Column: p.Column(),
			},
		})
		log.Fatalf("%+v\n", result)
	}

}
