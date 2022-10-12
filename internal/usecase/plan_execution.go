package usecase

import (
	"fmt"
	"os/exec"

	"github.com/fabianoshz/iflantis/internal/repository"
)

type Action int64

const (
	Apply Action = iota
	Destroy
	Plan
	PlanDestroy
)

type Resource struct {
	ID     string
	Path   string
	Action Action
}

type Changelist struct {
	Resources []Resource
}

// TODO use terragrunt show to get output in json
// TODO format the output in human readable way
func PlanExecution(changelist Changelist) bool {
	lock := AcquireLock()
	defer lock.Release()

	resources := CalculateDependencies(changelist.Resources)

	for _, v := range resources {
		var command string

		if v.Action == Plan {
			command = "/usr/bin/terragrunt plan -out=planfile"
		}

		if v.Action == PlanDestroy {
			command = "/usr/bin/terragrunt plan -destroy -out=planfile"
		}

		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = v.Path
		out, err := cmd.Output()

		repository.SavePlanfile(v.ID, "planfile", v.Path)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(out))
	}

	return true
}
