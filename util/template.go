package util

import (
	"bytes"
	"fmt"
	"html/template"
)

//ProcessTemplate takes a template name to create a new template and processes it with provided data
func ProcessTemplate(name, source string, data map[string]interface{}) string {
	buffer := new(bytes.Buffer)

	t := template.Must(template.New(name).Parse(source))
	err := t.Execute(buffer, data)
	if err != nil {
		fmt.Println("Unable to process template")
	}

	return buffer.String()
}
