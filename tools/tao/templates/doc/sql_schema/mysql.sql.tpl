{{- /*gotype: e.coding.net/miraclew/tao/tools/tao/mapper/sqlschema.CreateTables*/ -}}
{{range .Items}}
CREATE TABLE `{{.TableName}}` (
  {{- range .Columns }}
  `{{.Name}}` {{.Type}}{{if .Null}} null{{else}} not null{{end}}{{if .Default.Valid}} default {{.Default.String}}{{end}}{{if eq .Name "Id"}} AUTO_INCREMENT{{end}},
  {{- end }}
  {{- $n :=  len .Keys -}}
  {{- range $i, $elem :=  .Keys}}
  {{$elem.String}}{{if ne (add $i 1) $n}},{{end}}
  {{- end}}
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;

{{end -}}