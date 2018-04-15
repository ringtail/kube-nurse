package log

type KubeFlannelScalpel struct {
	KeywordsScalpel
}

func (ks *KubeFlannelScalpel) Name() string {
	return "kube-flannel"
}

func NewKubeFlannelScalpel() *KubeFlannelScalpel {
	return &KubeFlannelScalpel{}
}
