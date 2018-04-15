package log

type HeapsterScalpel struct {
	KeywordsScalpel
}

func (hs *HeapsterScalpel) Name() string {
	return "heapster"
}

func NewHeapsterScalpel() *HeapsterScalpel {
	return &HeapsterScalpel{}
}
