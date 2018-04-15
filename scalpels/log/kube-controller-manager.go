package log

type KubeControllerManagerScalpel struct {
	KeywordsScalpel
}

func (kcms *KubeControllerManagerScalpel) Name() string {
	return "kube-controller-manager"
}

func NewKubeControllerManagerScalpel() *KubeControllerManagerScalpel {
	return &KubeControllerManagerScalpel{}
}
