{{- define "uniqueSecrets" -}}
{{- $result := dict -}}
{{- $pathToKey := dict -}}
{{- range $envVar, $path := .Values.secrets -}}
  {{- if not (hasKey $pathToKey $path) -}}
    {{- $_ := set $pathToKey $path $envVar -}}
  {{- end -}}
  {{- $key := get $pathToKey $path -}}
  {{- $_ := set $result $envVar (dict "path" $path "key" $key) -}}
{{- end -}}
{{- $result | toJson -}}
{{- end -}}