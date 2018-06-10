package store

import (
	"encoding/json"
	"k8s.io/api/extensions/v1beta1"
	"github.com/ringtail/kube-nurse/types"
	"github.com/ringtail/kube-nurse/summary"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"strconv"
)

func HandleReplicasetList(replicasets []interface{}) (replicasetSummary *summary.ReplicasetSummary, err error) {
	replicasetSummary = &summary.ReplicasetSummary{
		summary.ResourceSummary{
			NormalItems:  make([]summary.Item, 0),
			WarningItems: make([]summary.Item, 0),
			ErrorItems:   make([]summary.Item, 0),
		},
	}
	for _, r := range replicasets {
		rt := r.(map[string]interface{})
		replicaset := &v1beta1.ReplicaSet{}
		replicasetBytes, err := json.Marshal(r)
		if err != nil {
			log.Warnf("Failed to marshal replicaset,Because of %s", err.Error())
			continue
		}
		err = json.Unmarshal(replicasetBytes, replicaset)
		if err != nil {
			log.Warnf("Failed to unmarshal replicaset,Because of %s", err.Error())
			continue
		}

		ans := rt["annotations"].(map[string]interface{})
		desiredSize := ans["deployment.kubernetes.io/desired-replicas"]

		status := replicaset.Status
		if desiredSize == "" {
			continue
		}
		desiredSizeInt32, err := strconv.ParseInt(desiredSize.(string), 10, 32)
		if err != nil {
			continue
		}

		rn := &summary.ResourceItem{
			Name:      rt["name"].(string),
			Namespace: replicaset.GetNamespace(),
			Healthy:   true,
		}
		if status.ReadyReplicas != (int32)(desiredSizeInt32) {
			rn.Healthy = false
			rn.Conditions = []types.Condition{
				types.Condition{
					Type:    "ReplicasetNotReady",
					Reason:  "replicaSet DesiredReplicas is not equals to ReadyReplicas",
					Message: fmt.Sprintf("DesiredReplicas:%d,ReadyReplicas:%d", desiredSizeInt32, status.ReadyReplicas),
				},
			}
			replicasetSummary.ErrorItems = append(replicasetSummary.ErrorItems, rn)
		} else {
			replicasetSummary.NormalItems = append(replicasetSummary.NormalItems, rn)
		}
	}

	return
}
