🛠 BUILD
---
{{ $status := .BuildStatus }}
{{- if eq $status "success" -}} 🟢 SUCCESS {{ else }} 🔴 FAILURE {{- end }} {{ $status := .BuildStatus}}
Status: {{- if eq $status "success" -}} 🟢 Success {{ else }} 🔴 Build Failed {{- end }}
Project: {{ .ProjectName }}
Initiator: {{ .User.Name }}