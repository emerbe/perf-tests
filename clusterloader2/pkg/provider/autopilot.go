/*
Copyright 2021 The Kubernetes Authors.

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

package provider

import (
	"context"

	clientset "k8s.io/client-go/kubernetes"
	prom "k8s.io/perf-tests/clusterloader2/pkg/prometheus/clients"
	sshutil "k8s.io/perf-tests/clusterloader2/pkg/util"
)

type AutopilotProvider struct {
	features Features
}

func NewAutopilotProvider(_ map[string]string) Provider {
	return &AutopilotProvider{
		features: Features{
			SupportProbe:                        false,
			SupportImagePreload:                 false,
			SupportSnapshotPrometheusDisk:       true,
			SupportEnablePrometheusServer:       true,
			SupportGrabMetricsFromKubelets:      true,
			SupportAccessAPIServerPprofEndpoint: true,
			SupportNodeKiller:                   false,
			ShouldPrometheusScrapeApiserverOnly: true,
			ShouldScrapeKubeProxy:               false,
		},
	}
}

func (p *AutopilotProvider) Name() string {
	return AutopilotName
}

func (p *AutopilotProvider) Features() *Features {
	return &p.features
}

func (p *AutopilotProvider) GetComponentProtocolAndPort(componentName string) (string, int, error) {
	return getComponentProtocolAndPort(componentName)
}

func (p *AutopilotProvider) GetConfig() Config {
	return Config{}
}

func (p *AutopilotProvider) RunSSHCommand(cmd, host string) (string, string, int, error) {
	// gce provider takes ssh key from GCE_SSH_KEY.
	r, err := sshutil.SSH(context.Background(), cmd, host, "gce")
	return r.Stdout, r.Stderr, r.Code, err
}

func (p *AutopilotProvider) Metadata(_ clientset.Interface) (map[string]string, error) {
	return nil, nil
}

func (p *AutopilotProvider) GetManagedPrometheusClient() (prom.Client, error) {
	return prom.NewGCPManagedPrometheusClient()
}
