package scalpels

import (
	"github.com/ringtail/kube-nurse/types"
	"github.com/ringtail/kube-nurse/scalpels/log"
)

var OneScapelBox *ScalpelBox

func init() {
	OneScapelBox = &ScalpelBox{}
	OneScapelBox.scalpels = make(map[string]types.Scalpel, 0)
	OneScapelBox.AddScalpel(&log.KubeApiServerScalpel{})
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
