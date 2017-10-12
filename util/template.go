package util

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"golang.org/x/net/context"

	"github.com/edgedagency/mulungu/logger"
)

//TemplateParse parses content data
func TemplateParse(ctx context.Context, source string, data interface{}) string {
	templateIdentifier := strings.Join([]string{"template", GenerateRandomCode(5, "")}, "")
	t, err := template.New(templateIdentifier).Parse(source)

	logger.Debugf(ctx, "template parse util", "parsing template, templateIdentifier: %s data: %#v", templateIdentifier, data)

	if err != nil {
		logger.Errorf(ctx, "template parse util", "parse template failed with error %s", err.Error())
		return fmt.Sprintf("failed to parse template %s error %s", source, err.Error())
	}

	var output bytes.Buffer

	if err := t.Execute(&output, data); err != nil {
		logger.Errorf(ctx, "template parse util", "execute template failed with error %s", err.Error())
		return fmt.Sprintf("failed to parse template, execution failed %s error %s", source, err.Error())
	}

	return output.String()
}
