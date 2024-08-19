{{- define "user_secret_prefix" -}}
{{ .Values.MongoDBCommunity.name }}-user
{{- end -}}
