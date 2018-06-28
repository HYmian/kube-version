package main

import (
	"fmt"
	"os"

	"k8s.io/kubernetes/cmd/kubeadm/app/constants"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
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

	result[constants.KubeDNS] = constants.KubeDNSVersion
	result[constants.CoreDNS] = constants.CoreDNSVersion
	result[constants.KubeProxy] = kubernetesImageTag

	fmt.Print(result[kubernetesComponent])
}
