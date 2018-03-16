package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/dns"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/pkg/util/version"
)

var (
	components = []string{
		"etcd",
		"kube-apiserver",
		"kube-controller-manager",
		"kube-scheduler",
	}

	addonDNS = []string{
		"kube-dns",
		"coredns",
	}
)

func main() {
	v := flag.String("k8s-version", "", "kubernetes version")
	flag.Parse()

	kubernetesImageTag := kubeadmutil.KubernetesVersionToImageTag(*v)
	etcdImageTag := constants.DefaultEtcdVersion

	result := map[string]string{
		constants.Etcd:                  etcdImageTag,
		constants.KubeAPIServer:         kubernetesImageTag,
		constants.KubeControllerManager: kubernetesImageTag,
		constants.KubeScheduler:         kubernetesImageTag,
	}

	k8sVersion, err := version.ParseSemantic(*v)
	if err != nil {
		fmt.Println(err)
		return
	}

	result["kube-dns"] = dns.GetKubeDNSVersion(k8sVersion)
	result["kube-proxy"] = kubernetesImageTag

	bs, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
