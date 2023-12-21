package handlers

import (
	"text/template"

	"git.sula.io/solevis/poultracker/web/templates"
)

// load templates
var Template = template.Must(template.ParseFS(templates.Files(), "*.go.html"))
