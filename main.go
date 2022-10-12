package main

import "github.com/fabianoshz/iflantis/internal/usecase"

func main() {

	changelist := usecase.Changelist{
		Resources: []usecase.Resource{
			{
				ID:     "abc",
				Path:   "internal/usecase/testdata/terragrunt/basic-terragrunt",
				Action: usecase.Plan,
			},
			{
				ID:     "123",
				Path:   "internal/usecase/testdata/terragrunt/basic-terragrunt-2",
				Action: usecase.Plan,
			},
		},
	}

	usecase.PlanExecution(changelist)
}
