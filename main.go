package main

import (
	"fmt"
	"os"

	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/dns"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
	"k8s.io/kubernetes/pkg/util/version"
)

var (
	kubernetesVesion    = os.Getenv("KUBERNETES_VERSION")
	kubernetesComponent = os.Getenv("KUBERNETES_COMPONENT")
)

func main() {
	kubernetesImageTag := kubeadmutil.KubernetesVersionToImageTag(kubernetesVesion)
	etcdImageTag := constants.DefaultEtcdVersion

	result := map[string]string{
		constants.Etcd:                  etcdImageTag,
		constants.KubeAPIServer:         kubernetesImageTag,
		constants.KubeControllerManager: kubernetesImageTag,
		constants.KubeScheduler:         kubernetesImageTag,
	}

	k8sVersion, err := version.ParseSemantic(kubernetesVesion)
	if err != nil {
		println(err)
		return
	}

	result["kube-dns"] = dns.GetKubeDNSVersion(k8sVersion)
	result["kube-proxy"] = kubernetesImageTag

	fmt.Print(result[kubernetesComponent])
}
