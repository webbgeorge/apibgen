package apib

import (
	"io"
	"strings"
	"text/template"
)

var (
	docTmpl *template.Template
	docFmt  = `{{with .Group}}{{.Render}}{{end}}`
)

func init() {
	docTmpl = template.Must(template.New("doc").Parse(docFmt))
}

type Doc struct {
	Group ResourceGroup
}

type Metadata struct {
	Format string
	Host   string
}

func NewDoc(resourceGroupName string, filePath string) (doc *Doc, err error) {
	doc = &Doc{
		Group: ResourceGroup{
			Title: strings.Title(resourceGroupName),
		},
	}

	return
}

func (d *Doc) AddResource(resource *Resource) {
	d.Group.Resources = append(d.Group.Resources, *resource)
}

func (d *Doc) Write(writer io.Writer) error {
	return docTmpl.Execute(writer, d)
}
