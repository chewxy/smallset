package smallset

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	basicTemplate,
	defaultEq,
	customEq,
}

var basicTemplate = &typewriter.Template{
	Name:           "smallset",
	Text:           basic,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}

var defaultEq = &typewriter.Template{
	Name: "defaultEq",
	Text: defaultEqTempl,
}

var customEq = &typewriter.Template{
	Name: "customEq",
	Text: customEqTempl,
}
