{{- /*gotype: github.com/miraclew/tao/tools/tao/mapper/ir.ProtoIR*/ -}}

import Foundation

{{- range .Enums }}
enum {{.Name}}: Int {
{{range .Values}} case {{.Name}} = {{.Value}}, {{end }}
}
{{- end }}

class {{.Name}}Service {
  static const app = "{{.App}}";

{{range .Service.Methods}}
  func {{.Name}}({{.Request}} req, completion: @escaping ({{.Response}}) -> ()) {
    guard let url = URL(string: "/v1/{{$.Name|lower}}/{{.Name|lower}}") else {return}
    var request = URLRequest(url: url)
    request.httpMethod = "POST"
    request.setValue("application/json", forHTTPHeaderField: "Content-Type")
    request.setValue("at", forHTTPHeaderField: "Authorization")
    guard let httpBody = try? JSONSerialization.data(withJSONObject: req, options: []) else {
    return
    }
    request.httpBody = httpBody

    URLSession.shared.dataTask(with: request) { (data, _, _) in
      let res = try! JSONDecoder().decode({{.Response}}.self, from: data!)

      DispatchQueue.main.async {
        completion(posts)
      }
    }.resume()
  }
{{end -}}
}
{{range .Messages}}
struct {{.Name}}: Codable, Identifiable {
{{- range .Fields}}
  var {{.Name}}: {{.Type.String}}
{{- end}}

  {{.Name}}({ {{range .Fields}}this.{{.Name}}, {{end}} });
}
{{end}}
