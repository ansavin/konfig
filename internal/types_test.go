package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestKubeconfig(t *testing.T) {
	tests := []struct {
		name       string
		kubeconfig string
		expected   Kubeconfig
	}{
		{
			name: "rancher cluster kubeconfig example",
			kubeconfig: `
                apiVersion: v1
                kind: Config
                clusters:
                - name: "rancher-cluster"
                  cluster:
                    server: "https://example.lan/foobar"
                - name: "rancher-cluster-fqdn"
                  cluster:
                    server: "https://rancher-cluster.appdevstage.com"
                    certificate-authority-data: "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURsRENDQ\
                      W55Z0F3SUJBZ0lVSHIrMUlWNzdUSXpIdkdQeURQS3BCaVRMRHc4d0RRWUpLb1pJaHZjTkFRRUwKQ\
                      lNVV0VlbUhzOG89Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0="

                users:
                - name: "rancher-cluster"
                  user:
                    token: "kubeconfig-user-fffffff:123"

                contexts:
                - name: "rancher-cluster"
                  context:
                    user: "rancher-cluster"
                    cluster: "rancher-cluster"
                - name: "rancher-cluster-fqdn"
                  context:
                    user: "rancher-cluster"
                    cluster: "rancher-cluster-fqdn"

                current-context: "rancher-cluster-fqdn"

            `,
			expected: Kubeconfig{
				APIVersion: "v1",
				Kind:       "Config",
				Clusters: []ClusterEntry{
					{
						Cluster: Cluster{
							Server:                   "https://example.lan/foobar",
							CertificateAuthorityData: "",
							Extensions:               nil,
							CertificateAuthority:     "",
						},
						Name: "rancher-cluster",
					},
					{
						Cluster: Cluster{
							Server:                   "https://rancher-cluster.appdevstage.com",
							CertificateAuthorityData: "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURsRENDQW55Z0F3SUJBZ0lVSHIrMUlWNzdUSXpIdkdQeURQS3BCaVRMRHc4d0RRWUpLb1pJaHZjTkFRRUwKQlNVV0VlbUhzOG89Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=",
							Extensions:               nil,
							CertificateAuthority:     "",
						},
						Name: "rancher-cluster-fqdn",
					},
				},
				Contexts: []ContextEntry{
					{
						Name: "rancher-cluster",
						Context: Context{
							Cluster: "rancher-cluster",
							User:    "rancher-cluster",
						},
					},
					{
						Name: "rancher-cluster-fqdn",
						Context: Context{
							Cluster: "rancher-cluster-fqdn",
							User:    "rancher-cluster",
						},
					},
				},
				CurrentContext: "rancher-cluster-fqdn",
				Users: []UserEntry{
					{
						Name: "rancher-cluster",
						User: User{
							ClientCertificateData: "",
							ClientKeyData:         "",
							Token:                 "kubeconfig-user-fffffff:123",
						},
					},
				},
			},
		},
		{

			name: "alibaba cluster kubeconfig example",
			kubeconfig: `
                apiVersion: v1
                clusters:
                - cluster:
                    server: https://192.168.0.1:6443
                    certificate-authority-data: dkfjbnjudferg==
                  name: kubernetes
                contexts:
                - context:
                    cluster: kubernetes
                    user: "1111111111111111111"
                  name: 1111111111111111111-sdkjghbsdvljnriignflwesifojh
                current-context: 1111111111111111111-sdkjghbsdvljnriignflwesifojh
                kind: Config
                preferences: {}
                users:
                - name: "1111111111111111111"
                  user:
                    client-certificate-data: LS0tLS1CRJQ0FURS0tLS0tCg==
                    client-key-data:  LS0tLS1CRUS0tLS0tCg==
            
            `,
			expected: Kubeconfig{
				APIVersion: "v1",
				Kind:       "Config",
				Clusters: []ClusterEntry{
					{
						Cluster: Cluster{
							Server:                   "https://192.168.0.1:6443",
							CertificateAuthorityData: "dkfjbnjudferg==",
							Extensions:               nil,
							CertificateAuthority:     "",
						},
						Name: "kubernetes",
					},
				},
				Contexts: []ContextEntry{
					{
						Name: "1111111111111111111-sdkjghbsdvljnriignflwesifojh",
						Context: Context{
							Cluster: "kubernetes",
							User:    "1111111111111111111",
						},
					},
				},
				CurrentContext: "1111111111111111111-sdkjghbsdvljnriignflwesifojh",
				Users: []UserEntry{
					{
						Name: "1111111111111111111",
						User: User{
							ClientCertificateData: "LS0tLS1CRJQ0FURS0tLS0tCg==",
							ClientKeyData:         "LS0tLS1CRUS0tLS0tCg==",
							Token:                 "",
						},
					},
				},
			},
		},
		{

			name: "google kubernetes kubeconfig example",
			kubeconfig: `
                apiVersion: v1
                clusters:
                - cluster:
                    certificate-authority-data: LS0tLS0tLS0tCg==
                    server: https://10.10.10.10
                  name: google-kubernetes-engine-example
                contexts:
                - context:
                    cluster: google-kubernetes-engine-example
                    user: google-kubernetes-engine-example
                  name: google-kubernetes-engine-example
                current-context: google-kubernetes-engine-example
                kind: Config
                preferences: {}
                users:
                - name: google-kubernetes-engine-example
                  user:
                    auth-provider:
                      config:
                        cmd-args: config config-helper --format=json
                        cmd-path: /foo/bar
                        expiry-key: '{.credential.token_expiry}'
                        token-key: '{.credential.access_token}'
                      name: gcp

            `,
			expected: Kubeconfig{
				APIVersion: "v1",
				Kind:       "Config",
				Clusters: []ClusterEntry{
					{
						Cluster: Cluster{
							Server:                   "https://10.10.10.10",
							CertificateAuthorityData: "LS0tLS0tLS0tCg==",
							Extensions:               nil,
							CertificateAuthority:     "",
						},
						Name: "google-kubernetes-engine-example",
					},
				},
				Contexts: []ContextEntry{
					{
						Name: "google-kubernetes-engine-example",
						Context: Context{
							Cluster: "google-kubernetes-engine-example",
							User:    "google-kubernetes-engine-example",
						},
					},
				},
				CurrentContext: "google-kubernetes-engine-example",
				Users: []UserEntry{
					{
						Name: "google-kubernetes-engine-example",
						User: User{
							ClientCertificateData: "",
							ClientKeyData:         "",
							Token:                 "",
						},
					},
				},
			},
		},
		{

			name: "mimkube kubeconfig example",
			kubeconfig: `
                apiVersion: v1
                clusters:
                - cluster:
                    certificate-authority: /foo/bar/ca.crt
                    extensions:
                    - extension:
                        last-update: Thu, 11 Aug 2022 16:22:25 +04
                        provider: minikube.sigs.k8s.io
                        version: v1.25.2
                    name: cluster_info
                    server: https://192.168.0.1:8443
                name: minikube
                contexts:
                - context:
                    cluster: minikube
                    extensions:
                    - extension:
                        last-update: Thu, 11 Aug 2022 16:22:25 +04
                        provider: minikube.sigs.k8s.io
                        version: v1.25.2
                    name: context_info
                    namespace: default
                    user: minikube
                name: minikube
                current-context: minikube
                kind: Config
                preferences: {}
                users:
                - name: minikube
                user:
                    client-certificate: /foo/bar/client.crt
                    client-key: /foo/bar/client.key
                
            `,
			expected: Kubeconfig{
				APIVersion: "v1",
				Kind:       "Config",
				Clusters: []ClusterEntry{
					{
						Cluster: Cluster{
							Server:                   "https://192.168.0.1:8443",
							CertificateAuthorityData: "",
							Extensions: []ExtensionEntry{
								{
									Name: "",
									Extension: Extension{
										Provider:   "minikube.sigs.k8s.io",
										Version:    "v1.25.2",
										LastUpdate: "Thu, 11 Aug 2022 16:22:25 +04",
									},
								},
							},
							CertificateAuthority: "/foo/bar/ca.crt",
						},
						Name: "",
					},
				},
				Contexts: []ContextEntry{
					{
						Name: "",
						Context: Context{
							Cluster: "minikube",
							User:    "minikube",
						},
					},
				},
				CurrentContext: "minikube",
				Users: []UserEntry{
					{
						Name: "minikube",
						User: User{
							ClientCertificateData: "",
							ClientKeyData:         "",
							Token:                 "",
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.name), func(t *testing.T) {
			result := Kubeconfig{}
			err := yaml.Unmarshal([]byte(tc.kubeconfig), &result)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
