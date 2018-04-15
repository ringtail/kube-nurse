package log

import (
	"github.com/ringtail/kube-nurse/types"
	"fmt"
)

type KeywordsScalpel struct {
	types.Scalpel
	ScalpelName    string
	ReferenceTable types.ReferenceTable
}

func (ks *KeywordsScalpel) Name() string {
	return ks.ScalpelName
}

func (ks *KeywordsScalpel) Fit(symptom *types.Symptom) bool {
	return true
}

func (ks *KeywordsScalpel) Cut(symptom *types.Symptom) error {
	fmt.Println("scalpel has handled symptom: " + symptom.Type)
	return nil
}
