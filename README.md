# KubeconfigMerge

[![Build Status](https://cloud.drone.io/api/badges/Vidimensional/kubeconfigmerge/status.svg)](https://cloud.drone.io/Vidimensional/kubeconfigmerge)

Merges kubernetes config files to your local kubeconfig.


## Install

```shell
go get github.com/Vidimensional/kubeconfigmerge
```

## Run

```
Usage:
kubeconfigmerge [flags] NEW_CONFIG...

Flags:
-h, --help                 help for main
-k, --kube-config string   Path to your local kubeconfig (default "${HOME}/.kube/config")
```