{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/ir.ProtoIR*/ -}}

import Foundation

{{ range .Enums }}
enum {{.Name}}: Int, Codable {
{{- range .Values}}
  case {{.Name}} = {{.Value}}
{{- end}}
}
{{end }}

{{- $app := .App }}
{{- range .Services}}
{{if eq .Type 2 -}}
protocol {{.Name}}Delegate {
  {{- range .Methods}}
  {{- if hasPrefix "recv" .Name}}
  func {{.Name}}(data: {{.Request}})
  {{- end}}
  {{- end}}
}
{{- end}}
{{end}}

{{- range .Services}}
class {{.Name}}Service {
  let app = "{{$app}}"
  static let shared = {{.Name}}Service()
  private init() {}
{{if eq .Type 1 -}}

  {{- range .Methods }}
  func {{.Name}}(req: {{.Request}}, completion: @escaping ({{.Response}}?, Error?) -> ()) {
    APIClient.shared.rpc(app: app, path: "/v1/{{$.Name|lower}}/{{.Name|lower}}", req: req, completion: completion)
  }
  {{end -}}
{{- else if eq .Type 2 }}
  var delegate: {{.Name}}Delegate?

  {{range .Methods}}{{if hasPrefix "send" .Name -}}
  func {{.Name}}(req: {{.Request}}) {
    SocketClient.shared.send(data: req)
  }{{end}}
  {{- end}}
{{end}}
}{{end}}

{{range .Messages -}}
struct {{.Name}}: Codable {
{{- range .Fields}}
  var {{.Name}}: {{.Type.String}}
{{- end}}
}

{{end}}
