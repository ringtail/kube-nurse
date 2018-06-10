package summary

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"fmt"
)

var SummaryManagerInstance *SummaryManager

func init() {
	SummaryManagerInstance = &SummaryManager{}
}

func GetSummaryManager() *SummaryManager {
	return SummaryManagerInstance
}

func AddNodeSummary(s Summary) {
	SummaryManagerInstance.NodeSummary = s
}
func AddApplicationSummary(appType string, namespace string, s Summary) {
	key := namespace + "-" + appType
	if (SummaryManagerInstance.AppSummary == nil ) {
		SummaryManagerInstance.AppSummary = make(map[string]Summary)
	}
	SummaryManagerInstance.AppSummary[key] = s
}

type SummaryManager struct {
	NodeSummary    Summary
	PodSummary     Summary
	AppSummary     map[string]Summary
	ServiceSummary Summary
}

func (s *SummaryManager) DumpSummary() {
	data := [][]string{}
	nodeInfo := s.NodeSummary.Info()

	data = append(data, []string{
		"Node",
		strconv.Itoa(len(nodeInfo.NormalItems)),
		strconv.Itoa(len(nodeInfo.WarningItems)),
		strconv.Itoa(len(nodeInfo.ErrorItems)),
	})

	for name, app := range s.AppSummary {
		appInfo := app.Info()
		data = append(data, []string{
			name,
			strconv.Itoa(len(appInfo.NormalItems)),
			strconv.Itoa(len(appInfo.WarningItems)),
			strconv.Itoa(len(appInfo.ErrorItems)),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"Type", "Normal", "Warning", "Error"})

	for _, v := range data {
		table.Append(v)
	}
	fmt.Println("Cluster Summary:")
	table.Render() // Send output
}

func (s *SummaryManager) DumpDetails() {
	s.DumpSummary()
	data := [][]string{}
	for name, app := range s.AppSummary {
		appInfo := app.Info()
		if len(appInfo.ErrorItems) != 0 {
			for _, item := range appInfo.ErrorItems {
				data = append(data, []string{
					name,
					item.GetName(),
					strconv.FormatBool(item.IsHealthy()),
					item.GetInfo(),
				})
			}
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Name", "Healthy", "Message"})
	table.SetRowLine(true)
	for _, v := range data {
		table.Append(v)
	}
	fmt.Println("Cluster Exceptions:")
	table.Render() // Send output
}
