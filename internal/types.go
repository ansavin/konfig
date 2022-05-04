package internal

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

// DefaultKubeconfig is path where kubectl config is stored by default
const DefaultKubeconfig = "/home/andrey/.kube/config"

// Context represents k8s context section of kubectl config file
type Context struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

// User represents k8s user section of kubectl config file
type User struct {
	ClientCertificateData string `yaml:"client-certificate-data"`
	ClientKeyData         string `yaml:"client-key-data"`
}

// Extension represents k8s extension section of kubectl config file
type Extension struct {
	Provider   string `yaml:"provider"`
	Version    string `yaml:"version"`
	LastUpdate string `yaml:"last-update"`
}

// ExtensionEntry represents list of extensions in kubectl config file
type ExtensionEntry struct {
	Name      string    `yaml:"name"`
	Extension Extension `yaml:"extension"`
}

// Cluster represents k8s cluster section of kubectl config file
type Cluster struct {
	Server                   string           `yaml:"server"`
	CertificateAuthorityData string           `yaml:"certificate-authority-data"`
	Extensions               []ExtensionEntry `yaml:"extensions"`
	CertificateAuthority     string           `yaml:"certificate-authority"`
}

// ClusterEntry represents list of clusters in kubectl config file
type ClusterEntry struct {
	Cluster Cluster `yaml:"cluster"`
	Name    string  `yaml:"name"`
}

// ContextEntry represents list of contexts in kubectl config file
type ContextEntry struct {
	Name    string  `yaml:"name"`
	Context Context `yaml:"context"`
}

// UserEntry represents list of users in kubectl config file
type UserEntry struct {
	Name string `yaml:"name"`
	User User   `yaml:"user"`
}

// Preferences represents k8s preferences section of kubectl config file
type Preferences struct{}

// Kubeconfig represents kubectl config file
type Kubeconfig struct {
	APIVersion     string         `yaml:"apiVersion"`
	Kind           string         `yaml:"kind"`
	Clusters       []ClusterEntry `yaml:"clusters"`
	Contexts       []ContextEntry `yaml:"contexts"`
	CurrentContext string         `yaml:"current-context"`
	Users          []UserEntry    `yaml:"users"`
	Preferences    Preferences    `yaml:"preferences"`
}

// PrettyPrint is debug func
func PrettyPrint(structure interface{}) {
	v := reflect.ValueOf(structure)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Tag, v.Field(i).Interface())
	}
}

// ReadConf is a helper func for reading kubeconfig files
func ReadConf(path string) (Kubeconfig, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Kubeconfig{}, fmt.Errorf("cannot open kubeconfig: %s", err)
	}

	config := Kubeconfig{}

	err = yaml.Unmarshal(raw, &config)
	if err != nil {
		return Kubeconfig{}, fmt.Errorf("cannot read kubeconfig: %s", err)
	}

	return config, nil
}

// Merge merges two kubeconfigs. If error happens, it always returns main config,
// which is assumed to be always correct, in order to continue working, because
// fails during merge kubeconfigs are assumed as normal usage of program
func Merge(MainConf, ExtraConf Kubeconfig) (Kubeconfig, error) {
	if MainConf.APIVersion != MainConf.APIVersion {
		return MainConf, fmt.Errorf("cannot merge configs with different versions")
	}

	if MainConf.Kind != ExtraConf.Kind {
		return MainConf, fmt.Errorf("cannot merge king: %s and kind: %s", MainConf.Kind, ExtraConf.Kind)
	}

	clusters := MainConf.Clusters
	clusters = append(clusters, ExtraConf.Clusters...)

	contexts := MainConf.Contexts
	contexts = append(contexts, ExtraConf.Contexts...)

	users := MainConf.Users
	users = append(users, ExtraConf.Users...)

	return Kubeconfig{
		APIVersion:     MainConf.APIVersion,
		Kind:           MainConf.Kind,
		Clusters:       clusters,
		Contexts:       contexts,
		CurrentContext: MainConf.CurrentContext,
		Users:          users,
		Preferences:    MainConf.Preferences,
	}, nil
}
