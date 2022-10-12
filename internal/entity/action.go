package entity

type Action string

const (
	Apply       Action = "apply"
	Destroy     Action = "destroy"
	Plan        Action = "plan"
	PlanDestroy Action = "plan-destroy"
)
