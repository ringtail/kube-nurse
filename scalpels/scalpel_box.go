package scalpels

import (
	"github.com/ringtail/kube-nurse/types"
	"github.com/ringtail/kube-nurse/scalpels/log"
	"github.com/ringtail/kube-nurse/scalpels/store"
)

var OneScapelBox *ScalpelBox

func init() {
	OneScapelBox = &ScalpelBox{}
	OneScapelBox.scalpels = make(map[string]types.Scalpel, 0)
	OneScapelBox.AddScalpel(log.NewKubeApiServerScalpel())
	OneScapelBox.AddScalpel(log.NewCloudControllerManagerScalpel())
	OneScapelBox.AddScalpel(log.NewHeapsterScalpel())
	OneScapelBox.AddScalpel(log.NewKubeControllerManagerScalpel())
	OneScapelBox.AddScalpel(log.NewKubeDnsScalpel())
	OneScapelBox.AddScalpel(log.NewKubeFlannelScalpel())
	OneScapelBox.AddScalpel(log.NewKubeProxyMasterScalpel())
	OneScapelBox.AddScalpel(log.NewKubeProxyWorkerScalpel())
	OneScapelBox.AddScalpel(log.NewKubeSchedulerScalpel())
	OneScapelBox.AddScalpel(store.NewJsonScapel())
}

type ScalpelBox struct {
	scalpels map[string]types.Scalpel
}

func (sb *ScalpelBox) AddScalpel(sl types.Scalpel) {
	if sb.scalpels[sl.Name()] != nil {
		return
	}
	sb.scalpels[sl.Name()] = sl
}

func (sb *ScalpelBox) FindFitScalpel(symptom *types.Symptom) types.Scalpel {
	for _, scalpel := range sb.scalpels {
		if scalpel.Fit(symptom) == true {
			return scalpel
		}
	}
	return nil
}

func (sb *ScalpelBox) FindScalpelByName(name string) types.Scalpel {
	return sb.scalpels[name]
}
