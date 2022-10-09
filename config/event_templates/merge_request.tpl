🟡 MERGE REQUEST
---
Initiator: {{ .User.Name }}
Project: {{ .Project.Name }}
MR Link: {{ .ObjectAttributes.URL }}
MR Status: {{ if eq .ObjectAttributes.MergeStatus "can_be_merged" }}💚 CAN BE MERGE{{ else }}💔 CANNOT BE MERGED{{- end }}
