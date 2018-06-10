package store

import (
	"k8s.io/api/core/v1"
	"encoding/json"
	"github.com/prometheus/common/log"
	"github.com/ringtail/kube-nurse/summary"
	"github.com/ringtail/kube-nurse/types"
)

func HandleNodesList(nodes []interface{}) (nodeSummary *summary.NodeSummary, err error) {
	nodeSummary = &summary.NodeSummary{
		summary.ResourceSummary{
			NormalItems:  make([]summary.Item, 0),
			WarningItems: make([]summary.Item, 0),
			ErrorItems:   make([]summary.Item, 0),
		},
	}
	for _, n := range nodes {
		node := &v1.Node{}
		nodeBytes, err := json.Marshal(n)
		if err != nil {
			log.Warnf("Failed to marshal node,Because of %s", err.Error())
			continue
		}
		err = json.Unmarshal(nodeBytes, node)
		if err != nil {
			log.Warnf("Failed to unmarshal node,Because of %s", err.Error())
			continue
		}
		status := node.Status
		var snode *summary.ResourceItem
		conditions := make([]types.Condition, 0)
		for _, conditon := range status.Conditions {
			if conditon.Status == v1.ConditionTrue && conditon.Type != v1.NodeReady {
				c := types.Condition{
					Type:    string(conditon.Type),
					Reason:  conditon.Reason,
					Message: conditon.Message,
				}
				conditions = append(conditions, c)
			}
		}
		snode = &summary.ResourceItem{
			Name: node.GetName(),
		}
		if len(conditions) > 0 {
			snode.Healthy = false
			snode.Conditions = conditions
			nodeSummary.ErrorItems = append(nodeSummary.ErrorItems, snode)
			continue
		}
		nodeSummary.NormalItems = append(nodeSummary.NormalItems, snode)
	}
	return
}
