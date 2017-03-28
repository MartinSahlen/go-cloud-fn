package template

import (
	"bytes"
	"html/template"
)

type IndexTemplateData struct {
	FunctionName string
	TargetDir    string
	TriggerHTTP  bool
}

func GenerateIndex(data IndexTemplateData) (string, error) {
	byt, err := Asset("index.js")
	if err != nil {
		return "", err
	}
	return parseTemplate(string(byt), data)
}

func parseTemplate(templateString string, data interface{}) (string, error) {
	tmpl, err := template.New("name").Parse(templateString)
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	err = tmpl.Execute(&b, data)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
