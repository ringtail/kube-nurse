package log

type KubeProxyMasterScalpel struct {
	KeywordsScalpel
}

func (kpms *KubeProxyMasterScalpel) Name() string {
	return "kube-proxy-master"
}

func NewKubeProxyMasterScalpel() *KubeProxyMasterScalpel {
	return &KubeProxyMasterScalpel{}
}
