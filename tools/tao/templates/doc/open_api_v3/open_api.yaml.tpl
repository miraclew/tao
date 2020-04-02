openapi: "3.0.0"
info:
  version: 1.0.0
  title: {{.Info.Title}} service
  license:
    name: MIT
servers:
  - url: 'http://localhost:8080'
    description: local
  - url: 'http://xxxx'
    description: dev
  - url: 'http://yyyyzzz'
    description: prod

{{ define "schema" }}
  {{.Name}}:
  {{- if ne .Ref "" }} 
    $ref: "{{.Ref}}"
  {{- else if eq .Type "array" }}
    type: {{.Type}}
    items: 
      $ref: "{{.Items.Ref}}"
  {{- else if eq .Type "object" }} 
    type: {{.Type}}
    {{ if gt (len .Properties) 0 }}properties:{{ end }}
  {{- range .Properties -}}
    {{- include "schema" . | indent 4 -}}
  {{- end -}}
  {{- else}}
    type: {{.Type}} 
  {{- end -}}
{{ end -}}

paths:
{{- range .Paths}}
  {{.Path}}:
  {{- range .Methods}}
    {{.Name}}:
      summary: {{ .Summary}}
      tags:
        {{- range .Tags}}
        - {{.}}
        {{- end }}
      requestBody:
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/{{.RequestBody.Ref}}"  
      responses:
        '200':
          description: Success
          content:
            application/json:    
              schema:
                $ref: "#/components/schemas/{{.Response.Ref}}"   
  {{ end }}
{{- end}} 

components:
  schemas:
  {{- range .Components.Schemas}}
    {{ include "schema" . | indent 4}}
  {{- end}}
