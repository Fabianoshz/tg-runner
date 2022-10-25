package entity

type Action string

const (
	Apply          Action = "apply"
	Destroy        Action = "destroy"
	Plan           Action = "plan"
	Planned        Action = "planned"
	PlanDestroy    Action = "plan-destroy"
	PlannedDestroy Action = "planned-destroy"
	Waiting        Action = "waiting"
	Undefined      Action = "undefined"
)
