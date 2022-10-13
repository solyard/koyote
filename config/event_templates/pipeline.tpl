🛠 <b>PIPELINE</b>
{{- $var := "" }}
{{- range $buildStages := .Builds }}
{{- if ne "success" $buildStages.Status }} {{ $var = "🔴 FAILURE"}} {{- else }} {{ $var = "🟢 SUCCESS"}} {{- end }}
<b>Stages:</b>
  {{ if ne "success" $buildStages.Status }} ⛔️ {{ else }} ✅ {{ end }} {{ $buildStages.Stage | ToUpper }}
{{- end}}
<b>Status:</b> {{ $var }}
<b>Project:</b> {{ .Project.Name }}