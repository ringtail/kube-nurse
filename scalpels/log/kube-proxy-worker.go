package log

type KubeProxyWorkerScalpel struct {
	KeywordsScalpel
}

func (kpws *KubeProxyWorkerScalpel) Name() string {
	return "kube-proxy-worker"
}

func NewKubeProxyWorkerScalpel() *KubeProxyWorkerScalpel {
	return &KubeProxyWorkerScalpel{}
}
