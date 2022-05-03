package main

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

type cluster struct {
	Server                   string           `yaml:"server"`
	CertificateAuthorityData string           `yaml:"certificate-authority-data"`
	Extensions               []extensionEntry `yaml:"extensions"`
	CertificateAuthority     string           `yaml:"certificate-authority"`
}

type context struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

type user struct {
	ClientCertificateData string `yaml:"client-certificate-data"`
	ClientKeyData         string `yaml:"client-key-data"`
}

type extension struct {
	Provider   string `yaml:"provider"`
	Version    string `yaml:"version"`
	LastUpdate string `yaml:"last-update"`
}

type extensionEntry struct {
	Name      string    `yaml:"name"`
	Extension extension `yaml:"extension"`
}

type clusterEntry struct {
	Cluster cluster `yaml:"cluster"`
	Name    string  `yaml:"name"`
}

type contextEntry struct {
	Name    string  `yaml:"name"`
	Context context `yaml:"context"`
}

type userEntry struct {
	Name string `yaml:"name"`
	User user   `yaml:"user"`
}

type preferences struct{}

type kubeconfig struct {
	APIVersion     string         `yaml:"apiVersion"`
	King           string         `yaml:"king"`
	Clusters       []clusterEntry `yaml:"clusters"`
	Contexts       []contextEntry `yaml:"contexts"`
	CurrentContext string         `yaml:"current-context"`
	Users          []userEntry    `yaml:"users"`
	Preferences    preferences    `yaml:"preferences"`
}

func prettyPrint(structure interface{}) {
	v := reflect.ValueOf(structure)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Tag, v.Field(i).Interface())
	}
}

func main() {
	raw, err := os.ReadFile("/home/andrey/.kube/vgadalov")
	if err != nil {
		fmt.Println("cannot open kubeconfig:", err)
	}
	conf := kubeconfig{}
	err = yaml.Unmarshal(raw, &conf)
	if err != nil {
		fmt.Println("cannot read kubeconfig:", err)
	}
	prettyPrint(conf)
}
