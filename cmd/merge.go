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
	"fmt"

	"github.com/spf13/cobra"

	"konfig/internal"
)

// mergeCmd represents command to merge kubeconfigs
var mergeCmd = &cobra.Command{
	Use:   "merge </path/to/config/file>",
	Short: "merge current config with one stored at </path/to/config/file>",
	Long: `Merges config from provided path with currently selected one.
	After execution currently selected config is modified and needs manual save.
		  `,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currentConfig, err := internal.ReadConf(internal.DefaultKubeconfig)
		if err != nil {
			panic(err)
		}
		extraConf, err := internal.ReadConf(args[0])
		if err != nil {
			panic(err)
		}
		currentConfig, err = internal.Merge(currentConfig, extraConf)
		if err != nil {
			// no need to exit here - just print error and no nothing
			fmt.Println(err)
			return
		}
		internal.PrettyPrint(currentConfig)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
