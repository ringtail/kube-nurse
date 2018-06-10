package store

import (
	"github.com/ringtail/kube-nurse/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	log "github.com/Sirupsen/logrus"
	"strings"
	"encoding/json"
	"github.com/ringtail/kube-nurse/summary"
)

const (
	Nodes                  = "nodes"
	Services               = "services"
	ReplicaSets            = "replicasets"
	Deployments            = "deployments"
	Daemonsets             = "daemonsets"
	Replicationcontrollers = "replicationcontrollers"
	Events                 = "events"
	Pods                   = "pods"
)

type ListItemMeta struct {
	v1.ListMeta
	Items []interface{}
}

func NewJsonScapel() *JsonScapel {
	return &JsonScapel{
		ScalpelName: "Json",
	}
}

type JsonScapel struct {
	types.Scalpel
	ScalpelName string
	Type        string
}

func (js *JsonScapel) Name() string {
	return js.ScalpelName
}

func (js *JsonScapel) Fit(symptom *types.Symptom) bool {
	return false
}

func (js *JsonScapel) Cut(symptom *types.Symptom) error {
	jsonBytes := SymptomContentToBytes(symptom.Content)
	list := &ListItemMeta{}
	err := json.Unmarshal(jsonBytes, list)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	err = HandleList(list)
	return err
}

func HandleList(list *ListItemMeta) error {
	arr := strings.Split(list.GetSelfLink(), "/")
	switch arr[len(arr)-1 ] {
	case Nodes:
		nodeSummary, err := HandleNodesList(list.Items)
		if err != nil {
			log.Errorf("Failed to diagnose nodes,because of %s\n", err.Error())
		} else {
			summary.AddNodeSummary(nodeSummary)
		}

	case ReplicaSets:
		replicasetSummary, err := HandleReplicasetList(list.Items)
		if err != nil {
			log.Errorf("Failed to diagnose replicasets,because of %s\n", err.Error())
		} else {
			//log.Infof("normal %d,warning %d,error %d", len(replicasetSummary.NormalItems), len(replicasetSummary.WarningItems), len(replicasetSummary.ErrorItems))
			summary.AddApplicationSummary(ReplicaSets, GetNamesapcesFromSelflink(list.GetSelfLink()), replicasetSummary)
		}
	}
	return nil
}

func SymptomContentToBytes(content []string) []byte {
	jsonString := ""
	for _, line := range content {
		jsonString += strings.Trim(line, "\n")
	}
	return []byte(jsonString)
}

func GetNamesapcesFromSelflink(selfLink string) string {
	arr := strings.Split(selfLink, "/")
	for index, item := range arr {
		if item == "namespaces" {
			return arr[index+1]
		}
	}
	return ""
}
