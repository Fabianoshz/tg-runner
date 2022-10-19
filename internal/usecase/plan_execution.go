package usecase

import (
	"fmt"
	"os/exec"

	"github.com/fabianoshz/tg-runner/internal/entity"
	"github.com/google/uuid"
)

// TODO use terragrunt show to get output in json
// TODO format the output in human readable way
// TODO treat errors
func (p PlanExecutionService) PlanExecution(changelist string) bool {
	resourcesService := NewLoadResourcesService()
	resources := resourcesService.LoadResources(changelist)

	// TODO check if requested resources are locked

	lock := entity.AcquireLock()
	defer lock.Release()

	dependencyCalcultorService := NewCalculateDependenciesService()
	// TODO get rootdir instead of hardcoding
	ordered := dependencyCalcultorService.CalculateDependencies(resources, "/home/fabiano/Projects/Fabianoshz/tg-runner/internal/usecase/testdata/terragrunt")

	for _, v := range ordered {
		for i, v := range v {
			if i == 0 {
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

				// TODO use something better than uuid.New().String() for file ID, maybe nothing?
				p.persistenceRepository.SavePlanfile(uuid.New().String(), "planfile", uuid.New(), v.Path)

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(string(out))
			} else {
				// TODO Add "plan later" to the resources in this level
			}
		}

	}

	return true
}
