package main

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cuego"
)

func main() {
	type Metadata struct {
		Labels []string
	}
	type Config struct {
		Metadata Metadata
		Length   string
		Filename string
		OptFile  string `json:",omitempty"`
		MaxCount int
		MinCount int
	}

	err := cuego.Constrain(&Config{}, `{
		let jsonFile = =~".json$"

		// Filename must be defined and have a .json extension
		Filename: jsonFile

		Length: strings.MaxRunes(3)
		Metadata: Labels: list.MaxItems(3)


		// OptFile must be undefined or be a file name with a .json extension
		OptFile?: jsonFile

		MinCount: >0 & <=MaxCount
		MaxCount: <=10_000
	}`)

	fmt.Println("error:", errMsg(err))

	fmt.Println("validate1:", errMsg(cuego.Validate(&Config{
		Metadata: Metadata{
			Labels: []string{"test", "test2"},
		},
		Filename: "foo.json",
		MaxCount: 1200,
		MinCount: 39,
	})))

	fmt.Println("validate2:", errMsg(cuego.Validate(&Config{
		Filename: "foo.json",
		Length:   "12344",
		MaxCount: 12,
		MinCount: 39,
	})))

	fmt.Println("validate:", errMsg(cuego.Validate(&Config{
		Filename: "foo.jso",
		MaxCount: 120,
		MinCount: 39,
	})))

	// TODO(errors): fix bound message (should be "does not match")

}

func errMsg(err error) string {
	a := []string{}
	for _, err := range errors.Errors(err) {
		a = append(a, err.Error())
	}
	s := strings.Join(a, "\n")
	if s == "" {
		return "nil"
	}
	return s
}
