package log

type CloudControllerManagerScalpel struct {
	KeywordsScalpel
}

func (ccms *CloudControllerManagerScalpel) Name() string {
	return "cloud-controller-manager"
}

func NewCloudControllerManagerScalpel() *CloudControllerManagerScalpel {
	return &CloudControllerManagerScalpel{}
}
