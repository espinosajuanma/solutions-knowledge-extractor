{{- range . -}}
# #{{.Number}} {{.Title}}

| Key | Value |
|---|---|
| **Project** | {{.Project.Label}} |
| **Status** | {{.Status}} |
| **Priority** | {{.Priority}} |
| **Assignee** | {{if .Assignee}}{{.Assignee.Label}}{{else}}Unassigned{{end}} |
| **Customer** | {{.Customer.Label}} |

## Description
{{.Description | render}}

{{if .Attachments}}
### Attachments
{{range .Attachments}}- [{{.File.Name}}]({{.Link}})
{{end}}
{{end}}

{{if .Notes}}
## Public Notes
{{range .Notes}}
> **{{.Label}}**
{{ .Note | render }}
{{end}}
{{end}}

{{if .InternalNotes}}
## Internal Notes
{{ range .InternalNotes }}
> **{{.Label}}**
{{ .Note | render }}
{{ end }}
{{ end }}

---
{{ end }}