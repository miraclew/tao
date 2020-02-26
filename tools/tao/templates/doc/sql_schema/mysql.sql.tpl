{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/sqlschema.CreateTable*/ -}}
CREATE TABLE `{{.TableName}}` (
  {{- range .Fields }}
  `{{.Name}}` {{.Type}}{{if .Null}} null{{end}}{{if .Default}} default {{.Default}}{{end}}{{if eq $.Pk .Name}} AUTO_INCREMENT{{end}},
  {{- end }}
  PRIMARY KEY (`{{.Pk}}`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;