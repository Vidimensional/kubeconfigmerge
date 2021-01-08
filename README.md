# KubeconfigMerge

## Install

```shell
go get github.com/Vidimensional/kubeconfigmerge
```

## Run

Merges kubernetes config files to your local kubeconfig.

```
Usage:
kubeconfigmerge [flags] NEW_CONFIG...

Flags:
-h, --help                 help for main
-k, --kube-config string   Path to your local kubeconfig (default "${HOME}/.kube/config")
```