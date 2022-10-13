🟡 <b>MERGE REQUEST</b>
<b>Initiator:</b> {{ .User.Name }}
<b>Project:</b> {{ .Project.Name }}
<b>MR Link:</b> {{ .ObjectAttributes.URL }}
<b>MR Status</b>: {{ if eq .ObjectAttributes.MergeStatus "can_be_merged" }}💚 CAN BE MERGE{{ else }}💔 CANNOT BE MERGED{{- end }}

<b>Please review and close merge request if you can do so.</b>