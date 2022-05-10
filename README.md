# konfig

konfig - kubectl config file manager, cli tool for backuping, viewing and merging
of kubectl config files that are usually stored in `~/.kube/` folder and also is
an example of simple but usefull app build using [cobra library](https://github.com/spf13/cobra)

If you need a more advanced tool, you might take a look at [kubecm](https://github.com/sunny0826/kubecm)

## install

To install this app, simply type
`go install github.com/ansavin/konfig`

## usage

`konfig show` - show current kubeconfig
`konfig merge /path/to/another/config` - merge current kubeconfig and another one situated at /path/to/another/config
`konfig backup` - to create a backup of current kubeconfig
`konfig restore` - to restore kubeconfig from backup
