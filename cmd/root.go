/*
Copyright © 2022 ansavin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"fmt"
	"reflect"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "konfig",
	Short: "CLI tool for managing kubectl config files",
	Long: `konfig - kubectl config file manager, cli tool for choosing, 
editing, backuping, viewing and merging of kubectl config files 
that are usually stored in '~/.kube/' folder
		  `,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.konfig.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
