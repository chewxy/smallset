package smallset

import (
	"io"
	"strings"
	"text/template"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewSetWriter())
	if err != nil {
		panic(err)
	}
}

// TypePlusPlus extends typewriter.Type as the templates do require extra information
type TypePlusPlus struct {
	typewriter.Type
	EqFn string
}

type SetWriter struct{}

func NewSetWriter() *SetWriter { return &SetWriter{} }

func (s *SetWriter) Name() string                                               { return "smallset" }
func (s *SetWriter) Imports(t typewriter.Type) (retVal []typewriter.ImportSpec) { return retVal }

func (s *SetWriter) Write(w io.Writer, t typewriter.Type) (err error) {
	tag, ok := t.FindTag(s)
	if !ok {
		return nil
	}

	if _, err = w.Write([]byte(licence)); err != nil {
		return
	}

	var tmpl *template.Template
	if tmpl, err = templates.ByTag(t, tag); err != nil {
		return err
	}

	//execute basic stuff first
	if err = tmpl.Execute(w, t); err != nil {
		return err
	}

	// work on Contains method

	// defaults
	tmpTV := typewriter.Tag{Name: "defaultEq"}
	if tmpl, err = templates.ByTag(t, tmpTV); err != nil {
		return err
	}

	// hacky stuff to allow for implementation of custom equality comparators
	var data interface{}
	data = t // default

	for _, v := range tag.Values {
		if strings.HasPrefix(v.Name, "eq_") {
			i := strings.Index(v.Name, "_")
			if i > -1 && i+1 < len(v.Name) {
				eqFn := v.Name[i+1:]
				data = TypePlusPlus{t, eqFn}

				tmpTV = typewriter.Tag{Name: "customEq"}
				if tmpl, err = templates.ByTag(t, tmpTV); err != nil {
					return err
				}

				break
			}
		}
	}

	if err = tmpl.Execute(w, data); err != nil {
		return err
	}

	return nil
}
