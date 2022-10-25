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
	dependencieGraphs := dependencyCalcultorService.CalculateDependencies(resources, "/home/fabiano/Projects/Fabianoshz/tg-runner/internal/usecase/testdata/terragrunt")

	var resourcesOrder [][]entity.Resource

	for position := 0; position < biggestGraphDepth(dependencieGraphs); position++ {
		resources := getResourcesWithPosition(dependencieGraphs, position)
		resourcesOrder = append(resourcesOrder, resources)
	}

	for i, v := range resourcesOrder {
		if i == 0 {
			for _, resource := range v {
				out, err := executeCommandForResource(resource)
				// TODO check for errors

				// TODO use something better than uuid.New().String() for file ID, maybe nothing?
				// TODO use some output from the command execution instead of hardcoding "planfile"
				p.persistenceRepository.SavePlanfile(uuid.New().String(), "planfile", uuid.New(), resource.Path)

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(string(out))
				// TODO mark the resources in this level as Action.Planned
			}
		} else {
			// TODO mark the resources in this level as Action.Waiting
		}
	}

	return true
}

func biggestGraphDepth(ordered [][]entity.Resource) int {
	var length = 0

	for _, v := range ordered {
		if length < len(v) {
			length = len(v)
		}
	}

	return length
}

func notInResources(resource entity.Resource, resources []entity.Resource) bool {
	add := true
	for _, z := range resources {
		if resource.Path == z.Path {
			add = false
		}
	}

	return add
}

// Return all resources from the dependencies
// graphs that reside in the requested position.
func getResourcesWithPosition(dependencieGraphs [][]entity.Resource, requestedPosition int) []entity.Resource {
	var resources []entity.Resource

	for _, resourceGraphs := range dependencieGraphs {
		for resourcePosition, resource := range resourceGraphs {
			if requestedPosition == resourcePosition && notInResources(resource, resources) {
				resources = append(resources, resource)
			}
		}
	}

	return resources
}

func executeCommandForResource(resource entity.Resource) ([]byte, error) {
	var command string

	if resource.Action == entity.Plan {
		command = "/usr/bin/terragrunt plan -out=planfile"
	}

	if resource.Action == entity.PlanDestroy {
		command = "/usr/bin/terragrunt plan -destroy -out=planfile"
	}

	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = resource.Path

	return cmd.Output()
}
