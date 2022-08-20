package internal

// OptionOutput is cli flag name for setting custom output file
const OptionOutput = "output"

// OptionKubeconfig is cli flag name for setting custom kubeconfig file
const OptionKubeconfig = "kubeconfig"

// OptionBackup is cli flag name for setting custom backup file
const OptionBackup = "backup"

// DefaultKubeconfigFolder is path where kubectl config is stored by default
const DefaultKubeconfigFolder = ".kube"

// DefaultBackupFolder is path where backups are stored
const DefaultBackupFolder = ".konfig"

// DefaultKubeconfigFile is default backup file name
const DefaultKubeconfigFile = "config"

// Context represents k8s context section of kubectl config file
type Context struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

// User represents k8s user section of kubectl config file
type User struct {
	ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
	ClientKeyData         string `yaml:"client-key-data,omitempty"`
	Token                 string `yaml:"token,omitempty"`
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
	Preferences    Preferences    `yaml:"preferences,omitempty"`
}
