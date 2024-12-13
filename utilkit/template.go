package utilkit

import (
	"bytes"
	"io"
	"text/template"
)

func TemplateGenerateString(str string, data interface{}) (string, error) {
	tmpl, err := template.New("").Parse(str)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func TemplateGenerateFileFromReader(reader io.Reader, dstPath string, data interface{}) error {
	b, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	content, err := TemplateGenerateString(BytesToString(b), data)
	if err != nil {
		return err
	}

	err = WriteFile(dstPath, content)

	return err
}
