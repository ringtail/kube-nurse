package log

type KubeDnsScalpel struct {
	KeywordsScalpel
}

func (kns *KubeDnsScalpel) Name() string {
	return "kubedns"
}

func NewKubeDnsScalpel() *KubeDnsScalpel {
	return &KubeDnsScalpel{}
}
