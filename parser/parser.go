package parser

import (
	"bytes"
	"embed"
	"fmt"

	htmlTemplate "html/template"
	textTemplate "text/template"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

//go:embed templates
var templatesFS embed.FS

func ToMarkdown(template string, data any) (string, error) {
	var output bytes.Buffer
	converter := md.NewConverter("", true, nil)

	funcMap := textTemplate.FuncMap{
		"render": func(s string) string {
			markdown, err := converter.ConvertString(s)
			if err != nil {
				return s
			}
			return markdown
		},
	}
	path := "templates/" + template + ".md"

	tmpl, err := textTemplate.New(template+".md").Funcs(funcMap).ParseFS(templatesFS, path)
	if err != nil {
		return "", fmt.Errorf("failed to parse Markdown template: %w", err)
	}
	if err := tmpl.Execute(&output, data); err != nil {
		return "", fmt.Errorf("failed to execute Markdown template: %w", err)
	}
	return output.String(), nil
}

func ToHTML(template string, data any) (string, error) {
	var output bytes.Buffer
	funcMap := htmlTemplate.FuncMap{
		"render": func(s string) htmlTemplate.HTML {
			return htmlTemplate.HTML(s)
		},
	}
	path := "templates/" + template + ".md"
	tmpl, err := htmlTemplate.New(template+".html").Funcs(funcMap).ParseFS(templatesFS, path)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML template: %w", err)
	}
	if err := tmpl.Execute(&output, data); err != nil {
		return "", fmt.Errorf("failed to execute HTML template: %w", err)
	}
	return output.String(), nil
}
