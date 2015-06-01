package apib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"text/template"
)

const (
	FORMAT  = "1A"
	outFile = "apidoc.apib"
)

var (
	apibTmpl *template.Template
	apibFmt  = `FORMAT: {{.Metadata.Format}}
HOST: {{.Metadata.Host}}

# {{.Title}}
{{.Description}}

{{range .ResourceGroups}}
{{.Render}}
{{end}}
`
)

func init() {
	apibTmpl = template.Must(template.New("apib").Parse(apibFmt))
}

type APIBlueprint struct {
	Title          string
	Description    string
	Metadata       *Metadata
	ResourceGroups []*ResourceGroup
	file           *os.File

	// TODO:
	// DataStructures
}

type Metadata struct {
	Format string
	Host   string
}

func NewAPIBlueprint(outDir string) (doc *APIBlueprint, err error) {
	var fi *os.File

	outPath := filepath.Join(outDir, outFile)
	fi, err = os.Create(outPath)
	if err != nil {
		return
	}

	doc = tmpDoc
	doc.file = fi

	err = apibTmpl.Execute(fi, doc)
	if err != nil {
		return
	}

	return
}

func (doc *APIBlueprint) RecordRequest(req *http.Request) error {
	body, err := getPayload(req)
	if err != nil {
		return err
	}

	fmt.Println(body)

	// err = doc.WriteRequestTitle("")
	// if err != nil {
	// 	return err
	// }

	// err = doc.WriteHeaders(req.Header)
	// if err != nil {
	// 	return err
	// }

	// return doc.WriteBody(string(body))

	return nil
}

func (doc *APIBlueprint) RecordResponse(handler http.Handler, req *http.Request) (resp *httptest.ResponseRecorder, err error) {
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	// err = doc.WriteResponseTitle(resp.Code, resp.Header().Get("Content-Type"))
	// if err != nil {
	// 	return
	// }

	// err = doc.WriteHeaders(resp.Header())
	// if err != nil {
	// 	return
	// }

	// err = doc.WriteBody(string(resp.Body.String()))
	return
}

func getPayload(req *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}
