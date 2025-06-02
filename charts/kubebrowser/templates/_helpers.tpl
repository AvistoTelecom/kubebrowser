{{/*
Author: Yann Lacroix <yann.lacroix@avisto.com>
*/}}

{{/*
Create a default fully qualified app name for kubebrowser server objects
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "kubebrowser.server.fullname" -}}
{{- printf "%s-%s" (include "common.names.fullname" .) .Values.server.name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Return the proper kubebrowser server image name
*/}}
{{- define "kubebrowser.server.image" -}}
{{ include "common.images.image" (dict "imageRoot" .Values.server.image "global" .Values.global) }}
{{- end -}}

{{/*
Return the proper Docker Image Registry Secret Names for kubebrowser server
*/}}
{{- define "kubebrowser.server.imagePullSecrets" -}}
{{ include "common.images.pullSecrets" (dict "images" (list .Values.server.image) "global" .Values.global) }}
{{- end -}}

{{/*
Get the kubebrowser server configuration ConfigMap name.
*/}}
{{- define "kubebrowser.server.configmapName" -}}
{{ printf "%s-configuration" (include "kubebrowser.server.fullname" .) }}
{{- end -}}

{{/*
Get the kubebrowser server configuration ConfigMap name.
*/}}
{{- define "kubebrowser.server.kubeconfigsName" -}}
{{ printf "%s-kubeconfigs" (include "kubebrowser.server.fullname" .) }}
{{- end -}}

{{/*
Return the ingress anotation
*/}}
{{- define "kubebrowser.ingress.annotations" -}}
{{ .Values.ingress.annotations | toYaml }}
{{- end -}}

{{/*
Return the ingress hostname
*/}}
{{- define "kubebrowser.ingress.hostname" -}}
{{- tpl .Values.ingress.hostname $ -}}
{{- end -}}
