// Package vfmt formats strings containing references to variables.
//
// The syntax for variable replacement is
//  {variable name}
// or
//  {variable name|format}
//
// Variables are provided in a map[string]interface{}.
package vfmt

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"
)

var varPattern = regexp.MustCompile(`\{([^|}]+)\}`)
var fmtPattern = regexp.MustCompile(`\{([^|}]+)\|([^}]+)\}`)

// Sprintf returns the formatted string.
// Every reference to a variable is replaced by its value
// obtained from the map vars, optionally formatted.
func Sprintf(s string, vars map[string]interface{}) (string, error) {
	for _, v := range varPattern.FindAllStringSubmatch(s, -1) {
		if _, found := vars[v[1]]; !found {
			return "", fmt.Errorf("vfmt: missing variable %q\n", v[1])
		}
	}
	for _, v := range fmtPattern.FindAllStringSubmatch(s, -1) {
		if _, found := vars[v[1]]; !found {
			return "", fmt.Errorf("vfmt: missing variable %q\n", v[1])
		}
	}
	s = varPattern.ReplaceAllString(s, "{{index . `$1`}}")
	s = fmtPattern.ReplaceAllString(s, "{{index . `$1` | printf `$2`}}")
	out := new(bytes.Buffer)
	err := template.Must(template.New("vfmt").Parse(s)).Execute(out, vars)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
