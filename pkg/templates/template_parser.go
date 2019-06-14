package templates

import (
	"bytes"
	"html/template"
)

//ITemplateParser is a encapsulation for template parsing
type ITemplateParser interface {
	ParseTemplate([]string, interface{}) (string, error)
}

// TemplateParser is a wrapper for template parser
type TemplateParser struct {
}

//ParseTemplate convert template and related data into html formated text
func (tp *TemplateParser) ParseTemplate(templateFileName []string, data interface{}) (string, error) {
	var parsedTemplate string
	t, err := template.ParseFiles(templateFileName...)
	if err != nil {
		return parsedTemplate, err
	}
	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return parsedTemplate, err
	}
	parsedTemplate = buf.String()
	return parsedTemplate, nil
}
