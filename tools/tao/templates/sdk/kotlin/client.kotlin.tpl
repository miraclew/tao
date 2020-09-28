{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/ir.ProtoIR*/ -}}
package {{.Package}}

import retrofit2.Call
import retrofit2.http.Body
import retrofit2.http.POST

{{ range .Enums }}
enum class {{.Name}}(val value: Int) { {{range .Values}} {{.Name}}({{.Value}}), {{end }} }
{{- end }}

{{ range .Services}}
public interface {{.Name}}Service {
{{- range .Methods}}
  @POST("/v1/{{$.Name|lower}}/{{.Name|lower}}")
  fun {{.Name}}(@Body req: {{.Request}}): Call<{{.Response}}>
{{- end}}
}
{{end}}

{{range .Messages}}
{{- $n :=  len .Fields -}}
{{- if gt $n 0 -}}
data class {{.Name}}({{range $i, $elem := .Fields}}val {{.Name}}: {{.Type.String}}{{if ne (add $i 1) $n}},{{end}} {{end}})
{{- else -}}
data class {{.Name}}(val _reserved:Int=0)
{{- end}}
{{end}}
