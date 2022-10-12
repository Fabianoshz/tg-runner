package entity

type Action int64

const (
	Apply Action = iota
	Destroy
	Plan
	PlanDestroy
)
