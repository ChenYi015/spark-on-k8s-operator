{{/*
Copyright 2024 The Kubeflow authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/}}

{{/*
Create the name of webhook component
*/}}
{{- define "spark-operator.webhook.name" -}}
{{- include "spark-operator.fullname" . }}-webhook
{{- end -}}

{{/*
Common labels for the webhook
*/}}
{{- define "spark-operator.webhook.labels" -}}
{{ include "spark-operator.labels" . }}
app.kubernetes.io/component: webhook
{{- end -}}

{{/*
Selector labels for the webhook
*/}}
{{- define "spark-operator.webhook.selectorLabels" -}}
{{ include "spark-operator.selectorLabels" . }}
app.kubernetes.io/component: webhook
{{- end -}}

{{/*
Create the name of service account to be used by webhook
*/}}
{{- define "spark-operator.webhook.serviceAccountName" -}}
{{- if .Values.webhook.serviceAccount.create -}}
{{ .Values.webhook.serviceAccount.name | default (include "spark-operator.webhook.name" .) }}
{{- else -}}
{{ .Values.webhook.serviceAccount.name | default "default" }}
{{- end -}}
{{- end -}}

{{/*
Create the name of the role to be used by webhook
*/}}
{{- define "spark-operator.webhook.roleName" -}}
{{- include "spark-operator.webhook.name" . }}
{{- end -}}

{{/*
Create the name of the role binding to be used by webhook
*/}}
{{- define "spark-operator.webhook.roleBindingName" -}}
{{- include "spark-operator.webhook.name" . }}
{{- end -}}

{{/*
Create the name of the secret to be used by webhook
*/}}
{{- define "spark-operator.webhook.secretName" -}}
{{ include "spark-operator.webhook.name" . }}-certs
{{- end -}}


{{/*
Create the name of the service to be used by webhook
*/}}
{{- define "spark-operator.webhook.serviceName" -}}
{{ include "spark-operator.webhook.name" . }}-svc
{{- end -}}

{{/*
Create the name of mutating webhook configuration
*/}}
{{- define "spark-operator.mutatingWebhookConfigurationName" -}}
webhook.sparkoperator.k8s.io
{{- end -}}

{{/*
Create the name of mutating webhook configuration
*/}}
{{- define "spark-operator.validatingWebhookConfigurationName" -}}
quotaenforcer.sparkoperator.k8s.io
{{- end -}}

{{/*
Create the name of the deployment to be used by webhook
*/}}
{{- define "spark-operator.webhook.deploymentName" -}}
{{ include "spark-operator.webhook.name" . }}
{{- end -}}

{{/*
Create the name of the lease resource to be used by leader election
*/}}
{{- define "spark-operator.webhook.leaderElectionName" -}}
{{ include "spark-operator.webhook.name" . }}-lock
{{- end -}}
