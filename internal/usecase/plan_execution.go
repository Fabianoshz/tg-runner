package usecase

import (
	"fmt"
	"os/exec"

	"github.com/fabianoshz/iflantis/internal/entity"
)

// TODO use terragrunt show to get output in json
// TODO format the output in human readable way
// TODO treat errors
func (p ExecutionPlannerService) PlanExecution(changelist entity.Changelist) bool {
	lock := entity.AcquireLock()
	defer lock.Release()

	resources := CalculateDependencies(changelist.Resources)

	for _, v := range resources {
		var command string

		if v.Action == entity.Plan {
			command = "/usr/bin/terragrunt plan -out=planfile"
		}

		if v.Action == entity.PlanDestroy {
			command = "/usr/bin/terragrunt plan -destroy -out=planfile"
		}

		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = v.Path
		out, err := cmd.Output()

		p.persistenceRepository.SavePlanfile(v.ID, "planfile", changelist.Id, v.Path)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(out))
	}

	return true
}
