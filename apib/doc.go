package apib

import (
	"os"
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
	file  *os.File
}

type Metadata struct {
	Format string
	Host   string
}

func NewDoc(resourceGroupName string, filePath string) (doc *Doc, err error) {
	fi, err := os.Create(filePath)
	if err != nil {
		return doc, err
	}

	doc = &Doc{
		Group: ResourceGroup{
			Title: strings.Title(resourceGroupName),
		},
		file: fi,
	}

	return
}

func (d *Doc) AddResource(resource *Resource) {
	d.Group.Resources = append(d.Group.Resources, *resource)
}

func (d *Doc) Write() error {
	return docTmpl.Execute(d.file, d)
}
