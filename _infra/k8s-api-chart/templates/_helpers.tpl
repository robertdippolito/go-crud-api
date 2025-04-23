{{- define "k8s-api-chart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "k8s-api-chart.fullname" -}}
{{- if .Values.fullnameOverride }}
{{-   .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{-   printf "%s-%s" (include "k8s-api-chart.name" .) .Release.Name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end -}}
