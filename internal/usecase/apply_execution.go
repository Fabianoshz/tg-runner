package usecase

import (
	"fmt"
	"os/exec"

	"github.com/fabianoshz/iflantis/internal/repository"
)

// TODO use terragrunt show to get output in json
// TODO format the output in human readable way
func ApplyExecution(changelist Changelist) bool {
	lock := AcquireLock()
	defer lock.Release()

	resources := CalculateDependencies(changelist.Resources)

	for _, v := range resources {
		var command string

		repository.GetPlanfile(v.ID, "planfile", v.Path)

		if v.Action == Plan {
			command = "/usr/bin/terragrunt apply planfile"
		}

		if v.Action == PlanDestroy {
			command = "/usr/bin/terragrunt destroy planfile"
		}

		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = v.Path
		out, err := cmd.Output()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(out))
	}

	return true
}
