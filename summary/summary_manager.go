package summary

var SummaryManagerInstance *SummaryManager

func init() {
	SummaryManagerInstance = &SummaryManager{}
}

func GetSummaryManager() *SummaryManager {
	return SummaryManagerInstance
}


func AddItem(item Item){

}

type SummaryManager struct {
	NodeSummary    Summary
	PodSummary     Summary
	AppSummary     Summary
	ServiceSummary Summary
}

func (s *SummaryManager) DumpSummary() {

}

func (s *SummaryManager) DumpDetails() {

}

