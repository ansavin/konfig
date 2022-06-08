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
	"github.com/spf13/cobra"

	"github.com/ansavin/konfig/internal"
)

// showCmd represents command to merge kubeconfigs
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restores current kubeconfig",
	Long: `copies backup file content from $HOME/.konfig/... to $HOME/.kube/config
		  `,
	Run: func(cmd *cobra.Command, args []string) {
		kubeconfig, err := internal.GetKubeconfigPath(cmd)
		if err != nil {
			panic(err)
		}

		backup, err := internal.GetBackupFilePath(cmd)
		if err != nil {
			panic(err)
		}

		err = internal.CopyFileContent(backup, kubeconfig)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	restoreCmd.Flags().String(internal.OptionBackup, "", "specify a custom backup file")
	rootCmd.AddCommand(restoreCmd)
}
