package internal

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

// PrettyPrint print current kubeconfig with colors
func PrettyPrint(k Kubeconfig) {
	magenta := color.New(color.FgMagenta)
	cyan := color.New(color.FgCyan)

	magenta.Print("apiVersion:")
	fmt.Println(k.APIVersion)

	magenta.Print("kind:")
	fmt.Println(k.Kind)

	magenta.Println("clusters:")
	flag := true
	for _, cluster := range k.Clusters {
		data, err := yaml.Marshal(cluster)
		if err != nil {
			panic(err)
		}

		switch flag {
		case true:
			fmt.Printf(`%s`, string(data))
		case false:
			cyan.Printf(`%s`, string(data))
		}

		flag = !flag
	}

	magenta.Println("contexts:")
	flag = true
	for _, context := range k.Contexts {
		data, err := yaml.Marshal(context)
		if err != nil {
			panic(err)
		}

		switch flag {
		case true:
			fmt.Printf(`%s`, string(data))
		case false:
			cyan.Printf(`%s`, string(data))
		}

		flag = !flag
	}

	magenta.Println("current-context:")
	fmt.Println(k.CurrentContext)

	magenta.Println("users:")
	flag = true
	for _, user := range k.Users {
		data, err := yaml.Marshal(user)
		if err != nil {
			panic(err)
		}

		switch flag {
		case true:
			fmt.Printf(`%s`, string(data))
		case false:
			cyan.Printf(`%s`, string(data))
		}

		flag = !flag
	}

	magenta.Print("preferences:")
	preferences, err := yaml.Marshal(k.Preferences)
	if err != nil {
		panic(err)
	}

	fmt.Printf(`%s`, string(preferences))

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

// CopyFileContent copies file content from src to dst
// If file exist, it gives it new uniq name prefix
func CopyFileContent(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		closeErr := in.Close()
		if err == nil {
			err = closeErr
		}
	}()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		closeErr := out.Close()
		if err == nil {
			err = closeErr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
