package types

type Scalpel interface {
	Name() string
	Fit(symptom *Symptom) bool
	Cut(symptom *Symptom) error
}
