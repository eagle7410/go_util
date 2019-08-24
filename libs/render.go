package lib

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

func SlowRenderContent(name *string, templ string, data interface{}) (buffer bytes.Buffer, err error){
	frame := template.New(*name)

	tpl, err := frame.Parse(templ);

	if err != nil {
		return  buffer, errors.New(fmt.Sprintf("Error on parse template %v : %s", *name, err))
	}

	if err := tpl.Execute(&buffer, data); err != nil {
		return  buffer, errors.New(fmt.Sprintf("Error on execute template %v : %s", *name, err))
	}

	return buffer, nil
}
