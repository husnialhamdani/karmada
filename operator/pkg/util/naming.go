/*
Copyright 2023 The Karmada Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"fmt"
	"strings"
)

// Namefunc defines a function to generate resource name according to karmada resource name.
type Namefunc func(karmada string) string

// AdminKarmadaConfigSecretName returns secret name of karmada-admin karmada-config
func AdminKarmadaConfigSecretName(karmada string) string {
	return generateResourceName(karmada, "admin-config")
}

// ComponentKarmadaConfigSecretName returns secret name of karmada component karmada-config
func ComponentKarmadaConfigSecretName(karmadaComponent string) string {
	return fmt.Sprintf("%s-config", karmadaComponent)
}

// KarmadaCertSecretName returns secret name of karmada certs
func KarmadaCertSecretName(karmada string) string {
	return generateResourceName(karmada, "cert")
}

// EtcdCertSecretName returns secret name of etcd cert
func EtcdCertSecretName(karmada string) string {
	return generateResourceName(karmada, "etcd-cert")
}

// WebhookCertSecretName returns secret name of karmada-webhook cert
func WebhookCertSecretName(karmada string) string {
	return generateResourceName(karmada, "webhook-cert")
}

// KarmadaAPIServerName returns secret name of karmada-apiserver
func KarmadaAPIServerName(karmada string) string {
	return generateResourceName(karmada, "apiserver")
}

// KarmadaAggregatedAPIServerName returns secret name of karmada-aggregated-apiserver
func KarmadaAggregatedAPIServerName(karmada string) string {
	return generateResourceName(karmada, "aggregated-apiserver")
}

// KarmadaSearchAPIServerName returns secret name of karmada-search
func KarmadaSearchAPIServerName(karmada string) string {
	return generateResourceName(karmada, "search")
}

// KarmadaEtcdName returns name of karmada-etcd
func KarmadaEtcdName(karmada string) string {
	return generateResourceName(karmada, "etcd")
}

// KarmadaEtcdClientName returns name of karmada-etcd client
func KarmadaEtcdClientName(karmada string) string {
	return generateResourceName(karmada, "etcd-client")
}

// KubeControllerManagerName returns name of kube-controller-manager
func KubeControllerManagerName(karmada string) string {
	return generateResourceName(karmada, "kube-controller-manager")
}

// KarmadaControllerManagerName returns name of karmada-controller-manager
func KarmadaControllerManagerName(karmada string) string {
	return generateResourceName(karmada, "controller-manager")
}

// KarmadaSchedulerName returns name of karmada-scheduler
func KarmadaSchedulerName(karmada string) string {
	return generateResourceName(karmada, "scheduler")
}

// KarmadaWebhookName returns name of karmada-webhook
func KarmadaWebhookName(karmada string) string {
	return generateResourceName(karmada, "webhook")
}

// KarmadaDeschedulerName returns name of karmada-descheduler
func KarmadaDeschedulerName(karmada string) string {
	return generateResourceName(karmada, "descheduler")
}

// KarmadaMetricsAdapterName returns name of karmada-metric-adapter
func KarmadaMetricsAdapterName(karmada string) string {
	return generateResourceName(karmada, "metrics-adapter")
}

// KarmadaSearchName returns name of karmada-search
func KarmadaSearchName(karmada string) string {
	return generateResourceName(karmada, "search")
}

func generateResourceName(karmada, suffix string) string {
	if strings.Contains(karmada, "karmada") {
		return fmt.Sprintf("%s-%s", karmada, suffix)
	}

	return fmt.Sprintf("%s-karmada-%s", karmada, suffix)
}
