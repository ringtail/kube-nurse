package summary

import (
	"fmt"
	"strings"
	"github.com/ringtail/kube-nurse/types"
)

type Summary interface {
	Info() *ResourceSummary
}

type Item interface {
	GetName() string
	IsHealthy() bool
	GetInfo() string
}

type ResourceItem struct {
	Name       string
	Healthy    bool
	Namespace  string
	Conditions []types.Condition
}

func (s *ResourceItem) GetName() string {
	return s.Name
}

func (s *ResourceItem) IsHealthy() bool {
	return s.Healthy
}

func (s *ResourceItem) GetNamespace() string {
	return s.Namespace
}

func (s *ResourceItem) GetInfo() string {
	r := make([]string, 0)
	for i, condition := range s.Conditions {
		line := fmt.Sprintf("%d. %s because of %s,%s.", i, condition.Type, condition.Reason, condition.Message)
		r = append(r, line)
	}
	return strings.Join(r, "\n")
}

type ResourceSummary struct {
	NormalItems  []Item
	WarningItems []Item
	ErrorItems   []Item
}

func (rs *ResourceSummary) Info() *ResourceSummary {
	return rs
}

func (rs *ResourceSummary) Merge(s Summary) {
	sy := s.Info()
	rs.NormalItems = append(rs.NormalItems, sy.NormalItems...)
	rs.WarningItems = append(rs.WarningItems, sy.WarningItems...)
	rs.ErrorItems = append(rs.ErrorItems, sy.ErrorItems...)
}
