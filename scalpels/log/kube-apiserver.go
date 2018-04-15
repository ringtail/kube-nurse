package log

type KubeApiServerScalpel struct {
	KeywordsScalpel
}

func (ks *KubeApiServerScalpel) Name() string {
	return "kube-apiserver"
}

func NewKubeApiServerScalpel() *KubeApiServerScalpel {
	return &KubeApiServerScalpel{}
}