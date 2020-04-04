{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/ir.ProtoIR*/ -}}

import Foundation

{{ range .Enums }}
enum {{.Name}}: Int, Codable {
{{- range .Values}}
  case {{.Name}} = {{.Value}}
{{- end}}
}
{{end }}

class {{.Name}}Service {
  let app = "{{.App}}"
  static let shared = {{.Name}}Service()

  private init() {}

{{range .Service.Methods}}
  func {{.Name}}(req: {{.Request}}, completion: @escaping ({{.Response}}) -> ()) {
    APIClient.shared.rpc(app: app, path: "/v1/{{$.Name|lower}}/{{.Name|lower}}", req: req, completion: completion)
  }
{{end -}}
}
{{range .Messages}}
struct {{.Name}}: Codable {
{{- range .Fields}}
  var {{.Name}}: {{.Type.String}}
{{- end}}
}
{{end}}
