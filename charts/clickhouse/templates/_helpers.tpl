{{/*
  Define Clickhouse System Users Access Secret Name. It contains uniq naming - Release Namespace
*/}}

{{- define "clickhouse-system-users-access-secret.name" -}}
  {{- if .Values.clickhouseServer.externalSecrets.systemUsersSecret -}}
    {{- .Values.clickhouseServer.externalSecrets.systemUsersSecret -}}
  {{- else -}}
    {{- printf "clickhouse-system-users-access-%s" .Release.Namespace -}}
  {{- end -}}
{{- end -}}

{{/*
  Define Clickhouse Custom Users Access Secret Name. It contains uniq naming - Release Namespace
*/}}

{{- define "clickhouse-custom-users-access-secret.name" -}}
  {{- if .Values.clickhouseServer.externalSecrets.customUsersSecret -}}
    {{- .Values.clickhouseServer.externalSecrets.customUsersSecret -}}
  {{- else -}}
    {{- printf "clickhouse-custom-users-access-%s" .Release.Namespace -}}
  {{- end -}}
{{- end -}}

{{/*
  Define Dockerhub Secret Name. It contains uniq naming - Release Namespace
*/}}

{{- define "dockerhub-secret.name" -}}
  {{- if .Values.imagePullSecret.name -}}
    {{- .Values.imagePullSecret.name -}}
  {{- else -}}
    {{- printf "dockerhub-credentials" -}}
  {{- end -}}
{{- end -}}

{{/*
  Define Ingress Name. It contains uniq naming - Release Namespace
*/}}

{{- define "ingress.name" -}}
  {{- if .Values.ingress.name -}}
    {{- .Values.ingress.name -}}
  {{- else -}}
    {{- printf "%s-ingress" .Release.Namespace -}}
  {{- end -}}
{{- end -}}

{{/*
  Define Clickhouse Settings Config Map Name. It contains uniq naming - Release Namespace
*/}}

{{- define "clickhouse-settings-configmap.name" -}}
  {{- if .Values.clickhouseServer.settings.configMap.name -}}
    {{- .Values.clickhouseServer.settings.configMap.name -}}
  {{- else -}}
    {{- printf "%s-settings-config" .Release.Namespace -}}
  {{- end -}}
{{- end -}}

{{/*
  Define Clickhouse Backup Service Account Name. It contains uniq naming - Release Namespace
*/}}

{{- define "clickhouse-backup-service-account.name" -}}
  {{- if .Values.backup.clickhouseBackupServer.serviceAccount -}}
    {{- .Values.backup.clickhouseBackupServer.serviceAccount -}}
  {{- else -}}
    {{- printf "%s-backup-service-account" .Release.Namespace -}}
  {{- end -}}
{{- end -}}

{{/*
  Define Clickhouse Keeper Cluster Name. It contains uniq naming - Release Namespace
  Name may not be longer than 15 and it is prohibited to use _ symbol.
*/}}

{{- define "clickhouse-keeper-cluster.name" -}}
  {{- if .Values.clickhouseKeeper.clusterName -}}
    {{- .Values.clickhouseKeeper.clusterName -}}
  {{- else -}}
  {{- printf "keeper" -}}
  {{- end -}}
{{- end -}}
