{{- /*gotype: e.coding.net/miraclew/tao/tools/tao/mapper/ir.ProtoIR*/ -}}
import 'package:douyin/data/client.dart';

{{- range .Enums }}
enum {{.Name}} { {{range .Values}} {{.Name}}, {{end }} }
{{- end }}

class {{.Name}}Service {
  static final {{.Name}}Service _singleton = new {{.Name}}Service._internal();
  static const app = "{{.App}}";

  factory {{.Name}}Service() {
    return _singleton;
  }
  {{.Name}}Service._internal();
{{range .Service.Methods}}
  Future<{{.Response}}> {{.Name}}({{.Request}} req) async {
    var data = await ApiClient().rpc(app, "/v1/{{$.Name|lower}}/{{.Name|lower}}", req);
    return {{.Response}}.fromJson(data);
  }
{{end -}}
}
{{range .Messages}}
class {{.Name}} {
{{- range .Fields}}
  {{.Type.String}} {{.Name}};
{{- end}}

  {{.Name}}({ {{range .Fields}}this.{{.Name}}, {{end}} });

  static {{.Name}} fromJson(Map<String, dynamic> data) {
    return {{.Name}}(
      {{- range .Fields}}
      {{- if .Type.Repeated}}
        {{- if .Type.Scalar}}
        {{.Name}}: {{.Type.String}}.from(data['{{.Name|title}}']),
        {{else}}
        {{.Name}}: data['{{.Name|title}}'] != null ? List.of(data['{{.Name|title}}']).map((e) => {{.Type.Name}}.fromJson(e)).toList() : [],
        {{end}}
      {{- else if .Type.Scalar }}
      {{.Name}}: data['{{.Name|title}}'],
      {{- else if .Type.Map }}
      {{.Name}}: data['{{.Name|title}}'],
      {{- else if .Type.Enum }}
      {{.Name}}: {{.Type.Name}}.values[data['{{.Name|title}}']],
      {{- else}}
        {{if eq .Type.Name "DateTime"}}
      {{.Name}}: DateTime.parse(data['{{.Name|title}}']),
        {{- else }}
      {{.Name}}: {{.Type.Name}}.fromJson(data['{{.Name|title}}']),
        {{- end}}
      {{- end}}
      {{- end}}
    );
  }

  Map toJson() {
    return Map.from({
      {{- range .Fields}}
      {{- if .Type.Enum}}
      "{{.Name|title}}": this.{{.Name}}.index,
      {{- else}}
      "{{.Name|title}}": this.{{.Name}},
      {{- end}}
      {{- end}}
    });
  }
}
{{end}}