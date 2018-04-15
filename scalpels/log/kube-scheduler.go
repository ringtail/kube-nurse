package log

type KubeSchedulerScalpel struct {
	KeywordsScalpel
}

func (kss *KubeSchedulerScalpel) Name() string {
	return "kube-scheduler-scalpel"
}

func NewKubeSchedulerScalpel() *KubeSchedulerScalpel {
	return &KubeSchedulerScalpel{}
}
