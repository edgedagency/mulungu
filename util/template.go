package util

import (
	"bytes"
	"fmt"
	"html/template"

	"golang.org/x/net/context"

	"github.com/edgedagency/mulungu/logger"
)

//TemplateParse parses content data
func TemplateParse(ctx context.Context, source string, data interface{}) string {
	t, err := template.New("messageTemplate").Parse(source)

	if err != nil {
		logger.Errorf(ctx, "messaging service", "parse template failed with error %s", err.Error())
		return fmt.Sprintf("failed to parse template %s error %s", source, err.Error())
	}

	var output bytes.Buffer

	if err := t.Execute(&output, data); err != nil {
		logger.Errorf(ctx, "messaging service", "execute template failed with error %s", err.Error())
		return fmt.Sprintf("failed to parse template, execution failed %s error %s", source, err.Error())
	}

	return output.String()
}
